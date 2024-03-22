package database

import (
	"encoding/json"
	"fmt"
	"ko44d/notion-api-golang-sample/pkg/infrastructure/config"
	"net/http"
	"net/url"
)

type NotionAPI interface {
	RetrieveADatabase(param *DatabaseParam) (DatabaseResponse, error)
	QueryADatabase(param *DatabaseParam) (interface{}, error)
}
type notionApi struct {
	config config.Config
}

func NewNotionAPI(config config.Config) NotionAPI {
	return &notionApi{config: config}
}

func (na notionApi) RetrieveADatabase(param *DatabaseParam) (DatabaseResponse, error) {

	u, err := url.Parse(fmt.Sprintf("%s://%s/v1/databases/%s", na.config.NotionAPIProtocol, na.config.NotionAPIHost, param.databaseId))
	if err != nil {
		return DatabaseResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return DatabaseResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", na.config.NotionAPIKey))

	req.Header.Set("Notion-Version", na.config.NotionVersion)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return DatabaseResponse{}, err
	}
	defer res.Body.Close()

	var dr DatabaseResponse
	if err := json.NewDecoder(res.Body).Decode(&dr); err != nil {
		return DatabaseResponse{}, err
	}

	return dr, nil
}

func (na notionApi) QueryADatabase(param *DatabaseParam) (interface{}, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s/v1/databases/%s/query", na.config.NotionAPIProtocol, na.config.NotionAPIHost, param.databaseId))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", na.config.NotionAPIKey))

	req.Header.Set("Notion-Version", na.config.NotionVersion)

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var qad interface{}
	if err := json.NewDecoder(res.Body).Decode(&qad); err != nil {
		return nil, err
	}

	return qad, nil
}
