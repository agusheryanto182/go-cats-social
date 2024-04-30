package domain

import "time"

type Matches struct {
	ID         uint64     `json:"id"`
	IssuedBy   uint64     `json:"issuedBy"`
	ReceiverBy uint64     `json:"-"`
	MatchCatID uint64     `json:"matchCatId"`
	UserCatID  uint64     `json:"userCatId"`
	IsApproved bool       `json:"isApproved"`
	Message    string     `json:"message"`
	DeletedAt  *time.Time `json:"-"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"-"`
}
