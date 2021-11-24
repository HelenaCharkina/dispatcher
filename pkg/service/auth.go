package service

import (
	"crypto/sha1"
	"dispatcher/pkg/repository"
	"dispatcher/types"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "gbklmbdflJLK5k5l4349fLKKlld0040392869vbkidgyCG" // соль для пароля
	tokenTTL   = 1 * time.Hour                                    // время жизни токена
	signingKey = "jgkr5494kfmed8d9dkKKKB76hBIUMK9r"               // ключ подписи
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(login string, password string) (string, error) {
	user, err := s.repo.GetUser(login, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id})
	return token.SignedString([]byte(signingKey))
}
