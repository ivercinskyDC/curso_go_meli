package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ivercinskyDC/curso/integrador/meli"
)

func main() {
	fmt.Fprintf(os.Stdout, "API para Sugerir Precio de Item\n")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})
	meliAPI := meli.API("MLA")
	r.GET("/categories/:category/prices", func(c *gin.Context) {
		category := c.Param("category")
		c.JSON(http.StatusOK, meliAPI.Prices(category))
	})
	r.GET("/search", func(c *gin.Context) {
		searchParams := &meli.SearchParams{}
		searchParams.MethodID = "category"
		searchParams.SearchID = c.Query("id")
		searchParams.SortID = "relevance"
		searchParams.FilterID = ""
		searchParams.Limit = "100"
		searchParams.Offset = ""
		search := meliAPI.Search(searchParams)
		c.JSON(http.StatusOK, search.SearchItems[0])
	})
	r.Run(":8080")
}
