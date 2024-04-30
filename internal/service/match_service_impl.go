package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/agusheryanto182/go-social-media/internal/repository"
	"github.com/jackc/pgx/v5"
)

type MatchServiceImpl struct {
	db        *pgx.Conn
	matchRepo repository.MatchRepository
}

// GetMatch implements MatchService.
func (s *MatchServiceImpl) GetMatch(ctx context.Context, userID uint64) ([]*dto.MatchGetRes, error) {
	return s.matchRepo.FindMatchByCatID(ctx, userID)
}

// IsHaveRequest implements MatchService.
func (s *MatchServiceImpl) IsHaveRequest(ctx context.Context, catID uint64) (bool, error) {
	return s.matchRepo.IsHaveRequest(ctx, catID)
}

// IsRequestExist implements MatchService.
func (s *MatchServiceImpl) IsRequestExist(ctx context.Context, matchCatID uint64, userCatID uint64) (bool, error) {
	return s.matchRepo.IsRequestExist(ctx, matchCatID, userCatID)
}

// Create implements MatchService.
func (s *MatchServiceImpl) Create(ctx context.Context, payload *dto.MatchReq) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	return s.matchRepo.Create(ctx, tx, &domain.Matches{
		IssuedBy:   payload.IssuedBy,
		ReceiverBy: payload.ReceiverBy,
		MatchCatID: payload.MatchCatInt,
		UserCatID:  payload.UserCatInt,
		Message:    payload.Message,
	})
}

func NewMatchService(db *pgx.Conn, matchRepo repository.MatchRepository) MatchService {
	return &MatchServiceImpl{
		db:        db,
		matchRepo: matchRepo,
	}
}
