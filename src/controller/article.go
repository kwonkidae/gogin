package controller

import (
	"db"
	"fmt"
	"image/png"
	"log"
	"model"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

var lock sync.Mutex

func initArticle() {

	g := r.Group("article")
	{
		g.POST("", writeArticle)
		g.PUT("/addlikecount", addLikeCount)
		g.PUT("/adddislikecount", addDislikeCount)
	}
	r.POST("/fileupload", fileUpload)
}

// WriteArticle ...
func writeArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	lock.Lock()
	defer lock.Unlock()

	var article model.Article
	c.BindJSON(&article)
	col := context.DbCollection("article")
	err := article.InsertArticle(col)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{})
}

func addLikeCount(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()
	col := context.DbCollection("article")
	var article model.Article
	c.BindJSON(&article)
	article.AddLikeCount(col)
	fmt.Println(article)
	c.JSON(200, gin.H{})
}

func addDislikeCount(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()
	col := context.DbCollection("article")
	var article model.Article
	c.BindJSON(&article)
	article.AddDislikeCount(col)
	fmt.Println(article)
	c.JSON(200, gin.H{})
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
