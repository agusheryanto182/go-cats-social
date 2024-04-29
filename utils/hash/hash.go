package hash

import (
	"github.com/agusheryanto182/go-social-media/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type HashInterfaceImpl struct {
	cfg *config.Global
}

type HashInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

func NewHash(cfg *config.Global) HashInterface {
	return &HashInterfaceImpl{
		cfg: cfg,
	}
}

func (h *HashInterfaceImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cfg.Bcrypt.Salt)
	return string(bytes), err
}

func (h *HashInterfaceImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
