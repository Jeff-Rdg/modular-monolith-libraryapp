package usecase

import (
	"context"
	"modular-monolith-libraryApp/modules/user/domain"
	"modular-monolith-libraryApp/modules/user/port"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CPF      string `json:"cpf"`
}

type CreateUserInterface interface {
	Execute(ctx context.Context, input CreateUserInput) error
}

type CreateUserUseCase struct {
	repo port.UserRepository
}

func NewCreateUserUseCase(repo port.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo}
}

func (c *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) error {
	err := c.repo.Save(ctx, domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		CPF:      input.CPF,
	})

	return err
}
