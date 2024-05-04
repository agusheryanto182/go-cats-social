package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/agusheryanto182/go-social-media/internal/repository"
)

type MatchServiceImpl struct {
	matchRepo repository.MatchRepository
	catRepo   repository.CatRepository
}

// DeleteMatchByIssuer implements MatchService.
func (s *MatchServiceImpl) DeleteMatchByIssuer(ctx context.Context, id uint64) error {

	return s.matchRepo.DeleteMatchByIssuer(ctx, id)
}

// Reject implements MatchService.
func (s *MatchServiceImpl) Reject(ctx context.Context, matchID, receiverID uint64) error {
	return s.matchRepo.Reject(ctx, matchID, receiverID)
}

// ApproveTheMatch implements MatchService.
func (s *MatchServiceImpl) ApproveTheMatch(ctx context.Context, matchID, matchCatID, userCatID, receiverID uint64) error {
	err := s.matchRepo.ApproveTheMatch(ctx, matchID, receiverID)
	if err != nil {
		return err
	}

	err = s.catRepo.DoubleUpdateHasMatched(ctx, matchCatID, userCatID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRequestByCatID implements MatchService.
func (s *MatchServiceImpl) DeleteRequestByCatID(ctx context.Context, catID, userCatID uint64) error {
	return s.matchRepo.DeleteRequestByCatIdAndUserCatID(ctx, catID, userCatID)
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
	return s.matchRepo.Create(ctx, &domain.Matches{
		IssuedBy:   payload.IssuedBy,
		ReceiverBy: payload.ReceiverBy,
		MatchCatID: payload.MatchCatInt,
		UserCatID:  payload.UserCatInt,
		Message:    payload.Message,
	})
}

func NewMatchService(matchRepo repository.MatchRepository, catRepo repository.CatRepository) MatchService {
	return &MatchServiceImpl{
		matchRepo: matchRepo,
		catRepo:   catRepo,
	}
}
