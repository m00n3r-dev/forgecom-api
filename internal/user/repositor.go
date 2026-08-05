package user

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users(name,email,password)
		VALUES($1,$2,$3)
		RETURNING id
	`

	return r.db.QueryRowContext(ctx, query, user.Name,user.Email, user.Password).Scan(&user.ID)
}
