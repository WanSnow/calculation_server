package redis

import (
	"github.com/go-redis/redis"
	"github.com/wansnow/calculation_server/config"
	"time"
)

var RedisClient *redis.Client

func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.RedisC.Addr,
		PoolSize:     config.RedisC.PoolSize,
		MinIdleConns: config.RedisC.MinIdleConns,
		DialTimeout:  time.Duration(config.RedisC.DialTimeout * int(time.Second)),
		ReadTimeout:  time.Duration(config.RedisC.ReadTimeout * int(time.Second)),
		WriteTimeout: time.Duration(config.RedisC.WriteTimeout * int(time.Second)),
		PoolTimeout:  time.Duration(config.RedisC.PoolTimeout * int(time.Second)),
		IdleTimeout:  time.Duration(config.RedisC.IdleTimeout * int(time.Minute)),
	})
}
