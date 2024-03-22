package domain

import "ko44d/notion-api-golang-sample/pkg/infrastructure/database"

type Database interface {
	GetDatabase(param database.DatabaseParam) database.DatabaseResponse
}
