package usecase

import (
	"context"
	"fmt"
	"poc/internal/onboarding/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTToken struct {
	secretKey string
	publicKey string
}

func NewJWTToken(secretKey, publicKey string) *JWTToken {
	return &JWTToken{
		secretKey: secretKey,
		publicKey: publicKey,
	}
}

func (j *JWTToken) GenerateToken(ctx context.Context, user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "poc",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:        fmt.Sprintf("%d", user.ID),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "poc",
		Subject:   fmt.Sprintf("%d", user.ID),
	})

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
