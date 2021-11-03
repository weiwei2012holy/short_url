package config

import (
    "github.com/go-redis/redis/v8"
)

const (
    RedisConnNameDefault = "default"
)

type redisConfig struct {
    Name   string
    Config interface{}
}

var (
    RedisDefaultAddr = ""
    // RedisDefault 默认配置
    RedisDefault = redisConfig{
        Name: RedisConnNameDefault,
        Config: redis.Options{
            Addr: RedisDefaultAddr,
            DB:   0,
        },
    }
)
