package port

import (
	"context"
	"modular-monolith-libraryApp/modules/author/domain"
)

type AuthorRepository interface {
	FindByID(ctx context.Context, id string) (domain.Author, error)
	Save(ctx context.Context, author domain.Author) error
	ExistsAll(ctx context.Context, ids []string) (bool, error)
}
