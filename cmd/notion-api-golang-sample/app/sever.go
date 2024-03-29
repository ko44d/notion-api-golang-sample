package app

import (
	"github.com/gin-gonic/gin"
	"ko44d/notion-api-golang-sample/pkg/application"
)

func Init(ndc application.NotionDatabaseController) *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	r.GET("/notionapi/database", ndc.GetDatabaseInfo)
	r.GET("/notionapi/database/query", ndc.GetData)
	return r
}
