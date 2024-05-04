package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

// FindByID implements UserRepository.
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	query := "SELECT id, email, name, password FROM users WHERE id = $1 AND deleted_at IS NULL"
	user := &domain.User{}
	if err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail implements UserRepository.
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "SELECT id, email, name, password FROM users WHERE email = $1 AND deleted_at IS NULL"

	user := &domain.User{}
	if err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

// IsEmailExist implements UserRepository.
func (r *UserRepositoryImpl) IsEmailExist(ctx context.Context, email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND deleted_at IS NULL) "

	var exists bool
	if err := r.db.QueryRow(ctx, query, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// Create implements UserRepository.
func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id"

	if err := tx.QueryRow(ctx, query, user.Email, user.Name, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
