package dto

type CatReq struct {
	UserID      uint64   `json:"userId" validate:"required,gte=1"`
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required,oneof=Persian MaineCoon Siamese Ragdoll Bengal Sphynx BritishShorthair Abyssinian ScottishFold Birman"`
	Sex         string   `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  uint     `json:"ageInMonth" validate:"required,gte=1,lte=120082"`
	Description string   `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string `json:"imageUrls" validate:"required,min=1,dive,required,url"`
}

const (
	Persian          string = "Persian"
	MaineCoon        string = "Maine Coon"
	Siamese          string = "Siamese"
	Ragdoll          string = "Ragdoll"
	Bengal           string = "Bengal"
	Sphynx           string = "Sphynx"
	BritishShorthair string = "British Shorthair"
	Abyssinian       string = "Abyssinian"
	ScottishFold     string = "Scottish Fold"
	Birman           string = "Birman"
)
