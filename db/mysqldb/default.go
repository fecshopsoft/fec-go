package mysqldb

import(
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/fecshopsoft/fec-go/config"
	"sync"
    "strconv"
)

var once sync.Once
var engine *(xorm.Engine)

func GetEngine() *(xorm.Engine){
    once.Do(func() {
		// 用于设置最大打开的连接数
        maxOpenConns := config.Get("maxOpenConns")
        // 用于设置闲置的连接数
        maxIdleConns := config.Get("maxIdleConns")
        mysql_user := config.Get("mysql_user")
        mysql_password := config.Get("mysql_password")
        mysql_host := config.Get("mysql_host")
        mysql_port := config.Get("mysql_port")
        mysql_db   := config.Get("mysql_db")
        charset    := config.Get("charset")
        autocommit := config.Get("autocommit")
        // engine, err := xorm.NewEngine("mysql", "root:Zhaoy34ggsd@tcp(127.0.0.1:3306)/fec-go?charset=utf8&autocommit=true")
        mysql_str := mysql_user + ":" + mysql_password + "@tcp(" + mysql_host + ":" + mysql_port + ")/" + mysql_db + "?charset=" + charset + "&autocommit=" + autocommit
        if mysql_str != "" {}
        var err error
        engine, err = xorm.NewEngine("mysql", "root:Zhaoyong2017fdsfds3f3GDs3fgsd@tcp(127.0.0.1:3306)/fec-go?charset=utf8&autocommit=true")
        if err != nil {
            panic(err.Error())
        }
        moc, _ := strconv.Atoi(maxOpenConns)
        mic, _ := strconv.Atoi(maxIdleConns)
        engine.SetMaxOpenConns(int(moc))
        engine.SetMaxIdleConns(int(mic))
	})
	return engine;
}