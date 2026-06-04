package repository

import (
	"context"
	"modular-monolith-libraryApp/modules/author/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthorRepository struct {
	db *pgxpool.Pool
}

func NewAuthorRepository(db *pgxpool.Pool) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) FindByID(ctx context.Context, id string) (domain.Author, error) {
	row := a.db.QueryRow(ctx,
		`
		SELECT id, 
			   name, 
			   created_at
		FROM authors WHERE id = $1`, id)
	var author domain.Author
	err := row.Scan(
		&author.ID,
		&author.Name,
		&author.CreatedAt,
	)

	if err != nil {
		return domain.Author{}, err
	}

	return author, nil
}

func (a *AuthorRepository) Save(ctx context.Context, author domain.Author) error {
	_, err := a.db.Exec(ctx,
		`
		INSERT INTO authors (id, name)
		VALUES ($1, $2)`,
		author.ID,
		author.Name,
	)

	return err
}

func (a *AuthorRepository) ExistsAll(ctx context.Context, ids []string) (bool, error) {
	var count int
	err := a.db.QueryRow(ctx,
		`
		SELECT count(*) FROM authors WHERE id = ANY($1)`, ids).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
