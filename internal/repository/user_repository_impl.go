package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepositoryImpl struct {
	db *pgx.Conn
}

// IsEmailExist implements UserRepository.
func (r *UserRepositoryImpl) IsEmailExist(ctx context.Context, email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"

	var exists bool
	if err := r.db.QueryRow(ctx, query, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// Create implements UserRepository.
func (r *UserRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id"

	if err := tx.QueryRow(ctx, query, user.Email, user.Name, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
