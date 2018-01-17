package controller

import (
	"db"
	"image/png"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func initArticle() {
	r.POST("/writeArticle", writeArticle)
	r.POST("/fileupload", fileUpload)
}

// WriteArticle ...
func writeArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()
	// _ := context.DbCollection("article")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// FileUpload ...
func fileUpload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	log.Println(file.Size)

	// decode jpeg into image.Image
	f, _ := file.Open()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Thumbnail(50, 50, img, resize.Lanczos3)

	out, err := os.Create("test_thumbnail111.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
	// c.SaveUploadedFile(file, dst)

	c.JSON(200, gin.H{
		"result": "true",
	})
}
