package dao

import "gorm.io/gorm"

type BaseDao interface {
    DB() *gorm.DB
}

