package customer

import(
    // "net/http"
    "errors"
    "log"
    "sync"
    "github.com/gin-gonic/gin"
    "github.com/go-xorm/xorm"
    // "github.com/fecshopsoft/fec-go/util"
    // "github.com/fecshopsoft/fec-go/config"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
)

type DeleteIds struct{
    Ids []int `form:"ids" json:"ids"`
}

type DeleteId struct{
    Id int `form:"id" json:"id"`
}
// 三种map类型，方便使用
type MapStrInterface map[string]interface{}
type MapIntStr map[int]string
type MapStrInt map[string]int
type MapInt64Str map[int64]string
type MapStrInt64 map[string]int64

type VueSelectOps struct{
    Key int64 `form:"key" json:"key"`
    DisplayName string `form:"display_name" json:"display_name"`
}

var once sync.Once
var engine *(xorm.Engine)
var reqMethodArr []VueSelectOps

// init 函数在程序启动时执行，后面不会再执行。
func init(){
    engine = mysqldb.GetEngine()
    log.Println("GetEngine complete")
}

// 得到当前的customerId
func GetCurrentCustomerId(c *gin.Context) (int64){
    return c.GetInt64("currentCustomerId");
}  

// 得到当前的customerType
func GetCurrentCustomerType(c *gin.Context) (int){
    return c.GetInt("currentCustomerType");
} 
 
// 得到当前的customer
func GetCurrentCustomer(c *gin.Context) (MapStrInterface){
    return c.GetStringMap("currentCustomer");
}  
// 得到当前的 customerUsername
func GetCurrentCustomerUsername(c *gin.Context) (string){
    return c.GetString("currentCustomerUsername");
}  

/**
 * 得到Request Method的ops数组
 */
func ReqMethodOps() ([]VueSelectOps){
    once.Do(func() {
        reqMethodArr = []VueSelectOps{
            VueSelectOps{
                Key: 1,
                DisplayName: "GET",
            },
            VueSelectOps{
                Key: 2,
                DisplayName: "POST",
            },
            VueSelectOps{
                Key: 3,
                DisplayName: "PATCH",
            },
            VueSelectOps{
                Key: 4,
                DisplayName: "DELETE",
            },
            VueSelectOps{
                Key: 5,
                DisplayName: "OPTIONS",
            },
        }
    })
    return reqMethodArr
}

/**
 * 根据用户的级别，通过own_id字段进行数据的过滤
 */
func OwnIdQueryFilter(c *gin.Context, whereParam mysqldb.XOrmWhereParam) (mysqldb.XOrmWhereParam, error){
    customerType := GetCurrentCustomerType(c)
    if customerType == AdminCommonType {
        whereParam["own_id"] = GetCurrentCustomerId(c)
    } else if customerType != AdminSuperType {
        return  nil, errors.New("you donot have role")
    }
    return whereParam, nil
}
