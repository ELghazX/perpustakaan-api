package service

import (
	"context"
	"errors"
	"time"

	"github.com/elghazx/perpustakaan/domain"
	"github.com/elghazx/perpustakaan/dto"
	"github.com/elghazx/perpustakaan/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

func (as authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := as.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	if user.ID == "" {
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	claim := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Duration(as.conf.Jwt.Exp) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(as.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}
