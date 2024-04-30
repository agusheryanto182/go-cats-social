package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
)

type CatService interface {
	Create(ctx context.Context, payload *dto.CatReq) (*dto.CatRes, error)
	GetByID(ctx context.Context, id uint64) (*domain.Cats, error)
	GetByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error)
	IsCatExist(ctx context.Context, catID, userID uint64) (bool, error)
	Update(ctx context.Context, cat *dto.CatReq) (*domain.Cats, error)
	GetByIdAndUserID(ctx context.Context, id, userID uint64) (*domain.Cats, error)
	Delete(ctx context.Context, catID, userID uint64) error
}
