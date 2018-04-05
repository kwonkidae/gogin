package controller

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

var authMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "test zone",
	Key:        []byte("secret key"),
	Timeout:    time.Hour,
	MaxRefresh: time.Hour,
	Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
		fmt.Println(userId, password)
		return userId, true
		// if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		// 	return userId, true
		// }

		// return userId, false
	},
	Authorizator: func(userId string, c *gin.Context) bool {
		fmt.Println("authorizator")
		return true
		// if userId == "admin" {
		// 	return true
		// }

		// return false
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	TokenLookup: "header:Authorization",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName: "Bearer",

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc: time.Now,
}

func init() {
	_init()
	r.Use(cors.Default())
	r.POST("/api/authenticate/", authMiddleware.LoginHandler)
	initArticle()
	initUser()
	r.Run("0.0.0.0:8081")
}

func _init() {
	r = gin.Default()
	r.Static("/assets", "./assets")
}
