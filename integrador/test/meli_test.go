package test

import (
	"testing"

	"github.com/ivercinskyDC/curso/integrador/meli"
)

func TestMeliCategory(t *testing.T) {
	meliAPI := meli.API("MLA")
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
	meliAPI := meli.API("MLA")
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

}
