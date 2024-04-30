package repository

import (
	"context"
	"time"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type MatchRepositoryImpl struct {
	db *pgx.Conn
}

// FindMatchByCatID implements MatchRepository.
func (r *MatchRepositoryImpl) FindMatchByCatID(ctx context.Context, userID uint64) ([]*dto.MatchGetRes, error) {
	query := `
	SELECT
		m.id,
		u.name AS issued_by_name,
		u.email AS issued_by_email,
		u.created_at AS issued_by_created_at,
		c1.id AS match_cat_detail_id,
		c1.name AS match_cat_detail_name,
		c1.race AS match_cat_detail_race,
		c1.sex AS match_cat_detail_sex,
		c1.description AS match_cat_detail_description,
		c1.age_in_month AS match_cat_detail_age_in_month,
		c1.image_urls AS match_cat_detail_image_urls,
		c1.has_matched AS match_cat_detail_has_matched,
		c1.created_at AS match_cat_detail_created_at,
		c2.id AS user_cat_detail_id,
		c2.name AS user_cat_detail_name,
		c2.race AS user_cat_detail_race,
		c2.sex AS user_cat_detail_sex,
		c2.description AS user_cat_detail_description,
		c2.age_in_month AS user_cat_detail_age_in_month,
		c2.image_urls AS user_cat_detail_image_urls,
		c2.has_matched AS user_cat_detail_has_matched,
		c2.created_at AS user_cat_detail_created_at,
		m.message,
		m.created_at
	FROM matches m
	JOIN users u ON m.issued_by = u.id
	JOIN cats c1 ON m.match_cat_id = c1.id 
	JOIN cats c2 ON m.user_cat_id = c2.id
	WHERE m.issued_by = $1 OR m.receiver_by = $1
	ORDER BY m.created_at DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	res := []*dto.MatchGetRes{}
	defer rows.Close()

	for rows.Next() {
		match := &dto.MatchGetRes{}
		var issuedByCreatedAt, matchCatDetailCreatedAt, userCatDetailCreatedAt, createdAt time.Time
		if err := rows.Scan(
			&match.ID,
			&match.Issued.Name,
			&match.Issued.Email,
			&issuedByCreatedAt,
			&match.MatchCatDetail.ID,
			&match.MatchCatDetail.Name,
			&match.MatchCatDetail.Race,
			&match.MatchCatDetail.Sex,
			&match.MatchCatDetail.Description,
			&match.MatchCatDetail.AgeInMonth,
			&match.MatchCatDetail.ImageUrls,
			&match.MatchCatDetail.HasMatched,
			&matchCatDetailCreatedAt,
			&match.UserCatDetail.ID,
			&match.UserCatDetail.Name,
			&match.UserCatDetail.Race,
			&match.UserCatDetail.Sex,
			&match.UserCatDetail.Description,
			&match.UserCatDetail.AgeInMonth,
			&match.UserCatDetail.ImageUrls,
			&match.UserCatDetail.HasMatched,
			&userCatDetailCreatedAt,
			&match.Message,
			&createdAt,
		); err != nil {
			return nil, err
		}
		match.Issued.CreatedAt = issuedByCreatedAt.Format(time.RFC3339)
		match.MatchCatDetail.CreatedAt = matchCatDetailCreatedAt.Format(time.RFC3339)
		match.UserCatDetail.CreatedAt = userCatDetailCreatedAt.Format(time.RFC3339)
		match.CreatedAt = createdAt.Format(time.RFC3339)
		res = append(res, match)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

// IsHaveRequest implements MatchRepository.
func (r *MatchRepositoryImpl) IsHaveRequest(ctx context.Context, catID uint64) (bool, error) {
	query := "SELECT EXISTS (SELECT * FROM matches WHERE match_cat_id = $1 OR user_cat_id = $1)"
	var exist bool
	if err := r.db.QueryRow(ctx, query, catID).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
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
		INSERT INTO matches (issued_by, receiver_by,  match_cat_id, user_cat_id, message)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`,
		&match.IssuedBy, &match.ReceiverBy, &match.MatchCatID, &match.UserCatID, &match.Message,
	).Scan(&match.ID)
}

func NewMatchRepository(db *pgx.Conn) MatchRepository {
	return &MatchRepositoryImpl{
		db: db,
	}
}