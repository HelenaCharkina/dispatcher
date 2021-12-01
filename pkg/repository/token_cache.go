package repository

import (
	"context"
	"dispatcher/pkg/settings"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenCache struct {
	client *redis.Client
}

func NewTokenCache(client *redis.Client) *TokenCache {
	return &TokenCache{
		client: client,
	}
}

func (c *TokenCache) GetToken(userId string) (string, error) {
	cacheToken, err := c.client.Get(context.TODO(), fmt.Sprintf("token-%s", userId)).Result()
	if err != nil {
		return "", err
	}
	return cacheToken, nil
}

func (c *TokenCache) SetToken(token string, userId string) error {
	return c.client.Set(context.TODO(), fmt.Sprintf("token-%s", userId), token, time.Minute*settings.Config.RefreshTokenTTL).Err()
}

func (c *TokenCache) RemoveToken(userId string) error {
	return c.client.Del(context.TODO(), fmt.Sprintf("token-%s", userId)).Err()
}
