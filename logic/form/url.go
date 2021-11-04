package form

type (
    CovReq struct {
        Url       string `json:"url" query:"url" form:"url" binding:"required"`
        ExpiredAt string `json:"expired_at" query:"expired_at" form:"expired_at"`
    }
    CovResp struct {
        Code   string `json:"code"`
        NewUrl string `json:"new_url"`
    }
)

type (
    RcovReq struct {
        NewUrl string `json:"new_url" query:"new_url" form:"new_url" binding:"required"`
    }
    RcovResp RcovData
    RcovData struct {
        Code      string `json:"code"`
        Url       string `json:"url"`
        ExpiredAt string `json:"expired_at"`
        NewUrl    string `json:"new_url"`
    }
)

type (
    UpdateReq struct {
        NewUrl    string `json:"new_url" query:"new_url" form:"new_url" binding:"required"`
        Url       string `json:"url" query:"url" form:"url" binding:"required"`
        ExpiredAt string `json:"expired_at" query:"expired_at" form:"expired_at"`
    }
    UpdateResp struct {
    }
)

type (
    DeleteReq struct {
        NewUrl string `json:"new_url" form:"new_url" query:"new_url" binding:"required"`
    }
    DeleteResp struct {
    }
)
