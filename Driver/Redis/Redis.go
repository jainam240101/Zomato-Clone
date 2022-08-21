package redis

import "github.com/go-redis/redis/v8"

type RedisClient struct{ *redis.Client }

func ConnectRedis() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisClient{rdb}
}
