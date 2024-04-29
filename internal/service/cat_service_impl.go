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

// GetByFilterAndArgs implements CatService.
func (s *CatServiceImpl) GetByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error) {
	return s.catRepo.FindByFilterAndArgs(ctx, query, args)
}

// GetByAgeInMonth implements CatService.
func (s *CatServiceImpl) GetByAgeInMonth(ctx context.Context, ageInMonth uint) ([]*domain.Cats, error) {
	panic("unimplemented")
}

// GetByHasMatched implements CatService.
func (s *CatServiceImpl) GetByHasMatched(ctx context.Context, hasMatched bool) ([]*domain.Cats, error) {
	panic("unimplemented")
}

// GetByID implements CatService.
func (s *CatServiceImpl) GetByID(ctx context.Context, id uint64) (*domain.Cats, error) {
	return s.catRepo.FindByID(ctx, id)
}

// GetByLimitAndOffset implements CatService.
func (s *CatServiceImpl) GetByLimitAndOffset(ctx context.Context, limit int, offset int) ([]*domain.Cats, error) {
	panic("unimplemented")
}

// GetByName implements CatService.
func (s *CatServiceImpl) GetByName(ctx context.Context, name string) ([]*domain.Cats, error) {
	panic("unimplemented")
}

// GetByRace implements CatService.
func (s *CatServiceImpl) GetByRace(ctx context.Context, race string) ([]*domain.Cats, error) {
	panic("unimplemented")
}

// GetBySex implements CatService.
func (s *CatServiceImpl) GetBySex(ctx context.Context, sex string) ([]*domain.Cats, error) {
	panic("unimplemented")
}

// GetByUserID implements CatService.
func (s *CatServiceImpl) GetByUserID(ctx context.Context, userID uint64) ([]*domain.Cats, error) {
	panic("unimplemented")
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
