package repository

import (
    "errors"
    "github.com/gin-gonic/gin"
)

type base struct {
}
var Base = new(base)

// GetUser 获取用户名称
func (b base) GetUser(c *gin.Context) (string, error) {
    user, exists := c.Get("auth_key")
    if !exists {
        return "", errors.New("用户不存在")
    }
    return user.(string), nil
}
