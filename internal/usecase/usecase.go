package usecase

import "github.com/vpbuyanov/gw-backend-go/internal/databases/repos"

type useCase struct {
	repos *repos.Repos
}

type UseCase interface {
}

func NewUseCase(repos *repos.Repos) UseCase {
	return &useCase{repos: repos}
}
