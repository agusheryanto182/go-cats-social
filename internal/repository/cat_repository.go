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
	FindByLimitAndOffset(ctx context.Context, limit int, offset int) ([]*domain.Cats, error)
	FindByRace(ctx context.Context, race string) ([]*domain.Cats, error)
	FindBySex(ctx context.Context, sex string) ([]*domain.Cats, error)
	FindByHasMatched(ctx context.Context, hasMatched bool) ([]*domain.Cats, error)
	FindByAgeInMonth(ctx context.Context, ageInMonth uint) ([]*domain.Cats, error)
	FindByOwned(ctx context.Context, userID uint64) ([]*domain.Cats, error)
	FindByName(ctx context.Context, name string) ([]*domain.Cats, error)
	FindByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error)
}
