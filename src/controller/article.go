package controller

import (
	"db"
	"fmt"
	"image/png"
	"log"
	"model"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

var lock sync.RWMutex

func initArticle() {

	g := r.Group("article")
	{
		g.POST("/", writeArticle)
		g.POST("/image", uploadImage)
		g.GET("/", getAllArticle)
		g.PUT("/", updateArticle)
		g.GET("/:articleNo", getArticle)
		g.PUT("/addlikecount", addLikeCount)
		g.PUT("/adddislikecount", addDislikeCount)
	}
	r.POST("/fileupload", fileUpload)
}

func uploadImage(c *gin.Context) {
	u1 := uuid.Must(uuid.NewV4())

	dst := "assets/article/image"
	file, err := c.FormFile("file")
	id := c.PostForm("id")
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	err = os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ext := path.Ext(file.Filename)
	path := filepath.Join(dst, fmt.Sprintf("%s_%s%s", id, u1, ext))
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"imageLink": path,
	})
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
	c.BindJSON(&article)
	col := context.DbCollection("article")
	err := article.InsertArticle(col)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, gin.H{"article_no": article.ArticleNo})
}

// updateArticle ..
func updateArticle(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	var article model.Article
	c.BindJSON(&article)
	fmt.Println(article)
	col := context.DbCollection("article")
	err := article.UpdateArticle(col)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"article_no": article.ArticleNo})
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
