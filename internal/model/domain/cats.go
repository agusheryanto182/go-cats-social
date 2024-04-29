package domain

type Cats struct {
	ID          uint64   `json:"id"`
	UserID      uint64   `json:"-"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  uint     `json:"ageInMonth"`
	ImageUrls   []string `json:"imageUrls"`
	Description string   `json:"description"`
	HasMatched  bool     `json:"hasMatched"`
	DeletedAt   string   `json:"-"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"-"`
}
