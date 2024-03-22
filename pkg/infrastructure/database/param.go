package database

type DatabaseParam struct {
	databaseId string `json:"database_id"`
}

func NewDatabaseParam(databaseId string) *DatabaseParam {
	return &DatabaseParam{databaseId: databaseId}
}
