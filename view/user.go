package view

import (
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"message": "add user",
	})
}

func UserDetail(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"message": "user detail",
	})
}

func UserIdentify(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"message": "user identify",
	})
}
