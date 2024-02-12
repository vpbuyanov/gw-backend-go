package postgresql

import (
	"github.com/vpbuyanov/gw-backend-go/internal/databases/postgres"
)

type postgresqlUS struct {
	repos *postgres.Postgresql
}

type USPostgresql interface {
}

func NewPostgresqlUseCase(repos *postgres.Postgresql) USPostgresql {
	return &postgresqlUS{repos: repos}
}
