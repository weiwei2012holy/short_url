package lib

import (
    "github.com/go-redis/redis/v8"
    "log"
    "short_url/config"
)

var RedisClientMap map[string]*redis.Client

func init() {
    //InitRedis()
}

func Redis(connName ...string) *redis.Client {
    var name string
    if len(connName) == 0 {
        name = config.RedisConnNameDefault
    } else {
        name = connName[0]
    }
    client, exist := RedisClientMap[name]
    if !exist {
        //。。。
    }
    return client
}

func InitRedis() {
    RedisClientMap = make(map[string]*redis.Client)
    options := config.RedisDefault.Config.(redis.Options)
    options.Addr = config.RedisDefault.Addr
    if options.Addr == "" {
        log.Fatal("redis 配置不完整")
    }
    RedisClientMap[config.RedisDefault.Name] = redis.NewClient(&options)
}
