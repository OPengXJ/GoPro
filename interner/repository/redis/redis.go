package redis

import (
	"github.com/OPengXJ/GoPro/configs"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

var repo = new(redis.Client)

func init() {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.DBaddr,
		Password:     cfg.DBpass,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	if err := client.Ping().Err(); err != nil {
		errors.Wrap(err, "ping redis err")
		panic(err.Error())
	}
	repo=client
}
func GetRedisRepo() *redis.Client {
	return repo
}
