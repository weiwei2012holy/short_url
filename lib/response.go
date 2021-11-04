package lib

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Success(c *gin.Context, data interface{}) {
    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "成功",
        "data":    data,
    })
}

func Failed(c *gin.Context, message string, code ...int) {
    var newCode int
    if len(code) >= 1 {
        newCode = code[0]
    } else {
        newCode = http.StatusBadRequest
    }
    c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
        "code":    newCode,
        "message": message,
        "data":    nil,
    })
}
