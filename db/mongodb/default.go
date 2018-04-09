package mongodb
import (
    "github.com/globalsign/mgo"
    "strconv"
    "github.com/fecshopsoft/fec-go/config"
    //  "labix.org/v2/mgo/bson"
)
/**
 * 参考：http://www.cnblogs.com/shenguanpu/p/5318727.html
 * 参考：http://www.jyguagua.com/?p=3126
 */
const (
    // USER string = "user"
    // MSG  string = "msg"
)
var (
    session      *mgo.Session
    ip = "127.0.0.1"
    port = "27017"
    databaseName = "fecshop_demo"
    maxPoolSize = "10"
    poolLimit = 10
)
func init(){
    ip              = config.Get("mgo_ip")
    port            = config.Get("mgo_port")
    databaseName    = config.Get("mgo_databaseName")
    maxPoolSize     = config.Get("mgo_maxPoolSize")
    poolLimitStr   := config.Get("mgo_poolLimit")
    poolLimit, _    = strconv.Atoi(poolLimitStr)
}
func Session() *mgo.Session {
    if session == nil {
        var err error
        session, err = mgo.Dial(ip + ":" + port + "?maxPoolSize=" + maxPoolSize) 
        session.SetPoolLimit(poolLimit)
        if err != nil {
            panic(err) // no, not really
        }
    }
    return session.Clone()
}

// 可以指定collection，database使用配置中的值
func MC(collection string, f func(*mgo.Collection) error ) error {
    session := Session()
    defer func() {
        session.Close()
        // if err = recover(); err != nil {
            // Log("M", err)
        // }
    }()
    c := session.DB(databaseName).C(collection)
    // 关于return 和 defer 执行的优先级参看:https://studygolang.com/articles/4809
    return f(c) 
}

// 可以指定database和collection
func MDC(dbName string, collection string, f func(*mgo.Collection) error ) error {
    session := Session()
    defer func() {
        session.Close()
        if err := recover(); err != nil {
            // Log("M", err)
        }
    }()
    c := session.DB(dbName).C(collection)
    return f(c)
}


