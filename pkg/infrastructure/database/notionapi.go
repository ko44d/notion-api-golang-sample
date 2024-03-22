package database

import (
	"ko44d/notion-api-golang-sample/pkg/domain"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
)

type notionApiDatabase struct {
	config config.Config
}

func NewNotionAPIDatabase(config config.Config) domain.Database {
	return &notionApiDatabase{config: config}
}

func (na notionApiDatabase) GetDatabase(param DatabaseParam) DatabaseResponse {
	return DatabaseResponse{}
}
