package core

import "github.com/gin-gonic/gin"

func Routers() (router *gin.Engine) {
	Router := gin.Default()
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})
	return Router
}
