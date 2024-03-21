package pkg

import (
	"github.com/gin-gonic/gin"
	"ko44d/notion-api-golang-sample/cmd/notion-api-golang-sample/app"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
)

type DI struct {
}

func (di DI) Init() *gin.Engine {
	return app.Init()
}

func (di DI) Config() config.Config {
	c, err := config.GetFromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
