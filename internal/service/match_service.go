package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
)

type MatchService interface {
	Create(ctx context.Context, payload *dto.MatchReq) error
	IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error)
}
