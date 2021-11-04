package config

import (
    "fmt"
    "github.com/spf13/viper"
    "log"
)

func InitLocal() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".") //如果有多个目录，这里可以多次添加
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    if viper.GetString("service_domain") != "" {
        ServiceDomain = viper.GetString("service_domain")
    }
    if viper.GetString("service_port") != "" {
        ServicePort = viper.GetString("service_port")
    }
    if viper.GetString("service_env") != "" {
        ServiceEnv = viper.GetString("service_env")
    }
    //mysql设置
    if viper.GetString("mysql_dsn") == "" {
        log.Fatal("config.json中{mysql_dsn}缺失")
    }
    MysqlDefault.Dsn = viper.GetString("mysql_dsn")
    //redis设置
    if viper.GetString("redis_addr") == "" {
        log.Fatal("config.json中{redis_addr}缺失")
    }
    RedisDefault.Addr = viper.GetString("redis_addr")
}
