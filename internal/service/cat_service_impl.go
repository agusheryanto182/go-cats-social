package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/agusheryanto182/go-social-media/internal/repository"
	"github.com/jackc/pgx/v5"
)

type CatServiceImpl struct {
	catRepo repository.CatRepository
	db      *pgx.Conn
}

// IsCatExist implements CatService.
func (s *CatServiceImpl) IsCatExist(ctx context.Context, catID, userID uint64) (bool, error) {
	return s.catRepo.IsCatExist(ctx, catID, userID)
}

// Update implements CatService.
func (s *CatServiceImpl) Update(ctx context.Context, cat *dto.CatReq) (*domain.Cats, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	payload := &domain.Cats{
		ID:          cat.ID,
		UserID:      cat.UserID,
		Name:        cat.Name,
		Race:        cat.Race,
		Sex:         cat.Sex,
		AgeInMonth:  cat.AgeInMonth,
		Description: cat.Description,
		ImageUrls:   cat.ImageUrls,
	}

	return s.catRepo.Update(ctx, tx, payload)
}

// GetByFilterAndArgs implements CatService.
func (s *CatServiceImpl) GetByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error) {
	return s.catRepo.FindByFilterAndArgs(ctx, query, args)
}

// GetByID implements CatService.
func (s *CatServiceImpl) GetByID(ctx context.Context, id uint64) (*domain.Cats, error) {
	return s.catRepo.FindByID(ctx, id)
}

// Create implements CatService.
func (s *CatServiceImpl) Create(ctx context.Context, payload *dto.CatReq) (*dto.CatRes, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	cat, err := s.catRepo.Create(ctx, tx, &domain.Cats{
		UserID:      payload.UserID,
		Name:        payload.Name,
		Race:        payload.Race,
		Sex:         payload.Sex,
		AgeInMonth:  payload.AgeInMonth,
		Description: payload.Description,
		ImageUrls:   payload.ImageUrls,
	})
	if err != nil {
		return nil, err
	}

	return &dto.CatRes{
		ID:        cat.ID,
		CreatedAt: cat.CreatedAt,
	}, nil
}

func NewCatService(catRepo repository.CatRepository, db *pgx.Conn) CatService {
	return &CatServiceImpl{
		catRepo: catRepo,
		db:      db,
	}
}
