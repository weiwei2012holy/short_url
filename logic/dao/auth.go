package dao

import (
    "short_url/logic/model"
)

type authDao struct {
}

var AuthDao = new(authDao)

func (d authDao) GetAuthUser(key string) model.Auth {
    var auth model.Auth
    model.Auth{}.DB().Where("`key`=?", key).First(&auth)
    return auth
}
