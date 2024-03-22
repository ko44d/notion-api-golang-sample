package application

import (
	"github.com/gin-gonic/gin"
	"ko44d/notion-api-golang-sample/pkg/domain"
)

type NotionDatabaseController interface {
	GetDatabase(ctx *gin.Context)
}

type notionDatabaseController struct {
	database domain.NotionDatabase
}

func NewDatabaseService(database domain.NotionDatabase) NotionDatabaseController {
	return &notionDatabaseController{database: database}
}

func (ndc notionDatabaseController) GetDatabase(ctx *gin.Context) {
	ctx.JSON(200, ndc.database.GetDatabase())
}
