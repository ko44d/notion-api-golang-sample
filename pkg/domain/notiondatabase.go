package domain

import (
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/notionapi"
	"log"
)

type NotionDatabase interface {
	GetDatabaseInfo() interface{}
	GetData() interface{}
}

type notionDatabase struct {
	notionapi notionapi.NotionAPI
	config    config.Config
}

func NewNotionDatabase(notionapi notionapi.NotionAPI, config config.Config) NotionDatabase {
	return &notionDatabase{notionapi: notionapi, config: config}
}

func (nd notionDatabase) GetDatabaseInfo() interface{} {
	param := notionapi.NewDatabaseParam(nd.config.NotionPageId)

	getDatabase, err := nd.notionapi.RetrieveADatabase(param)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return getDatabase
}

func (nd notionDatabase) GetData() interface{} {
	param := notionapi.NewDatabaseParam(nd.config.NotionPageId)

	qad, err := nd.notionapi.QueryADatabase(param)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return qad
}
