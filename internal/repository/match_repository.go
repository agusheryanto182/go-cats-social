package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type MatchRepository interface {
	Create(ctx context.Context, tx pgx.Tx, match *domain.Matches) error
	IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error)
	IsHaveRequest(ctx context.Context, catID uint64) (bool, error)
	FindMatchByCatID(ctx context.Context, catID uint64) ([]*dto.MatchGetRes, error)
	IsMatchExist(ctx context.Context, id, userID uint64) (*domain.Matches, error)
	DeleteRequestByCatIdAndUserCatID(ctx context.Context, tx pgx.Tx, catID, userCatID uint64) error
	ApproveTheMatch(ctx context.Context, tx pgx.Tx, matchID, receiverID uint64) error
}
