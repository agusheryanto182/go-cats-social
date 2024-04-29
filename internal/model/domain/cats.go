package domain

import "time"

type Cats struct {
	ID               uint64    `json:"id"`
	UserID           uint64    `json:"userId"`
	Name             string    `json:"name"`
	Race             string    `json:"race"`
	Sex              string    `json:"sex"`
	Description      string    `json:"description"`
	AgeInMonth       uint      `json:"ageInMonth"`
	IsAlreadyMatched bool      `json:"hasMatched"`
	ImageUrls        []string  `json:"imageUrls"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	User             *User     `json:"user"`
}
