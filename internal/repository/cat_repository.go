package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
)

type CatRepository interface {
	Create(ctx context.Context, cat *domain.Cats) (*domain.Cats, error)
	FindByID(ctx context.Context, id uint64) (*domain.Cats, error)
	IsCatExist(ctx context.Context, catID, userID uint64) (bool, error)
	FindByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error)
	Update(ctx context.Context, cat *domain.Cats) error
	FindByIdAndUserID(ctx context.Context, id, userID uint64) (*domain.Cats, error)
	Delete(ctx context.Context, catID, userID uint64) error
	DoubleUpdateHasMatched(ctx context.Context, catID, userCatID uint64) error
	CheckCats(ctx context.Context, matchCatID, userCatID uint64) ([]*dto.CatResCheck, error)
}
