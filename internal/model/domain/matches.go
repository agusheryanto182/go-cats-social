package domain

import "time"

type Match struct {
	ID         uint64    `json:"id"`
	UserID     uint64    `json:"userId"`
	CatID      uint64    `json:"catId"`
	User       *User     `json:"user"`
	Cat        *Cats     `json:"cat"`
	IsApproved bool      `json:"isApproved"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
