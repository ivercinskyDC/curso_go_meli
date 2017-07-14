package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ivercinskyDC/curso/integrador/meli"
)

var server *gin.Engine

func TestMain(m *testing.M) {
	server = setUpServer()
	exit := m.Run()
	os.Exit(exit)
}

func makeHTTPCall(r *gin.Engine, catID string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", "/categories/"+catID+"/prices", nil)
	if err != nil {
		return nil, err

	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp, nil
}

func testPrice(t *testing.T, r *gin.Engine, sug chan *meli.Suggestion, catID string, suggested float64) error {
	resp, err := makeHTTPCall(r, catID)
	if err != nil {
		return err
	}
	sugRes := &meli.Suggestion{}
	json.Unmarshal(resp.Body.Bytes(), sugRes)
	if sugRes.Suggested < suggested-10 && sugRes.Suggested > suggested+10 {
		t.Fatalf("Se esperaba que el precio estuviera entre %f, %f, pero dio %f", suggested-10, suggested+10, sugRes.Suggested)
	}
	sug <- sugRes
	return nil
}

func TestApiPriceOk(t *testing.T) {
	cant := 100
	suggestions := make(chan (*meli.Suggestion), cant)
	for i := 0; i < cant; i++ {
		go testPrice(t, server, suggestions, "MLA31372", 2410.0)
	}
	<-suggestions
}

type ErrorResponse struct {
	ErrorMsg string `json:"error"`
}

func TestApiPriceError(t *testing.T) {
	resp, err := makeHTTPCall(server, "123")
	if err != nil {
		t.Fatalf("No se esperaba que falle el request, pero tuvo un error: %s", err.Error())
	}
	if resp.Code != 500 {
		t.Fatalf("Se esperaba codigo %d, pero vino %d", 500, resp.Code)
	}
	respError := &ErrorResponse{}
	err = json.Unmarshal(resp.Body.Bytes(), respError)
	if err != nil {
		t.Fatalf("Se espera que el body de la respuesta fuera de tipo Error")
	}
	if respError.ErrorMsg != "Category not found: 123" {
		t.Fatalf("Se espera Mensaje de Error %s, pero recibi %s", "Category not found: 123", respError.ErrorMsg)
	}

	resp, err = makeHTTPCall(server, "MLA4769")
	if err != nil {
		t.Fatalf("No se esperaba que falle el request, pero tuvo un error: %s", err.Error())
	}
	if resp.Code != 500 {
		t.Fatalf("Se esperaba codigo %d, pero vino %d", 500, resp.Code)
	}

	respError = &ErrorResponse{}
	err = json.Unmarshal(resp.Body.Bytes(), respError)
	if err != nil {
		t.Fatalf("Se espera que el body de la respuesta fuera de tipo Error")
	}
	if respError.ErrorMsg != "La categoria no puede tener subcategorias" {
		t.Fatalf("Se espera Mensaje de Error %s, pero recibi %s", "La categoria no puede tener subcategorias", respError.ErrorMsg)
	}
}

func BenchmarkPricesOK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeHTTPCall(server, "MLA31372")
	}
}

func BenchmarkWrongCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeHTTPCall(server, "123")
	}
}

func BenchmarkCatWithSubCAts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeHTTPCall(server, "MLA4769")
	}
}
