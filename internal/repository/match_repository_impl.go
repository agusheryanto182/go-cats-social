package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type MatchRepositoryImpl struct {
	db *pgx.Conn
}

func (r *MatchRepositoryImpl) IsRequestExist(ctx context.Context, matchCatID, userCatID uint64) (bool, error) {
	query := "SELECT EXISTS (SELECT * FROM matches WHERE match_cat_id = $1 AND user_cat_id = $2)"
	var exist bool
	if err := r.db.QueryRow(ctx, query, matchCatID, userCatID).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

// Create implements MatchRepository.
func (r *MatchRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, match *domain.Matches) error {
	return tx.QueryRow(ctx, `
		INSERT INTO matches (issued_by, match_cat_id, user_cat_id, message)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		&match.IssuedBy, &match.MatchCatID, &match.UserCatID, &match.Message,
	).Scan(&match.ID)
}

func NewMatchRepository(db *pgx.Conn) MatchRepository {
	return &MatchRepositoryImpl{
		db: db,
	}
}
