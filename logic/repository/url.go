package repository

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/golang-module/carbon/v2"
    "gorm.io/gorm"
    "net/http"
    "net/url"
    "short_url/config"
    "short_url/lib"
    "short_url/logic/dao"
    "short_url/logic/form"
    "short_url/logic/model"
    "strings"
    "time"
)

type shortUrl struct{}

var ShortUrl = new(shortUrl)

func (u shortUrl) Cov(c *gin.Context, req *form.CovReq) (*form.CovResp, error) {
    user, _ := Base.GetUser(c)
    url := model.ShortUrl{
        Url:  req.Url,
        User: user,
    }
    if req.ExpiredAt != "" {
        et := time.Unix(carbon.Parse(req.ExpiredAt).Timestamp(), 0)
        url.ExpiredAt = &et
    }
    resp := new(form.CovResp)
    err := dao.ShortUrlDao.DB().Transaction(func(tx *gorm.DB) error {
        err := tx.Create(&url).Error
        if err != nil {
            return err
        }
        code, err := lib.CovCode(url.ID)
        if err != nil {
            return nil
        }
        url.Code = code
        //写入短链标记
        err = tx.Model(&url).Update("code", code).Error
        if err != nil {
            return err
        }

        //写入Redis缓存
        err = lib.Redis().Set(c, url.Code, url.Url, GetTtl(url.ExpiredAt)).Err()
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    resp.Code = url.Code
    resp.NewUrl = u.GetFullPath(url.Code)
    return resp, nil
}

func GetTtl(endTime *time.Time) time.Duration {
    var ttl time.Duration
    if endTime != nil {
        ttl = time.Until(*endTime)
    } else {
        ttl = -1
    }
    return ttl
}

func (u shortUrl) Rcov(c *gin.Context, req *form.RcovReq) (*form.RcovResp, error) {
    code, err := u.GetCode(req.NewUrl)
    if err != nil {
        return nil, err
    }
    url := new(model.ShortUrl)
    dao.ShortUrlDao.DB().Where("code = ?", code).First(&url)
    data := new(form.RcovResp)
    if url.ID > 0 {
        data.Code = url.Code
        data.Url = url.Url
        if url.ExpiredAt != nil {
            data.ExpiredAt = carbon.CreateFromTimestamp(url.ExpiredAt.Unix()).ToDateTimeString()
        }
        data.NewUrl = u.GetFullPath(data.Code)
    } else {
        return nil, errors.New("NOT FOUND")
    }
    return data, nil
}

func (u shortUrl) GetFullPath(code string) string {
    return config.ServiceDomain + "/" + code
}

// GetCode 从短链中解析CODE
func (u shortUrl) GetCode(newUrl string) (string, error) {
    p, err := url.Parse(newUrl)
    if err != nil {
        return "", err
    }
    res := strings.Split(p.Path, "/")
    return res[len(res)-1], nil
}

func (u shortUrl) DeleteCov(c *gin.Context, req *form.DeleteReq) (*form.DeleteResp, error) {
    code, err := u.GetCode(req.NewUrl)
    if err != nil {
        return nil, err
    }
    user, err := Base.GetUser(c)
    if err != nil {
        return nil, err
    }
    url := u.GetUserUrl(user, code)
    if url.ID == 0 {
        return nil, errors.New("NOT FOUND")
    }
    dao.ShortUrlDao.DB().Delete(&url)
    lib.Redis().Del(c, code)
    return nil, nil
}

func (u shortUrl) UpdateCov(c *gin.Context, req *form.UpdateReq) (*form.UpdateResp, error) {
    code, err := u.GetCode(req.NewUrl)
    if err != nil {
        return nil, err
    }
    user, err := Base.GetUser(c)
    if err != nil {
        return nil, err
    }
    url := u.GetUserUrl(user, code)
    if url.ID == 0 {
        return nil, errors.New("NOT FOUND")
    }
    url.Url = req.Url
    if req.ExpiredAt != "" {
        et := time.Unix(carbon.Parse(req.ExpiredAt).Timestamp(), 0)
        url.ExpiredAt = &et
    }
    url.DB().Save(&url)
    lib.Redis().Set(c, code, url.Url, GetTtl(url.ExpiredAt))
    return nil, nil
}

func (u shortUrl) Trans(c *gin.Context, code string) (interface{}, error) {
    url := lib.Redis().Get(c, code).Val()
    if url == "" {
        return nil, errors.New("NOT FOUND")
    }
    c.Redirect(http.StatusPermanentRedirect, url)
    return nil, nil
}

func (u shortUrl) GetUserUrl(user string, code string) *model.ShortUrl {
    url := new(model.ShortUrl)
    dao.ShortUrlDao.DB().Where("user = ? and code = ?", user, code).First(&url)
    return url
}
