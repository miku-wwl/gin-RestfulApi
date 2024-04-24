package cache

import (
	"context"
	"restApi/config"
	"restApi/pkg/logger"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})

	Rctx = context.Background()
	logger.Info(map[string]interface{}{"Rdb": Rdb})
	logger.Info(map[string]interface{}{"Rctx": Rctx})
}

func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
