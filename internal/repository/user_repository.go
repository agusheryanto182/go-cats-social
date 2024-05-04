package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uint64) (*domain.User, error)
}
