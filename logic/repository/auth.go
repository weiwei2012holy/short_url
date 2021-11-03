package repository

import (
    "errors"
    "short_url/logic/dao"
)

type authRepository struct{}

var AuthRepository = new(authRepository)

func (r authRepository) CheckAuth(key string, secret string) error {
    if key == "" {
        return errors.New("Auth Key Is Missing ")
    }
    //认证用户放到数据库，也可以
    auth := dao.AuthDao.GetAuthUser(key)
    if auth.ID == 0 {
        return errors.New("Auth Key Is Illegal ")
    }
    if auth.Secret != secret {
        return errors.New("Auth Failed ")
    }
    return nil
}
