package model

import "gorm.io/gorm"

type BaseModel interface {
    DB() *gorm.DB
    ConnName() string
}
