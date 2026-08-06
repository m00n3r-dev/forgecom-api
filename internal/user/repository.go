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

	return r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID)
}

func (r *Repository) GetUserFromEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id,name,email,password FROM users
		WHERE email = $1
	`

	var user User

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
