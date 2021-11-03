package lib

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "short_url/config"
)

var MysqlClient *gorm.DB
var MysqlClientMap map[string]*gorm.DB

func init() {
    //InitDB()
}

func DB(connName ...string) *gorm.DB {
    var name string
    if len(connName) == 0 {
        name = config.MysqlConnNameDefault
    } else {
        name = connName[0]
    }
    client, exist := MysqlClientMap[name]
    if !exist {
        //。。。
    }
    return client
}

func InitDB() {
    MysqlClientMap = make(map[string]*gorm.DB)
    options := config.MysqlDefault.Config.(gorm.Config)
    db, err := gorm.Open(mysql.Open(config.MysqlDefault.Dsn), &options)
    if err != nil {
        log.Fatal(err)
    }
    MysqlClient = db
    MysqlClientMap[config.MysqlConnNameDefault] = db
}
