package user

import (
	"context"
	"errors"

	"github.com/m00n3r-dev/forgecom-api/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
	jwtService *auth.JwtService
}

func NewService(repository *Repository, jwtService *auth.JwtService) *Service {
	return &Service{
		repository: repository,
		jwtService: jwtService,
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

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	user, err := s.repository.GetUserFromEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.jwtService.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil

}
