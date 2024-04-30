package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
)

type MatchService interface {
	Create(ctx context.Context, payload *dto.MatchReq) error
	IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error)
	IsHaveRequest(ctx context.Context, catID uint64) (bool, error)
	GetMatch(ctx context.Context, userID uint64) ([]*dto.MatchGetRes, error)
	IsMatchExist(ctx context.Context, id, userID uint64) (*domain.Matches, error)
	DeleteRequestByCatID(ctx context.Context, catID, userCatID uint64) error
	ApproveTheMatch(ctx context.Context, matchID, matchCatID, userCatID, receiverID uint64) error
	Reject(ctx context.Context, matchID, receiverID uint64) error
}
