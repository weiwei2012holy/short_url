package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "short_url/lib"
    "short_url/logic/repository"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        key := c.GetHeader("X-Auth-Key")
        secret := c.GetHeader("X-Auth-Secret")
        //验证用户
        err := repository.AuthRepository.CheckAuth(key, secret)
        if err != nil {
            lib.Failed(c, err.Error(), http.StatusUnauthorized)
        }
        //写入全局变量
        c.Set("auth_key", key)

        c.Next()
        // 执行完对应的回调函数之后, 继续回到这个地方进行执行(但是响应还没有返回给客户端)
    }

    // 当中间件执行完之后, 才真正把响应返回给客户端
}
