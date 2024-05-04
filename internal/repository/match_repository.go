package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
)

type MatchRepository interface {
	Create(ctx context.Context, match *domain.Matches) error
	IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error)
	IsHaveRequest(ctx context.Context, catID uint64) (bool, error)
	FindMatchByIssuedID(ctx context.Context, issuedID uint64) ([]*dto.MatchGetRes, error)
	IsMatchExist(ctx context.Context, id uint64) (*domain.Matches, error)
	DeleteRequestByCatIdAndUserCatID(ctx context.Context, catID, userCatID uint64) error
	ApproveTheMatch(ctx context.Context, matchID, receiverID uint64) error
	Reject(ctx context.Context, matchID, receiverID uint64) error
	DeleteMatchByIssuer(ctx context.Context, id uint64) error
}
