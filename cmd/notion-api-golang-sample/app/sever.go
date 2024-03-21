package app

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	return r
}
