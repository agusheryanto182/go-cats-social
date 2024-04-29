package service

import (
	"context"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/domain"
	"github.com/agusheryanto182/go-social-media/internal/repository"
	"github.com/agusheryanto182/go-social-media/utils/hash"
	"github.com/agusheryanto182/go-social-media/utils/jwt"
	"github.com/jackc/pgx/v5"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	db   *pgx.Conn
	hash hash.HashInterface
	jwt  jwt.IJwt
}

// IsEmailExist implements UserService.
func (s *UserServiceImpl) IsEmailExist(ctx context.Context, email string) (bool, error) {
	return s.repo.IsEmailExist(ctx, email)
}

// Create implements UserService.
func (s *UserServiceImpl) Create(ctx context.Context, payload *dto.UserReq) (*dto.UserCreateRes, error) {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	userReq := &domain.User{
		Email: payload.Email,
		Name:  payload.Name,
	}

	hashed, err := s.hash.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	userReq.Password = hashed

	result, err := s.repo.Create(ctx, tx, userReq)
	if err != nil {
		return nil, err
	}

	token, err := s.jwt.GenerateJWT(result.ID, result.Email)
	if err != nil {
		return nil, err
	}

	return &dto.UserCreateRes{
		Email:       result.Email,
		Name:        result.Name,
		AccessToken: token,
	}, nil
}

func NewUserService(repo repository.UserRepository, db *pgx.Conn, hash hash.HashInterface, jwt jwt.IJwt) UserService {
	return &UserServiceImpl{
		repo: repo,
		db:   db,
		hash: hash,
		jwt:  jwt,
	}
}
