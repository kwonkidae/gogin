package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	_init()
	initArticle()
	r.Use(cors.Default())
	r.Run("0.0.0.0:8081")
}

func _init() {
	r = gin.Default()
}
