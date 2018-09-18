# fec-go
fec-go


### 安装

1.库包安装

1.1安装xorm

```
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
```

2.安装gin

```
go get github.com/gin-gonic/gin
```

参考资料： [centos6 安装go框架gin的步骤，以及中间遇到的坑](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85go%e6%a1%86%e6%9e%b6gin%e7%9a%84%e6%ad%a5%e9%aa%a4%ef%bc%8c%e4%bb%a5%e5%8f%8a%e4%b8%ad%e9%97%b4%e9%81%87%e5%88%b0%e7%9a%84%e5%9d%91/)

上面如果都安装成功了

```
go get github.com/fecshopsoft/fec-go
```

下载完成后，即可下载玩所有的文件



### 配置



1.将 config/config.ini 的内容，复制到 `/etc/fec-go/config.ini`
,上面的配置文件中配置相应的参数

新建文件：`/etc/fec-go/config.ini`

内容填写如下：

```
redis_user =
redis_password =
redis_port = 127.0.0.1:2183

# mysql
mysql_user      = root
mysql_password  = trew4ffr
mysql_host      = 127.0.0.1
mysql_port      = 3306
mysql_db        = fec-go
charset         = utf8
autocommit      = true
maxOpenConns    = 2
maxIdleConns    = 2

#mongodb
mgo_ip              = 127.0.0.1
mgo_port            = 27017
mgo_databaseName    = fec_go_trace_info
mgo_maxPoolSize     = 4
mgo_poolLimit       = 4

#elastic
elastic_host = http://elasticsearch1:9200

//release,debug,test
#log_mode = release
output_log = false
router_info_log = /www/web_logs/fec-go/router_info.log
router_error_log = /www/web_logs/fec-go/router_error.log
global_log = /www/web_logs/fec-go/global.log


#shell 日志

shell_output_log = false
shell_router_info_log = /www/web_logs/fec-go-shell/router_info.log
shell_router_error_log = /www/web_logs/fec-go-shell/router_error.log
shell_global_log = /www/web_logs/fec-go-shell/global.log


#
saveUploadFileDir = /www/test/xlsx


#用户信息

#userName = terry
#userEmail = 2358269014@qq.com
#userTelephone = 18620432962
#userToken = xxxxxxxxxxxxxxxxxx
#userUrl = http://120.24.37.249:3001/verify

#http服务监听的端口
httpHost = 0.0.0.0:3000

```

上面出现的log文件，自行新建


在src/main 下面创建   fec-go.go  和 fec-go-verify.go 文件

fec-go.go 的内容如下：

```
package main
/**
 * 服务端入口部分
 * 1.初始化log输出文件
 * 2.监听ip
 *
 */
import(
    "github.com/fecshopsoft/fec-go/router"
    "github.com/fecshopsoft/fec-go/initialization"
    "github.com/fecshopsoft/fec-go/config"
    "log"
    "time"
    "os"
    "os/signal"
    "syscall"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
)

func main() {
    // 初始化log输出，log.Println("---") 输出的内容将输出到globalLog文件里面
    log.Println("------start：" + time.Now().String())
        initialization.InitGlobalLog()
        log.Println("------start：" + time.Now().String())
    log.SetFlags(log.LstdFlags | log.Llongfile)
    SetupCloseHandler()
    listenIp := config.Get("httpHost")
    router.Listen(listenIp);

}

func VerifyRemote(){

}

func SetupCloseHandler() {
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        err := mysqldb.CloseEngine()
        if err != nil {
            log.Println(err.Error())
        }
        log.Println("\r- Ctrl+C pressed in Terminal")
        log.Println("------close：" + time.Now().String())
        os.Exit(0)
    }()
}

```


fec-go-verify.go的内容如下：


