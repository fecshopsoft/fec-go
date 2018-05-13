package mysqldb

import(
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/fecshopsoft/fec-go/config"
	"sync"
    "reflect"
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
        //if mysql_str != "" {}
        var err error
        engine, err = xorm.NewEngine("mysql", mysql_str)
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

func CloseEngine() error{
    engine := GetEngine()
    return engine.Close()
}

type XOrmWhereParam map[string]interface{}

/**
 * 根据传递的格式条件，得到字符串和相应的值，
 *    {
 *        "id": 1,   // int类型的完全匹配
 *        "username": "terry",   // 字符串类型的完全匹配
 *        "age": ["scope", 1, 9],  // 范围类型的查询  age >=1  and  age <9  
 *        "email":["like", "34@qq.com"],  // 模糊查询
 *    }
 */ 
func GetXOrmWhere(whereParam XOrmWhereParam) (string, []interface{}){
    // 组织where 语句
    var whereVal []interface{}
    var whereStr string
    for columnName, columnVal := range whereParam{
        if columnVal == nil {
            //
        } else if _, ok := columnVal.(int); ok {
            if columnVal != 0 {
                if whereStr != "" {
                    whereStr += " and " + columnName + " = ? "
                } else {
                    whereStr += columnName + " = ? "
                }
                whereVal = append(whereVal, columnVal)
            }
        } else if _, ok := columnVal.(int64); ok {
            if columnVal != 0 {
                if whereStr != "" {
                    whereStr += " and " + columnName + " = ? "
                } else {
                    whereStr += columnName + " = ? "
                }
                whereVal = append(whereVal, columnVal)
            }
        } else if _, ok := columnVal.(string); ok {
            if columnVal != "" {
                if whereStr != "" {
                    whereStr += " and " + columnName + " = ? "
                } else {
                    whereStr += columnName + " = ? "
                }
                whereVal = append(whereVal, columnVal)
            }
        } else if v := reflect.ValueOf(columnVal); v.Kind() == reflect.Slice {
            sColumnVal := columnVal.([]string)
            valLen := len(sColumnVal)
            if valLen == 3 && sColumnVal[0] == "scope" {
                if whereStr != "" {
                    if sColumnVal[1] != "" {
                        whereStr += " and " + columnName + " >= ? "
                        whereVal = append(whereVal, sColumnVal[1])
                    } 
                    if sColumnVal[2] != "" {
                        whereStr += " and " + columnName + " < ? "
                        whereVal = append(whereVal, sColumnVal[2])
                    } 
                    // and " + columnName + " < ? 
                } else {
                    // whereStr += columnName + " >= ? and " + columnName + " < ? "
                    if sColumnVal[1] != "" {
                        whereStr += columnName + " >= ? "
                        whereVal = append(whereVal, sColumnVal[1])
                    } 
                    if sColumnVal[2] != "" && whereStr == "" {
                        whereStr += columnName + " < ? "
                        whereVal = append(whereVal, sColumnVal[2])
                    } else if sColumnVal[2] != "" && whereStr != "" {
                        whereStr += " and " + columnName + " < ? "
                        whereVal = append(whereVal, sColumnVal[2])
                    }
                }
            } else if valLen == 2 && sColumnVal[0] == "like" {
                if whereStr != "" {
                    whereStr += " and " + columnName + " like ? "
                } else {
                    whereStr +=  columnName + " like ? "
                }
                whereVal = append(whereVal, "%"+sColumnVal[1]+"%")
            }
        }
    }
    return whereStr, whereVal
}
















