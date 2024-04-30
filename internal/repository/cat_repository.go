package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type CatRepository interface {
	Create(ctx context.Context, tx pgx.Tx, cat *domain.Cats) (*domain.Cats, error)
	FindByID(ctx context.Context, id uint64) (*domain.Cats, error)
	IsCatExist(ctx context.Context, catID, userID uint64) (bool, error)
	FindByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error)
	Update(ctx context.Context, tx pgx.Tx, cat *domain.Cats) (*domain.Cats, error)
}