```
package main

import(
    "log"
    "time"
    "errors"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    "github.com/fecshopsoft/fec-go/helper"

)

func main() {
    // 初始化log输出，log.Println("---") 输出的内容将输出到globalLog文件里面
    log.Println("------start：" + time.Now().String())

        log.Println("------start：" + time.Now().String())
    log.SetFlags(log.LstdFlags | log.Llongfile)
    listenIp := "0.0.0.0:3001"
    Listen(listenIp);

}

func Listen(listenIp string) {
    // log.Println("------444：" + time.Now().String())
    r := gin.Default()
    r.GET("/verify", f)

    r.Run(listenIp) // 这里改成您的ip和端口
}


func f(c *gin.Context){
    //access_token, _ := security.JwtSignAccessToken("xxxxxxxxxxxxxx")
    // 通过参数，获取jwt的值，解密，获取 telephone, email, token.
    // 1.按照tel查询 2.验证电话，email，token，不对就报错 3.验证时间戳
    // 返回jwt 加密的格式
    tk := c.DefaultQuery("tk", "")
    log.Println(tk)
    data, _, _, err := security.JwtParse(tk)
    log.Println(data)
    // 类型断言， 解析出来
    decodeData, ok := data.(map[string]interface{})
    if !ok {
        log.Println(errors.New("decode data error"))
    }
    log.Println(decodeData["email"])
    log.Println(decodeData["name"])
    log.Println(decodeData["telephone"])
    log.Println(decodeData["time"])
    log.Println(decodeData["token"])


    // 得到里面的各个值
    email, _ := decodeData["email"].(string)
    name, _ := decodeData["name"].(string)
    telephone, _ := decodeData["telephone"].(string)
    timess, _ := decodeData["time"].(string)
    token, _ := decodeData["token"].(string)
    timestamps := helper.DateTimestamps()

    // 判断时间间隔不要超过5分钟否则，将报错
    timeInt64, _ := helper.Int64(timess)
    if timestamps - timeInt64 > 300 {
        s := gin.H{
            "status": "fail",
            "error_info": "time error",
        }
        log.Println(s)
        enS := getJwtSign(s)
        c.String(200, enS)
        return
        
    }
    // 查询，得到user数据,判断是否过期，然后得到数据。
    user, err := GetOneByEmail(email, name, telephone, token)
    if err != nil {
        s := gin.H{
            "status": "fail",
            "error_info": err.Error(),
        }
        log.Println(s)
        enS := getJwtSign(s)
        c.String(200, enS)
        return
    }

    s := gin.H{
        "status": "success",
        "pv_count": user.PvCount,
        "site_count": user.SiteCount,
        "timestamps":timestamps,
        "error_info": "",
    }
    log.Println(s)
    enS := getJwtSign(s)
    c.String(200, enS)

}

func getJwtSign(s gin.H) string{
    b, _ := json.Marshal(s)
    sr, _ :=  security.JwtSignAccessToken(string(b[:]))
    return sr
}

  type UserInfo struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    Telephone string `form:"telephone" json:"telephone"`
    Email string `form:"email" json:"email"`
    SiteCount int64 `form:"site_count" json:"site_count"`
    PvCount int64 `form:"pv_count" json:"pv_count"`
    EndDate string `form:"end_date" json:"end_date"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    Token string `form:"token" json:"token"`
}

/**
 * 通过id查询一条记录
 */
func GetOneByEmail(email string, name string, telephone string, token string) (UserInfo, error){
    engine := mysqldb.GetEngine()
    var user UserInfo
    var user1 UserInfo
    has, err := engine.Where("email = ? and name = ? and telephone = ? and token = ? ", email, name, telephone, token).Get(&user)
    if err != nil {
        return user, err
    }
    if has == false {
        return user, errors.New("get userInfo by email error, empty data.")
    }
    // 查看时间是否过期

    t := helper.GetTimestampsByDate(user.EndDate)
    timestamps := helper.DateTimestamps()
    if t < timestamps {
        return user1, errors.New("website Time has expired")
    }

    return user, nil
}


      
        



```


2.数据处理脚本


```
go run fec-go-shell.go 1 removeEsAllIndex

```

第一个参数：`1`,代表处理N天内的数据统计

第二个参数： `removeEsAllIndex`,如果删除elasticSearch里面的数据，加上这个参数值。


3.cron

```
1 1 * * * /usr/bin/wget http://120.24.37.249:3000/fec/trace/cronssss > /dev/null
01 * * * *  /root/go/src/main/fec-go-shell   >> /www/web_logs/fec-go-shell.log  2>&1
```


第一个是更新数据，将ip和端口换成你自己的即可

第二个是周期跑数据的脚本，一天跑一次。，





