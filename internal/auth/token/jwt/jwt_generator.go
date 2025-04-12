package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtService struct {
	secretKey string
	ttl       time.Duration
}

func NewJWTService(secret string, ttl time.Duration) *jwtService {
	return &jwtService{secretKey: secret, ttl: ttl}
}

func (j *jwtService) GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(j.ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}
