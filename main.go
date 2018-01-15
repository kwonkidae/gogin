package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Conext) {
		name := r.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
