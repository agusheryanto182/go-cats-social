package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type MatchRepository interface {
	Create(ctx context.Context, tx pgx.Tx, match *domain.Matches) error
	IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error)
	IsHaveRequest(ctx context.Context, catID uint64) (bool, error)
}
