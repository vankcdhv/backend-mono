package cachemanager

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) GetString(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *RedisCache) GetInt64(ctx context.Context, s string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) GetInterface(ctx context.Context, s string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) Add(ctx context.Context, s string, i interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) IsExisted(ctx context.Context, s string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) Lock(ctx context.Context, s string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) Release(ctx context.Context, s string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache) IsLocked(ctx context.Context, s string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisCache() (CacheManager, error) {
	network := viper.GetString("redis.network")
	if network == "" {
		network = "tcp"
	}
	address := viper.GetString("redis.address")
	if address == "" {
		address = "localhost:6379"
	}
	username := viper.GetString("redis.username")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Network:  network,
			Addr:     address,
			Username: username,
			Password: password,
			DB:       db,
		}),
	}, nil
}
