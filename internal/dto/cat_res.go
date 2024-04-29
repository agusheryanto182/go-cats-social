package dto

import (
	"database/sql"
	"time"
)

type CatRes struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type CatAllsRes struct {
	ID          uint64       `json:"id"`
	UserID      uint64       `json:"-"`
	Name        string       `json:"name"`
	Race        string       `json:"race"`
	Sex         string       `json:"sex"`
	AgeInMonth  uint         `json:"ageInMonth"`
	ImageUrls   []string     `json:"imageUrls"`
	Description string       `json:"description"`
	HasMatched  bool         `json:"hasMatched"`
	DeletedAt   sql.NullTime `json:"-"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"-"`
}
