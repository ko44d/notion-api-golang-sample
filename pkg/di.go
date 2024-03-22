package pkg

import (
	"github.com/gin-gonic/gin"
	"ko44d/notion-api-golang-sample/cmd/notion-api-golang-sample/app"
	"ko44d/notion-api-golang-sample/pkg/application"
	"ko44d/notion-api-golang-sample/pkg/domain"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/database"
)

type DI struct {
}

func (di DI) Init() *gin.Engine {
	return app.Init(di.NotionDatabaseController())
}

func (di DI) Config() config.Config {
	c, err := config.GetFromEnv()
	if err != nil {
		panic(err)
	}
	return c
}

func (di DI) NotionDatabaseController() application.NotionDatabaseController {
	return application.NewDatabaseService(di.NotionDatabase())
}

func (di DI) NotionDatabase() domain.NotionDatabase {
	return domain.NewNotionDatabase(di.NotionAPI(), di.Config())
}

func (di DI) NotionAPI() database.NotionAPI {
	return database.NewNotionAPI(di.Config())
}
