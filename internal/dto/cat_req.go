package dto

type CatReq struct {
	ID          uint64   `json:"id" validate:"required,gte=1"`
	UserID      uint64   `json:"userId" validate:"required,gte=1"`
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required,oneof=Persian 'Maine Coon' Siamese Ragdoll Bengal Sphynx 'British Shorthair' Abyssinian 'Scottish Fold' Birman"`
	Sex         string   `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  uint     `json:"ageInMonth" validate:"required,gte=1,lte=120082"`
	Description string   `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string `json:"imageUrls" validate:"required,min=1,dive,required,url"`
}
