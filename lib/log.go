package lib

import (
    "fmt"
    "github.com/sirupsen/logrus"
    "os"
    "path"
    "time"
)

func Logger() *logrus.Logger {
    now := time.Now()
    logFilePath := ""
    if dir, err := os.Getwd(); err == nil {
        logFilePath = dir + "/logs/"
    }
    if err := os.MkdirAll(logFilePath, 0777); err != nil {
        fmt.Println(err.Error())
    }
    logFileName := now.Format("2006-01-02") + ".log"
    //日志文件
    fileName := path.Join(logFilePath, logFileName)
    if _, err := os.Stat(fileName); err != nil {
        if _, err := os.Create(fileName); err != nil {
            fmt.Println(err.Error())
        }
    }
    //写入文件
    src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        fmt.Println("err", err)
    }

    //实例化
    logger := logrus.New()

    //设置输出
    logger.Out = src

    //设置日志级别
    logger.SetLevel(logrus.DebugLevel)

    //设置日志格式
    logger.SetFormatter(&logrus.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })
    return logger
}
