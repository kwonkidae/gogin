package controller

import (
	"db"
	"fmt"
	"image/png"
	"log"
	"model"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"gopkg.in/mgo.v2/bson"
)

var lock sync.RWMutex

func initArticle() {

	g := r.Group("article")
	{
		g.POST("/", writeArticle)
		g.GET("/", getAllArticle)
		g.GET("/:articleNo", getArticle)
		g.PUT("/addlikecount", addLikeCount)
		g.PUT("/adddislikecount", addDislikeCount)
	}
	r.POST("/fileupload", fileUpload)
}

func getArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	lock.RLock()
	defer lock.RUnlock()

	col := context.DbCollection("article")
	var article model.Article

	articleNo := c.Param("articleNo")
	no, err := strconv.Atoi(articleNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	article.ArticleNo = no
	article.GetArticle(col)
	fmt.Println(article)
	c.JSON(200, article)
}

// getAllArticle is 게시글 목록을 모두 다 가지고 온다.
func getAllArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	lock.RLock()
	defer lock.RUnlock()

	var articles []model.Article
	col := context.DbCollection("article")
	if err := col.Find(bson.M{}).All(&articles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, articles)
}

// WriteArticle ...
func writeArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	lock.Lock()
	defer lock.Unlock()

	var article model.Article

	col := context.DbCollection("article")
	err := article.InsertArticle(col)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
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
