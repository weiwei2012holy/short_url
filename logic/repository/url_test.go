package repository

import (
    "context"
    "fmt"
    "github.com/stretchr/testify/assert"
    "short_url/lib"
    "testing"

)

func TestShortUrl_GetCode(t *testing.T) {

    code := "abcd"
    domain := "https://127.0.0.1"

    list := []string{
        domain + "/" + code,
        domain + "/asda/asd/" + code,
        domain + "/asda/asd/" + code+"?",
        domain + "/asda/asd/" + code+"?id=123",
        domain + "/asda/asd/" + code+"?idasd/2a",
        domain + "/asda/asd/" + code+"#aaa",
    }
    for _, s := range list {
        res, err := ShortUrl.GetCode(s)
        assert.NoError(t, err)
        assert.Equal(t, res, code)
    }

}

func TestShortUrl_Redis(t *testing.T) {

    ctx := context.Background()
    res := lib.Redis().Get(ctx,"j").Val()

    fmt.Println(res)


}
