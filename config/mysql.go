package config

import (
    "gorm.io/gorm"
)

type mysqlConfig struct {
    Name   string
    Dsn    string
    Config interface{}
}

const (
    MysqlConnNameDefault = "default"
)

var (

    // MysqlDefault 默认配置
    MysqlDefault = mysqlConfig{
        Name:   MysqlConnNameDefault,
        Dsn:    "",
        Config: gorm.Config{},
    }
)
