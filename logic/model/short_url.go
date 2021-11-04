package model

import (
    "gorm.io/gorm"
    "short_url/config"
    "short_url/lib"
    "time"
)

// 短网址
type ShortUrl struct {
    ID        uint64         `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
    Code      string         `gorm:"type:char(6);column:code;NOT NULL;index:idx_code;comment:短链标识" json:"code"` // 短链标识
    Url       string         `gorm:"type:varchar(1000);column:url;NOT NULL;comment:跳转地址" json:"url"`            // 跳转地址
    User      string         `gorm:"type:varchar(30);column:user;NOT NULL;comment:跳转地址" json:"user"`            // 用户标识
    ExpiredAt *time.Time     `gorm:"type:timestamp NULL;column:expired_at" json:"expired_at"`                   // 过期时间
    DeletedAt gorm.DeletedAt `gorm:"type:timestamp NULL;column:deleted_at" json:"deleted_at"`
    CreatedAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`
    UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0);" json:"updated_at"`
}

func (m *ShortUrl) TableName() string {
    return "short_url"
}

func (m ShortUrl) DB() *gorm.DB {
    return lib.DB(m.ConnName())
}

func (m ShortUrl) ConnName() string {
    return config.MysqlConnNameDefault
}
