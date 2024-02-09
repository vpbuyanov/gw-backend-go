package repos

import (
	"database/sql"
	"fmt"
)

type repos struct {
	db *sql.DB
}

type Repos interface {
}

func NewRepos(db *sql.DB) (Repos, error) {
	if db == nil {
		return nil, fmt.Errorf("point db is nil")
	}

	return &repos{
		db: db,
	}, nil
}
