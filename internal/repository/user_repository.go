package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context, tx pgx.Tx, user *domain.User) (*domain.User, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
