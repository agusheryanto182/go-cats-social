package dto

import (
	"database/sql"
	"time"
)

type CatRes struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type CatAllsRes struct {
	ID          string       `json:"id"`
	UserID      uint64       `json:"-"`
	Name        string       `json:"name"`
	Race        string       `json:"race"`
	Sex         string       `json:"sex"`
	AgeInMonth  uint         `json:"ageInMonth"`
	ImageUrls   []string     `json:"imageUrls"`
	Description string       `json:"description"`
	HasMatched  bool         `json:"hasMatched"`
	DeletedAt   sql.NullTime `json:"-"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   time.Time    `json:"-"`
}

type CatResCheck struct {
	HasMatched bool       `json:"hasMatched"`
	DeletedAt  *time.Time `json:"deletedAt"`
}
