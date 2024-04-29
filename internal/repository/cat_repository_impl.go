package repository

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/jackc/pgx/v5"
)

type CatRepositoryImpl struct {
	db *pgx.Conn
}

// Create implements CatRepository.
func (r *CatRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, cat *domain.Cats) (*domain.Cats, error) {
	query := "INSERT INTO cats (user_id, name, race, sex, age_in_month, description, image_urls) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at"
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
