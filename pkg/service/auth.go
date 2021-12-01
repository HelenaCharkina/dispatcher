package service

import (
	"crypto/sha1"
	"dispatcher/pkg/repository"
	"dispatcher/pkg/settings"
	"dispatcher/types"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

const (
	salt       = "gbklmbdflJLK5k5l4349fLKKlld0040392869vbkidgyCG" // соль для пароля
	signingKey = "jgkr5494kfmed8d9dkKKKB76hBIUMK9r"               // ключ подписи
)

type AuthService struct {
	repo  repository.Authorization
	cache repository.Token
}

func NewAuthService(repo repository.Authorization, cache repository.Token) *AuthService {
	return &AuthService{
		repo:  repo,
		cache: cache,
	}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) SignIn(login string, password string) (*types.Response, error) {
	user, err := s.repo.GetUser(login, s.generatePasswordHash(password))
	if err != nil {
		return nil, err
	}

	return s.GenerateTokens(user)
}

func (s *AuthService) GenerateTokens(user *types.User) (*types.Response, error) {
	accessToken, err := s.GenerateToken(user.Id)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	response := types.Response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: types.User{
			Name: user.Name,
			Id:   user.Id,
		},
	}

	return &response, nil
}

func (s *AuthService) GenerateToken(userId string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(settings.Config.TokenTTL * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &types.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*types.TokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) SetRefreshToken(token string, userId string) error {
	return s.cache.SetToken(token, userId)
}

func (s *AuthService) CheckRefreshToken(token string, userId string) error {
	refreshToken, err := s.cache.GetToken(userId)
	if err != nil {
		return err
	}

	if refreshToken != token {
		return errors.New("invalid refresh token")
	}
	return nil
}
