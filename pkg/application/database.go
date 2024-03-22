package application

import (
	"github.com/gin-gonic/gin"
	"ko44d/notion-api-golang-sample/pkg/domain"
)

type NotionDatabaseController interface {
	GetDatabaseInfo(ctx *gin.Context)
	GetData(ctx *gin.Context)
}

type notionDatabaseController struct {
	database domain.NotionDatabase
}

func NewDatabaseService(database domain.NotionDatabase) NotionDatabaseController {
	return &notionDatabaseController{database: database}
}

func (ndc notionDatabaseController) GetDatabaseInfo(ctx *gin.Context) {
	ctx.JSON(200, ndc.database.GetDatabaseInfo())
}

func (ndc notionDatabaseController) GetData(ctx *gin.Context) {
	ctx.JSON(200, ndc.database.GetData())
}
