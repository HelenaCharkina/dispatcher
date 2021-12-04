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
	repo  repository.User
	cache repository.Token
}

func NewAuthService(repo repository.User, cache repository.Token) *AuthService {
	return &AuthService{
		repo:  repo,
		cache: cache,
	}
}

func (s *AuthService) SignIn(login string, password string) (*types.Response, error) {
	user, err := s.repo.GetUser(login, s.generatePasswordHash(password))
	if err != nil {
		return nil, err
	}

	response, err := s.generateTokens(user)
	if err != nil {
		return nil, err
	}

	err = s.setRefreshToken(response.RefreshToken, response.User.Id)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AuthService) RefreshToken(refreshToken string, userId string) (*types.Response, error) {
	err := s.checkRefreshToken(refreshToken, userId)
	if err != nil {
		return nil, err
	}

	response, err := s.generateTokens(&types.User{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	err = s.setRefreshToken(response.RefreshToken, userId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AuthService) Logout(userId string) error {
	return s.cache.RemoveToken(userId)
}

func (s *AuthService) CheckToken(accessToken string) (string, error) {
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

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) generateTokens(user *types.User) (*types.Response, error) {
	accessToken, err := s.generateToken(user.Id)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateRefreshToken()
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

func (s *AuthService) generateToken(userId string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(settings.Config.TokenTTL * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generateRefreshToken() (string, error) {
	b := make([]byte, 32)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) setRefreshToken(token string, userId string) error {
	return s.cache.SetToken(token, userId)
}

func (s *AuthService) checkRefreshToken(token string, userId string) error {
	refreshToken, err := s.cache.GetToken(userId)
	if err != nil {
		return err
	}

	if refreshToken != token {
		return errors.New("invalid refresh token")
	}
	return nil
}
