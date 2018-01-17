package controller

import (
	"db"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initArticle() {
	r.GET("/", WriteArticle)
	r.POST("/fileupload", FileUpload)
}

// WriteArticle ...
func WriteArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()
	// _ := context.DbCollection("article")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// FileUpload ...
func FileUpload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
