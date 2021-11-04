package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "short_url/lib"
)

func Visit() gin.HandlerFunc {
    return func(c *gin.Context) {
        lib.Logger().Debug(c.Request.Header)
        fmt.Println(c.Request.Header)
        c.Next()
    }

}
