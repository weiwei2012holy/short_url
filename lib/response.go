package lib

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "成功",
        "data":    data,
    })
    c.Abort()
}

func Failed(c *gin.Context, message string, code ...int) {
    var newCode int
    if len(code) >= 1 {
        newCode = code[0]
    } else {
        newCode = http.StatusBadRequest
    }
    c.JSON(http.StatusBadRequest, gin.H{
        "code":    newCode,
        "message": message,
        "data":    nil,
    })
    c.Abort()
}
