package domain

import (
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/database"
	"log"
)

type NotionDatabase interface {
	GetDatabase() interface{}
}

type notionDatabase struct {
	notionapi database.NotionAPI
	config    config.Config
}

func NewNotionDatabase(notionapi database.NotionAPI, config config.Config) NotionDatabase {
	return &notionDatabase{notionapi: notionapi, config: config}
}

func (nd notionDatabase) GetDatabase() interface{} {
	param := database.NewDatabaseParam(nd.config.NotionPageId)

	getDatabase, err := nd.notionapi.GetDatabase(param)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return getDatabase
}
