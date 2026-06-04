package usecase

import (
	"context"
	"modular-monolith-libraryApp/modules/author/domain"
	"modular-monolith-libraryApp/modules/author/port"
)

type CreateAuthorInput struct {
	Name string `json:"name" validate:"required"`
}

type CreateAuthor interface {
	Execute(ctx context.Context, input CreateAuthorInput) error
}

type createAuthorUseCase struct {
	repo port.AuthorRepository
}

func NewCreateAuthor(repo port.AuthorRepository) CreateAuthor {
	return &createAuthorUseCase{repo: repo}
}

func (c *createAuthorUseCase) Execute(ctx context.Context, input CreateAuthorInput) error {
	return c.repo.Save(ctx, domain.Author{
		Name: input.Name,
	})
}
