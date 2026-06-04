package port

import (
	"context"
	"modular-monolith-libraryApp/modules/user/domain"
)

type UserRepository interface {
	FindByCPF(ctx context.Context, cpf string) (domain.User, error)
	Save(ctx context.Context, user domain.User) error
}
