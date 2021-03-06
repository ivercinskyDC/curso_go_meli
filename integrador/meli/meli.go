package meli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Category returns the Category
func (m *Meli) Category(CatID string) (*Category, error) {
	//TODO: No hardcodear la URL
	url := "https://api.mercadolibre.com/categories/"
	url += CatID
	fmt.Fprintf(os.Stdout, "Calling %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	category := &Category{}
	err = json.Unmarshal(rawBody, category)
	if err != nil {
		return nil, err
	}
	if category.ID == "" {
		catNotFound := &CategoryNotFound{}
		err = json.Unmarshal(rawBody, catNotFound)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(catNotFound.Message)
	}
	return category, nil
}

//Search is a wrapper for the http call to
func (m *Meli) Search(params *SearchParams) (*SearchResult, error) {
	//TODO: No hardcodear la URL
	url := "https://api.mercadolibre.com/sites/" + m.SiteID + "/search?"
	if params.MethodID == "" {
		return nil, errors.New("MethoID is required")
	}
	if params.SearchID == "" {
		return nil, errors.New("SearchID is required")
	}
	url += params.MethodID
	url += "=" + params.SearchID
	if params.SortID != "" {
		url += "&sort=" + params.SortID
	}
	if params.FilterID != "" {
		url += "&filter=" + params.FilterID
	}
	if params.Limit != "" {
		url += "&limit=" + params.Limit
	}
	if params.Offset != "" {
		url += "&offset=" + params.Offset
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	searchResponse := &SearchResult{}
	rawBody, readError := ioutil.ReadAll(resp.Body)
	if readError != nil {
		return nil, readError
	}
	unmarshalError := json.Unmarshal(rawBody, searchResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return searchResponse, nil

}

func (m *Meli) getSuggestion(items []SearchItem) (*Suggestion, error) {
	if len(items) == 0 {
		return &Suggestion{0, 0, 0}, nil
	}
	var min, max, avg float64
	var solds int
	min = items[0].Price
	max = 0
	avg = 0
	solds = 0
	//TODO: Se podria mejorar haciendo un kMeans.
	//Con el Kmeans podemos encontrar subcategorias en los items.
	//Luego podriamos devolver el promedio del grupo con mas elementos
	//O varias sugerencias. (Quizas agregar mas categorias en estos casos)
	for _, item := range items {
		if item.Condition != "new" {
			continue
		}
		if item.Price < min {
			min = item.Price
		}
		if item.Price > max {
			max = item.Price
		}
		avg += item.Price * float64(item.SoldQuantity)
		solds += item.SoldQuantity
	}
	avg /= float64(solds)
	return &Suggestion{max, avg, min}, nil
}

//Prices returns an estimated price for the category
func (m *Meli) Prices(CatID string) (*Suggestion, error) {
	cat, err := m.Category(CatID)
	//Busco la categoria para no calcular el precio sugerido de una categoria con muchos elementos diferentes.
	if err != nil {
		return nil, err
	}
	if len(cat.ChildrenCategories) > 0 {
		return nil, errors.New("La categoria no puede tener subcategorias")
	}
	//TODO: hacer un multiget para recuperar todos los items de la categoria. haciendo ItemsEnCategoria / 200
	searchParams := &SearchParams{}
	searchParams.MethodID = "category"
	searchParams.SearchID = cat.ID
	searchParams.SortID = "relevance"
	searchParams.FilterID = ""
	searchParams.Limit = "200"
	searchParams.Offset = ""
	response, err := m.Search(searchParams)
	if err != nil {
		return nil, err
	}
	//meli limita el limite a 200. Hay que recuperar en varios llamados
	return m.getSuggestion(response.SearchItems)
}

//API get the API Wrapper for a specific SITE
func API(SiteID string) *Meli {
	m := &Meli{SiteID: SiteID}
	return m
}
