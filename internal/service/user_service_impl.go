package service

import (
	"context"
	"errors"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/agusheryanto182/go-social-media/internal/repository"
	"github.com/agusheryanto182/go-social-media/utils/hash"
	"github.com/agusheryanto182/go-social-media/utils/jwt"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	hash hash.HashInterface
	jwt  jwt.IJwt
}

// GetByID implements UserService.
func (s *UserServiceImpl) GetByID(ctx context.Context, id uint64) (*dto.UserResWithID, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResWithID{
		ID:    result.ID,
		Email: result.Email,
		Name:  result.Name,
	}, nil
}

// Login implements UserService.
func (s *UserServiceImpl) Login(ctx context.Context, payload *dto.UserLoginReq) (*dto.UserRes, error) {
	user, err := s.repo.FindByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	if !s.hash.CheckPasswordHash(payload.Password, user.Password) {
		return nil, errors.New("wrong password")
	}

	token, err := s.jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.UserRes{
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}

// IsEmailExist implements UserService.
func (s *UserServiceImpl) IsEmailExist(ctx context.Context, email string) (bool, error) {
	return s.repo.IsEmailExist(ctx, email)
}

// Create implements UserService.
func (s *UserServiceImpl) Create(ctx context.Context, payload *dto.UserReq) (*dto.UserRes, error) {

	userReq := &domain.User{
		Email: payload.Email,
		Name:  payload.Name,
	}

	hashed, err := s.hash.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	userReq.Password = hashed

	result, err := s.repo.Create(ctx, userReq)
	if err != nil {
		return nil, err
	}

	token, err := s.jwt.GenerateJWT(result.ID, result.Email)
	if err != nil {
		return nil, err
	}

	return &dto.UserRes{
		Email:       result.Email,
		Name:        result.Name,
		AccessToken: token,
	}, nil
}

func NewUserService(repo repository.UserRepository, hash hash.HashInterface, jwt jwt.IJwt) UserService {
	return &UserServiceImpl{
		repo: repo,
		hash: hash,
		jwt:  jwt,
	}
}
