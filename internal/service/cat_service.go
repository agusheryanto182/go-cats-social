package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
)

type CatService interface {
	Create(ctx context.Context, payload *dto.CatReq) (*dto.CatRes, error)
	GetByID(ctx context.Context, id uint64) (*domain.Cats, error)
	GetByLimitAndOffset(ctx context.Context, limit int, offset int) ([]*domain.Cats, error)
	GetByRace(ctx context.Context, race string) ([]*domain.Cats, error)
	GetBySex(ctx context.Context, sex string) ([]*domain.Cats, error)
	GetByHasMatched(ctx context.Context, hasMatched bool) ([]*domain.Cats, error)
	GetByAgeInMonth(ctx context.Context, ageInMonth uint) ([]*domain.Cats, error)
	GetByUserID(ctx context.Context, userID uint64) ([]*domain.Cats, error)
	GetByName(ctx context.Context, name string) ([]*domain.Cats, error)
	GetByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error)
}
