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
	r.StaticFile("/", "./frontend/index.html")
	meliAPI := meli.API("MLA")
	r.GET("/categories/:category/prices", func(c *gin.Context) {
		category := c.Param("category")
		resp, err := meliAPI.Prices(category)
		if err == nil {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	r.Run(":8080")
}
