package lib

/**
 * Note: ID和短链转换
 * Author: 余伟<weiwei2012holy@hotmail.com>
 * Time: 2021/10/29
 */

import (
    "errors"
    "math"
)

const Scale = 62

// CovCode 通过ID生成短链标识
func CovCode(number uint64) (string, error) {
    dic := IndexMap()
    if len(dic) != Scale {
        return "", errors.New("字典配置有误")
    }
    tran := Ten2Any(number, Scale)
    var res string
    for i := 0; i < len(tran); i++ {
        res += dic[tran[i]]
    }
    return res, nil
}

// RcovCode 解析短链，获取ID
func RcovCode(code string) (uint64, error) {
    return Any2Ten(code, Scale)
}

// Ten2Any 转换
func Ten2Any(number uint64, scale uint64) []uint64 {
    var a uint64
    var b uint64
    var res []uint64
    a = number / scale
    b = number % scale
    if a > 0 {
        res = append(res, Ten2Any(a, scale)...)
    }
    res = append(res, b)
    return res
}

// Any2Ten 解码
func Any2Ten(str string, scale uint64) (uint64, error) {

    runeStr := []rune(str)
    dic := IndexMap()

    if int(scale) > len(dic) {
        return 0, errors.New("字典配置有误")
    }
    revDicMap := make(map[string]int)
    for i, s := range dic {
        revDicMap[s] = i
    }
    var res uint64
    for i := 0; i < len(runeStr); i++ {
        v := revDicMap[string(runeStr[i])]
        fv := math.Pow(float64(scale), float64(len(runeStr)-i-1)) * float64(v)
        res += uint64(fv)
    }
    return res, nil
}

// IndexMap 进制转换字典，不要改动
func IndexMap() [62]string {
    i := [...]string{
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "a",
        "b",
        "c",
        "d",
        "e",
        "f",
        "g",
        "h",
        "i",
        "j",
        "k",
        "l",
        "m",
        "n",
        "o",
        "p",
        "q",
        "r",
        "s",
        "t",
        "u",
        "v",
        "w",
        "x",
        "y",
        "z",
        "A",
        "B",
        "C",
        "D",
        "E",
        "F",
        "G",
        "H",
        "I",
        "J",
        "K",
        "L",
        "M",
        "N",
        "O",
        "P",
        "Q",
        "R",
        "S",
        "T",
        "U",
        "V",
        "W",
        "X",
        "Y",
        "Z",
    }
    return i
}
