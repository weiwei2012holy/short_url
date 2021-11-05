package middleware

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "short_url/lib"
)

func Visit() gin.HandlerFunc {
    return func(c *gin.Context) {
        code := c.Param("code")
        h := c.Request.Header
        str, _ := json.Marshal(h)
        lib.Logger().Infof("code=%s user=%s", code, lib.Md5(string(str)))
        //lib.Logger().Debug(c.Request.Header)
        c.Next()
    }

}
