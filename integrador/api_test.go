package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ivercinskyDC/curso/integrador/meli"
)

func testPrice(r *gin.Engine, t *testing.T, sug chan *meli.Suggestion, catID string) {
	fmt.Printf("Calling %s", catID)
	req, err := http.NewRequest("GET", "/categories/"+catID+"/prices", nil)
	if err != nil {
		t.Fatalf("No se esperaba que falle el test, pero tuvo un error: %s", err.Error())
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatalf("Se esperaba codigo %d, pero vino %d", 200, resp.Code)
	}
	sugRes := &meli.Suggestion{}
	json.Unmarshal(resp.Body.Bytes(), sugRes)
	sug <- sugRes
}

func TestApiPriceOk(t *testing.T) {
	r := setUpServer()
	suggestions := make(chan (*meli.Suggestion), 2)
	go testPrice(r, t, suggestions, "MLA31372")

	go testPrice(r, t, suggestions, "MLA31372")

	<-suggestions
}
