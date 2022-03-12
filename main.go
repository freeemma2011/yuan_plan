package main

import (
	//"github.com/freeemma2011/yuan_plan/dal"
	"github.com/gin-gonic/gin"
)

func main() {
	//dal.InitDB()
	r := gin.Default() //创建默认路由，其实就是一个handler　　 //添加路由，不用为什么是 /ping(你可以可以设置为/user, /login)，劳资复制GitHub官方的案例，保持原汁原味，所以多看官方文档
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})      //默认开启8080，如果需要开启其他的，可以设置r.Run(9090)
	r.Run() // listen and serve on 0.0.0.0:8080
}
