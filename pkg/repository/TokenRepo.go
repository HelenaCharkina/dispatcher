package repository

import (
	"context"
	"dispatcher/pkg/settings"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenRepo struct {
	client *redis.Client
}

func NewTokenRepo(client *redis.Client) *TokenRepo {
	return &TokenRepo{
		client: client,
	}
}

func (c *TokenRepo) GetToken(userId string) (string, error) {
	cacheToken, err := c.client.Get(context.TODO(), fmt.Sprintf("token-%s", userId)).Result()
	if err != nil {
		return "", err
	}
	return cacheToken, nil
}

func (c *TokenRepo) SetToken(token string, userId string) error {
	return c.client.Set(context.TODO(), fmt.Sprintf("token-%s", userId), token, time.Minute*settings.Config.RefreshTokenTTL).Err()
}

func (c *TokenRepo) RemoveToken(userId string) error {
	return c.client.Del(context.TODO(), fmt.Sprintf("token-%s", userId)).Err()
}
