package repository

import (
	"context"
	"modular-monolith-libraryApp/modules/user/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindByCPF(ctx context.Context, cpf string) (domain.User, error) {
	row := u.db.QueryRow(ctx,
		`
		SELECT id, 
       		   name, 
       		   password,
       		   cpf,
       		   created_at,
       		   updated_at
		FROM users WHERE cpf = $1`, cpf)

	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.CPF,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Save(ctx context.Context, user domain.User) error {
	_, err := u.db.Exec(ctx,
		`
		INSERT INTO users (name, password, cpf, email)
		VALUES ($1, $2, $3, $4)`,
		user.Name,
		user.Password,
		user.CPF,
		user.Email,
	)
	return err
}
