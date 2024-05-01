package dto

import "time"

type MatchGetRes struct {
	ID             uint64     `json:"id"`
	IssuedBy       uint64     `json:"-"`
	Issued         UserMatch  `json:"issuedBy"`
	MatchCatID     uint64     `json:"-"`
	MatchCatDetail CatMatches `json:"matchCatDetail"`
	UserCatID      uint64     `json:"-"`
	UserCatDetail  CatMatches `json:"userCatDetail"`
	IsApproved     bool       `json:"-"`
	Message        string     `json:"message"`
	CreatedAt      string     `json:"createdAt"`
	DeletedAt      *time.Time `json:"-"`
}

type UserMatch struct {
	ID        uint64 `json:"-"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type CatMatches struct {
	ID          uint64   `json:"id"`
	UserID      uint64   `json:"-"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  uint     `json:"ageInMonth"`
	ImageUrls   []string `json:"imageUrls"`
	Description string   `json:"description"`
	HasMatched  bool     `json:"hasMatched"`
	CreatedAt   string   `json:"createdAt"`
}
