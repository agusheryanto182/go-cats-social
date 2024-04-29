package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type CatRepository interface {
	Create(ctx context.Context, tx pgx.Tx, cat *domain.Cats) (*domain.Cats, error)
}
