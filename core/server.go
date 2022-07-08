package core

import "github.com/gin-gonic/gin"

func ServerRun() {
	Router := gin.Default()

	// router

	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})
	//运行
	Router.Run(":8088")
}
