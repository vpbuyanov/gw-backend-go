package usecase

import (
	"github.com/vpbuyanov/gw-backend-go/internal/databases/postgres/repos"
)

type useCase struct {
	repos *repos.Pg
}

type UseCase interface {
}

func NewUseCase(repos *repos.Pg) UseCase {
	return &useCase{repos: repos}
}
