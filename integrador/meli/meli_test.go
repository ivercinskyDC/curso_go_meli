package meli

import (
	"testing"
)

func TestMeliAPI(t *testing.T) {
	meliAPI := API("MLA")

	if meliAPI.SiteID != "MLA" {
		t.Fatalf("Se esperaba que el Site de la API sea %s y fue %s", "MLA", meliAPI.SiteID)
	}
}
func TestMeliCategory(t *testing.T) {
	meliAPI := API("MLA")
	catID := "MLA31372"
	cat, err := meliAPI.Category(catID)
	if err != nil {
		t.Fatal("Error Inesperado", err)
	}
	expected := 0

	if len(cat.ChildrenCategories) != expected {
		t.Fatalf("Error Esperaba que esa %s no tengo subcategorias pero tiene %d", catID, len(cat.ChildrenCategories))
	}
}

func TestMeliCategoryFail(t *testing.T) {
	meliAPI := API("MLA")
	catID := "123"
	_, err := meliAPI.Category(catID)
	if err == nil {
		t.Fatal("Error Esperado", err)
	}
	expected := "Category not found: " + catID

	if err.Error() != expected {
		t.Fatalf("\nEsperaba %s, pero recibi %s", expected, err.Error())
	}
}

func TestMeliSearch(t *testing.T) {
	meliAPI := API("MLA")
	catID := "MLA31372"
	searchParams := &SearchParams{}
	searchParams.MethodID = "category"
	searchParams.SearchID = catID
	searchParams.SortID = "relevance"
	searchParams.FilterID = ""
	searchParams.Limit = "200"
	searchParams.Offset = ""
	searchResult, err := meliAPI.Search(searchParams)
	if err != nil {
		t.Fatalf("\nNo se esperaba que falle el test")
	} else {
		if searchResult.SiteID != "MLA" {
			t.Fatalf("\nSe esperaba que el site sea %s y fue %s", "MLA", searchResult.SiteID)
		}
		if searchResult.SearchItems[0].CategoryID != catID {
			t.Fatalf("\nSe esperaba que el primer resultado tuviera ID %s y fue %s", catID, searchResult.SearchItems[0].CategoryID)
		}
	}
}

func TestMeliSearchSinResultados(t *testing.T) {
	meliAPI := API("MLA")
	catID := "123"
	searchParams := &SearchParams{}
	searchParams.MethodID = "category"
	searchParams.SearchID = catID
	searchParams.SortID = "relevance"
	searchParams.FilterID = ""
	searchParams.Limit = "200"
	searchParams.Offset = ""
	searchResult, err := meliAPI.Search(searchParams)
	if err != nil {
		t.Fatalf("\nNo se esperaba que falle el test")
	} else {
		if searchResult.SiteID != "MLA" {
			t.Fatalf("\nSe esperaba que el site sea %s y fue %s", "MLA", searchResult.SiteID)
		}
		if len(searchResult.SearchItems) > 0 {
			t.Fatalf("\nSe esperaba que el resultado no tuviera resultados, pero tuvo %d", len(searchResult.SearchItems))
		}
	}
}

func TestPrices(t *testing.T) {
	meliAPI := API("MLA")
	catID := "MLA31372"
	expected := &Suggestion{22900 + 10, 2280, 45 - 10}
	suggestion, err := meliAPI.Prices(catID)
	if err != nil {
		t.Fatalf("\nNo se esperaba que falle el test")
	} else {
		if suggestion.Max > expected.Max {
			t.Fatalf("\nSe esperaba que el precio maximo fuera mayor a : %f , pero fue %f", expected.Max, suggestion.Max)
		}
		if suggestion.Min < expected.Min {
			t.Fatalf("\nSe esperaba que el precio minimo fuera: %f , pero fue %f", expected.Min, suggestion.Min)
		}
		if suggestion.Suggested > expected.Suggested+10 && suggestion.Suggested < expected.Suggested-10 {
			t.Fatalf("\nSe esperaba que el precio sugerido estuviera entre: (%f,%f) , pero fue %f", expected.Suggested-10, expected.Suggested+10, suggestion.Suggested)
		}
	}

}

func TestPricesCatIncorrecta(t *testing.T) {
	meliAPI := API("MLA")
	catID := "123"
	_, err := meliAPI.Prices(catID)
	if err == nil && err.Error() != "Category not found: 123" {
		t.Fatalf("\nSe esperaba que falle el test con error %s, pero fue %s", "Category not found: 123", err.Error())
	}
}

func TestPricesCatConSubCats(t *testing.T) {
	meliAPI := API("MLA")
	catID := "MLA4769"
	_, err := meliAPI.Prices(catID)
	if err == nil && err.Error() != "La categoria no puede tener subcategorias" {
		t.Fatalf("\nSe esperaba que falle el test con error %s, pero fue %s", "La categoria no puede tener subcategorias", err.Error())
	}
}
