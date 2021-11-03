package controller

import (
    "github.com/gin-gonic/gin"
    "short_url/lib"
    "short_url/logic/form"
    "short_url/logic/repository"
)

type url struct {
}

var Url = new(url)

// Trans 转发
func (u *url) Trans(c *gin.Context) {
    code := c.Param("code")
    if code == "" {
        lib.Failed(c, "无效的地址")
    }
    res, err := repository.ShortUrl.Trans(c,code)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    lib.Success(c, res)
}

// Cov 短链->完整链接
func (u *url) Cov(c *gin.Context) {
    form := new(form.CovReq)
    err := c.ShouldBind(&form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    res, err := repository.ShortUrl.Cov(c,form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    lib.Success(c, res)
}

// Rcov 链接->短链
func (u *url) Rcov(c *gin.Context) {
    form := new(form.RcovReq)
    err := c.ShouldBind(&form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    res, err := repository.ShortUrl.Rcov(c,form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    lib.Success(c, res)
}

func (u *url) DeleteCov(c *gin.Context) {
    form := new(form.DeleteReq)
    err := c.ShouldBind(&form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    res, err := repository.ShortUrl.DeleteCov(c, form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    lib.Success(c, res)
}

func (u *url) UpdateCov(c *gin.Context) {
    form := new(form.UpdateReq)
    err := c.ShouldBind(&form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    res, err := repository.ShortUrl.UpdateCov(c, form)
    if err != nil {
        lib.Failed(c, err.Error())
    }
    lib.Success(c, res)
}
