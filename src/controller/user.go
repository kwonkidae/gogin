package controller

import (
	"db"
	"log"
	"model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initUser() {
	g := r.Group("user")
	{
		g.POST("/", createUser)
		g.PUT("/", updateUser)
		g.DELETE("/", deleteUser)
	}
}

func createUser(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	icol := context.DbCollection("identity")
	ucol := context.DbCollection("user")
	var user model.User

	c.BindJSON(&user)
	err := user.CreateUser(ucol, icol)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, user)
}

func updateUser(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	var user model.User
	c.BindJSON(&user)
	col := context.DbCollection("user")
	err := user.UpdateUser(col)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user_no": user.UserNo})
}

func deleteUser(c *gin.Context) {
	context := db.NewContext()
	defer context.Close()

	var user model.User
	c.BindJSON(&user)
	col := context.DbCollection("user")
	err := user.DeleteUser(col)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user_no": user.UserNo})
}
