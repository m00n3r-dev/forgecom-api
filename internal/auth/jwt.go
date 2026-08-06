package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secret string
}

func NewJwtService(secret string) *JwtService {
	return &JwtService{
		secret: secret,
	}
}

func (j *JwtService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}
