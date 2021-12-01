package repository

import (
	"context"
	"dispatcher/pkg/settings"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisCache(ctx context.Context) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", settings.Config.RedisHost, settings.Config.RedisPort)

	fmt.Println("addr ", addr)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Connection ping error: %v ", err)
		return nil, err
	}

	return client, nil
}
