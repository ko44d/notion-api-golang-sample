package config

import (
	"errors"
	"os"
)

type Config struct {
	NotionAPIProtocol string
	NotionAPIHost     string
	NotionAPIKey      string
	NotionPageId      string
	NotionVersion     string
}

func GetFromEnv() (Config, error) {

	nak := os.Getenv("NOTION_API_KEY")
	if nak == "" {
		return Config{}, errors.New("NOTION_API_KEY is empty.")
	}

	npi := os.Getenv("NOTION_PAGE_ID")
	if npi == "" {
		return Config{}, errors.New("NOTION_PAGE_ID is empty")
	}

	return Config{
		NotionAPIProtocol: "https",
		NotionAPIHost:     "api.notion.com",
		NotionAPIKey:      nak,
		NotionPageId:      npi,
		NotionVersion:     "2022-06-28",
	}, nil
}
