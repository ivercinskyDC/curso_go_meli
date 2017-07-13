package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ivercinskyDC/curso/integrador/meli"
)

func makeHTTPCall(t *testing.T, r *gin.Engine, catID string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/categories/"+catID+"/prices", nil)
	if err != nil {
		t.Fatalf("No se esperaba que falle el request, pero tuvo un error: %s", err.Error())
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}

func testPrice(t *testing.T, r *gin.Engine, sug chan *meli.Suggestion, catID string, suggested float64) {
	resp := makeHTTPCall(t, r, catID)
	if resp.Code != 200 {
		t.Fatalf("Se esperaba codigo %d, pero vino %d", 200, resp.Code)
	}
	sugRes := &meli.Suggestion{}
	json.Unmarshal(resp.Body.Bytes(), sugRes)
	if sugRes.Suggested < suggested-10 && sugRes.Suggested > suggested+10 {
		t.Fatalf("Se esperaba que el precio estuviera entre %f, %f, pero dio %f", suggested-10, suggested+10, sugRes.Suggested)
	}
	sug <- sugRes
}

func TestApiPriceOk(t *testing.T) {
	r := setUpServer()
	suggestions := make(chan (*meli.Suggestion), 2)
	go testPrice(t, r, suggestions, "MLA31372", 2410.0)

	go testPrice(t, r, suggestions, "MLA31372", 2390.0)

	<-suggestions
}

type ErrorResponse struct {
	ErrorMsg string `json:"error"`
}

func TestApiPriceError(t *testing.T) {
	r := setUpServer()
	resp := makeHTTPCall(t, r, "123")
	if resp.Code != 500 {
		t.Fatalf("Se esperaba codigo %d, pero vino %d", 500, resp.Code)
	}
	respError := &ErrorResponse{}
	err := json.Unmarshal(resp.Body.Bytes(), respError)
	if err != nil {
		t.Fatalf("Se espera que el body de la respuesta fuera de tipo Error")
	}
	if respError.ErrorMsg != "Category not found: 123" {
		t.Fatalf("Se espera Mensaje de Error %s, pero recibi %s", "Category not found: 123", respError.ErrorMsg)
	}

	resp = makeHTTPCall(t, r, "MLA4769")
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
