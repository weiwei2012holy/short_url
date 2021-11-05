package main

import (
    "short_url/config"
    "short_url/lib"
    "short_url/logic/model"
)

func main() {
    config.InitLocal()
    //初始化数据库
    lib.InitDB()

    lib.MysqlClient.AutoMigrate(&model.ShortUrl{})
    lib.MysqlClient.AutoMigrate(&model.Auth{})

    //testAuth := model.Auth{
    //    Key:    "test",
    //    Secret: "123456",
    //}
    //model.Auth{}.DB().Where(model.Auth{Key: "test"}).FirstOrCreate(&testAuth)
}
