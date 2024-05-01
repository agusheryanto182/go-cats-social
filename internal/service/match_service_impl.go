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
	catRepo   repository.CatRepository
}

// DeleteMatchByIssuer implements MatchService.
func (s *MatchServiceImpl) DeleteMatchByIssuer(ctx context.Context, id uint64) error {
	tx, err := s.db.Begin(ctx)

	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	return s.matchRepo.DeleteMatchByIssuer(ctx, tx, id)
}

// Reject implements MatchService.
func (s *MatchServiceImpl) Reject(ctx context.Context, matchID, receiverID uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	return s.matchRepo.Reject(ctx, tx, matchID, receiverID)
}

// ApproveTheMatch implements MatchService.
func (s *MatchServiceImpl) ApproveTheMatch(ctx context.Context, matchID, matchCatID, userCatID, receiverID uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	err = s.matchRepo.ApproveTheMatch(ctx, tx, matchID, receiverID)
	if err != nil {
		return err
	}

	err = s.catRepo.DoubleUpdateHasMatched(ctx, tx, matchCatID, userCatID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRequestByCatID implements MatchService.
func (s *MatchServiceImpl) DeleteRequestByCatID(ctx context.Context, catID, userCatID uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	return s.matchRepo.DeleteRequestByCatIdAndUserCatID(ctx, tx, catID, userCatID)
}

// IsMatchExist implements MatchService.
func (s *MatchServiceImpl) IsMatchExist(ctx context.Context, id uint64) (*domain.Matches, error) {
	return s.matchRepo.IsMatchExist(ctx, id)
}

// GetMatch implements MatchService.
func (s *MatchServiceImpl) GetMatch(ctx context.Context, userID uint64) ([]*dto.MatchGetRes, error) {
	return s.matchRepo.FindMatchByIssuedID(ctx, userID)
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

func NewMatchService(db *pgx.Conn, matchRepo repository.MatchRepository, catRepo repository.CatRepository) MatchService {
	return &MatchServiceImpl{
		db:        db,
		matchRepo: matchRepo,
		catRepo:   catRepo,
	}
}
