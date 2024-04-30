package repository

import (
	"context"
	"strings"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type CatRepositoryImpl struct {
	db *pgx.Conn
}

// FindDoubleCats implements CatRepository.
func (r *CatRepositoryImpl) CheckCats(ctx context.Context, matchCatID, userCatID uint64) ([]*dto.CatResCheck, error) {
	query := "SELECT has_matched, deleted_at FROM cats WHERE (id = $1 OR id = $2)"
	rows, err := r.db.Query(ctx, query, matchCatID, userCatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cats := []*dto.CatResCheck{}

	for rows.Next() {
		cat := &dto.CatResCheck{}
		err = rows.Scan(&cat.HasMatched, &cat.DeletedAt)
		if err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

// DoubleUpdateHasMatched implements CatRepository.
func (r *CatRepositoryImpl) DoubleUpdateHasMatched(ctx context.Context, tx pgx.Tx, catID uint64, userCatID uint64) error {
	query := "UPDATE cats SET has_matched = true WHERE id = $1 OR id = $2 AND deleted_at IS NULL"
	rows, err := tx.Exec(ctx, query, catID, userCatID)
	if err != nil {
		return err
	}

	if rows.RowsAffected() != 2 {
		return pgx.ErrNoRows
	}

	return nil
}

// Delete implements CatRepository.
func (r *CatRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, catID uint64, userID uint64) error {
	query := "UPDATE cats SET deleted_at = NOW() WHERE id = $1 AND user_id = $2"
	res, err := tx.Exec(ctx, query, catID, userID)
	if err != nil {
		return err
	}
	numRowsAffected := res.RowsAffected()

	if numRowsAffected == 0 {
		return err
	}
	return nil
}

// FindByIdAndUserID implements CatRepository.
func (r *CatRepositoryImpl) FindByIdAndUserID(ctx context.Context, id uint64, userID uint64) (*domain.Cats, error) {
	query := "SELECT id, name, race, sex, age_in_month, description, image_urls, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at FROM cats WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL"
	cat := &domain.Cats{}

	if err := r.db.QueryRow(ctx, query, id, userID).Scan(&cat.ID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls, &cat.CreatedAt); err != nil {
		return nil, err
	}
	cat.UserID = userID
	return cat, nil
}

// IsCatExist implements CatRepository.
func (r *CatRepositoryImpl) IsCatExist(ctx context.Context, catID, userID uint64) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM cats WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL) "
	var exist bool
	if err := r.db.QueryRow(ctx, query, catID, userID).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

// Update implements CatRepository.
func (r *CatRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, cat *domain.Cats) (*domain.Cats, error) {
	query := "UPDATE cats SET name = $2, race = $3, sex = $4, age_in_month = $5, description = $6, image_urls = $7 WHERE id = $1 AND user_id = $8 AND deleted_at IS NULL RETURNING name, race, sex, age_in_month, description, image_urls, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at"

	row := tx.QueryRow(ctx, query, cat.ID, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ImageUrls, cat.UserID)
	if err := row.Scan(
		&cat.Name,
		&cat.Race,
		&cat.Sex,
		&cat.AgeInMonth,
		&cat.Description,
		&cat.ImageUrls,
		&cat.CreatedAt,
	); err != nil {
		return nil, err
	}
	return cat, nil
}

// FindByFilterAndArgs implements CatRepository.
func (r *CatRepositoryImpl) FindByFilterAndArgs(ctx context.Context, query string, args []interface{}) ([]*dto.CatAllsRes, error) {
	rows, err := r.db.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cats := []*dto.CatAllsRes{}

	for rows.Next() {
		cat := &dto.CatAllsRes{}
		err := rows.Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.HasMatched, &cat.ImageUrls, &cat.CreatedAt)
		if err != nil {
			return nil, err
		}

		for i, img := range cat.ImageUrls {
			cat.ImageUrls[i] = strings.TrimSpace(img)
		}

		cats = append(cats, cat)
	}
	return cats, nil
}

// FindByID implements CatRepository.
func (r *CatRepositoryImpl) FindByID(ctx context.Context, id uint64) (*domain.Cats, error) {
	query := "SELECT id, user_id, name, race, sex, description, age_in_month, has_matched, image_urls, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at FROM cats WHERE id = $1 AND deleted_at IS NULL LIMIT 1"

	cat := &domain.Cats{}
	if err := r.db.QueryRow(ctx, query, id).Scan(
		&cat.ID,
		&cat.UserID,
		&cat.Name,
		&cat.Race,
		&cat.Sex,
		&cat.Description,
		&cat.AgeInMonth,
		&cat.HasMatched,
		&cat.ImageUrls,
		&cat.CreatedAt,
	); err != nil {
		return nil, err
	}

	return cat, nil
}

// Create implements CatRepository.
func (r *CatRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, cat *domain.Cats) (*domain.Cats, error) {
	query := "INSERT INTO cats (user_id, name, race, sex, age_in_month, description, image_urls) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at"
	if err := tx.QueryRow(ctx, query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ImageUrls).Scan(&cat.ID, &cat.CreatedAt); err != nil {
		return nil, err
	}
	return cat, nil
}

func NewCatRepository(db *pgx.Conn) CatRepository {
	return &CatRepositoryImpl{
		db: db,
	}
}
