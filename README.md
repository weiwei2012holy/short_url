## GO版本短网址服务

### 相关依赖

1. 基础框架：GIN(https://github.com/gin-gonic/gin)
2. 数据库：GORM(https://gorm.io/zh_CN/docs/)
3. Redis：https://github.com/go-redis/redis
4. 配置管理：https://github.com/spf13/viper
5. 日期处理：Carbon

### 部署方法

1. Clone 本项目
2. 修改配置`cp config.json.tmp config.json`
3. 启动服务`go run main/service`

### 使用方法

1. 联系管理员获取Key+Secret，注意每个Key的数据相互隔离，数据存在`auth`数据表中
2. 通过自定义Header传入验证参数（服务端调用，简单验证）
    1. `X-Auth-Key` = Key
    2. `X-Auth-Secret` = Secret 
    3. 3.调用接口：POST `127.0.0.1:8090/api/url` ,生成短链,访问短链即可跳转
3. 服务地址
   1. 测试：https://test-short-url.atido.com
   2. 线上：

### 提供的接口

详见：test目录下 `short_url.http`