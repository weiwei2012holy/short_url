package dao

import (
    "gorm.io/gorm"
    "short_url/logic/model"
)

type shortUrlDao struct {
}

var ShortUrlDao = new(shortUrlDao)

func (d *shortUrlDao) DB() *gorm.DB {
    return model.ShortUrl{}.DB()
}
