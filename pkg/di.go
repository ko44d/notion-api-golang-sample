package pkg

import (
	"github.com/gin-gonic/gin"
	"ko44d/notion-api-golang-sample/cmd/notion-api-golang-sample/app"
)

type DI struct {
}

func (d DI) Init() *gin.Engine {
	return app.Init()
}
