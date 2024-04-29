package dto

type UserRes struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type UserResWithID struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
