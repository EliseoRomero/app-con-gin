package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ejempro de Go Web Gin
/*
func main() {
	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Buenas": "Goland",
		})
	})
	r.Run()
}
*/

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Bienvenido a nuestra primera app de Go con Gin %s", name)
	})
	router.Run(":8080")
}
