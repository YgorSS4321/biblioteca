package dicontainer

import (
	"treinamento/core/interfaces/repository"
	"treinamento/infra/postgres"
)

func GetBookRepository() repository.BookLoader {
	return postgres.NewBookRepository(GetPSQLConnectionManager())
}

func GetPSQLConnectionManager() postgres.Connector {
	return &postgres.DatabaseConnectionManager{}
}
