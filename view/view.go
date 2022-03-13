package view

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {

	r.GET("/user/detail/", UserDetail)
	r.GET("/user/add/", AddUser)
	r.GET("/user/identify/", UserIdentify)

	r.NoRoute(func(c *gin.Context) {
		log.Printf("url doesn't match, uri=%v", c.Request.RequestURI)
		c.Header("tt_stable", "1")
		c.JSON(200, gin.H{
			"status_code": 404,
			"status_msg":  "url doesn't match",
		})
	})
}
