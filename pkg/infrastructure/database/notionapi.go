package database

import (
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
)

type NotionAPI interface {
	GetDatabase(param *DatabaseParam) DatabaseResponse
}
type notionApi struct {
	config config.Config
}

func NewNotionAPI(config config.Config) NotionAPI {
	return &notionApi{config: config}
}

func (na notionApi) GetDatabase(param *DatabaseParam) DatabaseResponse {
	return DatabaseResponse{}
}
