package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/m00n3r-dev/forgecom-api/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository       *Repository
	jwtService       *auth.JwtService
	refreshTokenRepo *auth.RefreshTokenRepository
}

func NewService(repository *Repository, jwtService *auth.JwtService, refreshTokenRepo *auth.RefreshTokenRepository) *Service {
	return &Service{
		repository:       repository,
		jwtService:       jwtService,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest, userAgent string) (string, string, error) {
	user, err := s.repository.GetUserFromEmail(ctx, req.Email)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.jwtService.GenerateToken(user.ID.String())
	if err != nil {
		return "", "", err
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	familyID := uuid.NewString()

	err = s.refreshTokenRepo.Create(ctx, auth.RefreshToken{
		UserID:    user.ID.String(),
		FamilyID:  familyID,
		TokenHash: string(hashedRefreshToken),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		UserAgent: userAgent,
		IPAddress: "127.0.0.1",
	})

	return accessToken, refreshToken, nil

}
