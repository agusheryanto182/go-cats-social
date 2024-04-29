package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
)

type CatService interface {
	Create(ctx context.Context, payload *dto.CatReq) (*dto.CatRes, error)
}
