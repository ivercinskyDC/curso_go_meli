package meli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Search is a wrapper for the http call to
func (m *Meli) Search(params *SearchParams) *SearchResult {
	url := "https://api.mercadolibre.com/sites/" + m.SiteID + "/search?"
	if params.MethodID == "" {
		panic("MethoID is required")
	}
	if params.SearchID == "" {
		panic("SearchID is required")
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
	fmt.Fprintf(os.Stdout, "Calling %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		searchResponse := &SearchResult{}
		rawBody, readError := ioutil.ReadAll(resp.Body)
		if readError != nil {
			panic(readError)
		} else {
			unmarshalError := json.Unmarshal(rawBody, searchResponse)
			if unmarshalError != nil {
				panic(unmarshalError)
			}
			return searchResponse
		}
	}
}

func (m *Meli) getSuggestions(items []SearchItem) *Suggestion {
	if len(items) == 0 {
		return &Suggestion{0, 0, 0}
	}
	var min, max, avg float64
	min = items[0].Price
	max = 0
	avg = 0
	for _, item := range items {
		if item.Price < min {
			min = item.Price
		}
		if item.Price > max {
			max = item.Price
		}
		avg += item.Price
	}
	avg /= float64(len(items))
	return &Suggestion{max, avg, min}
}

//Prices returns an estimated price for the category
func (m *Meli) Prices(CatID string) *Suggestion {
	searchParams := &SearchParams{}
	searchParams.MethodID = "category"
	searchParams.SearchID = CatID
	searchParams.SortID = "relevance"
	searchParams.FilterID = ""
	searchParams.Limit = "100"
	searchParams.Offset = ""
	response := m.Search(searchParams)
	res := m.getSuggestions(response.SearchItems)
	return res
}

//API get the API Wrapper for a specific SITE
func API(SiteID string) *Meli {
	m := &Meli{SiteID: SiteID}
	return m
}
