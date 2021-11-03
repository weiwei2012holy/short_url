package model

import (
    "gorm.io/gorm"
    "short_url/config"
    "short_url/lib"
)

type Auth struct {
    ID     uint64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
    Key    string `gorm:"column:key;varchar(30);index:idx_key,unique;comment:授权key"`
    Secret string `gorm:"column:secret;varchar(100);comment:密钥"`
}

func (m *Auth) TableName() string {
    return "auth"
}

func (m Auth) DB() *gorm.DB {
    return lib.DB(m.ConnName())
}

func (m Auth) ConnName() string {
    return config.MysqlConnNameDefault
}
