package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
)

type MatchService interface {
	Create(ctx context.Context, payload *dto.MatchReq) error
	IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error)
	IsHaveRequest(ctx context.Context, catID uint64) (bool, error)
	GetMatch(ctx context.Context, userID uint64) ([]*dto.MatchGetRes, error)
}
