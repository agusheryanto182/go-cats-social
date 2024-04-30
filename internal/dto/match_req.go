package dto

type MatchReq struct {
	IssuedBy    uint64 `json:"issuedBy" validate:"required,gte=1"`
	ReceiverBy  uint64
	MatchCatID  string `json:"matchCatId" validate:"required"`
	UserCatID   string `json:"userCatId" validate:"required"`
	MatchCatInt uint64
	UserCatInt  uint64
	Message     string `json:"message" validate:"required,min=1,max=120"`
}

type MatchIdReq struct {
	MatchID    string `json:"matchId" validate:"required,gte=1"`
	MatchIdInt uint64
}
