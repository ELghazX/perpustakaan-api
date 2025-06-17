package domain

import (
	"context"

	"github.com/elghazx/perpustakaan/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
