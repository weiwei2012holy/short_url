package lib

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCov(t *testing.T) {

    var num, scale uint64
    num = 12345
    scale = 2

    res := Ten2Any(num, scale)

    fmt.Println(res)

}

func TestBuildIndex(t *testing.T) {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    runeStr := []rune(str)
    for _, s := range runeStr {
        fmt.Println("\"" + string(s) + "\",")
    }
}

func TestGetCode(t *testing.T) {
    conMap := map[uint64]string{
        123456789:  "8m0Kx",
        10000000:   "FXsk",
        3847632:    "g8WA",
        0:          "0",
        63:         "11",
        1000:       "g8",
        3463876124: "3Mq4pK",
        9999999999: "aUKYOz",
    }
    for num, code := range conMap {
        res, err := CovCode(num)
        t.Log("CovCode=>", num, res)
        assert.NoError(t, err)
        assert.Equal(t, code, res)

        res2, err2 := RcovCode(res)

        t.Log("RcovCode=>", res2, num)
        assert.NoError(t, err2)
        assert.Equal(t, num, res2)
    }
}
