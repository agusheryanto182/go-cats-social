package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
)

type UserService interface {
	Create(ctx context.Context, payload *dto.UserReq) (*dto.UserCreateRes, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
