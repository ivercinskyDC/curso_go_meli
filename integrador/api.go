package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ivercinskyDC/curso/integrador/meli"
)

func pricesHandler(c *gin.Context) {
	meliAPI := meli.API("MLA")
	category := c.Param("category")
	resp, err := meliAPI.Prices(category)
	if err == nil {
		c.JSON(http.StatusOK, resp)
	} else {
		//se deberia diferenciar los errores 500 de lo 4xx
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func setUpServer() *gin.Engine {
	r := gin.Default()
	r.StaticFile("/", "./frontend/index.html")
	r.GET("/categories/:category/prices", pricesHandler)
	return r
}
func main() {
	fmt.Fprintf(os.Stdout, "API para Sugerir Precio de Item\n")
	r := setUpServer()
	r.Run(":8080")
}
