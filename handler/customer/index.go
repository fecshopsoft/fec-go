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
    "github.com/fecshopsoft/fec-go/helper"
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

type VueMutilSelect map[string][]MapStrInterface

type VueSelectOps struct{
    Key int64 `form:"key" json:"key"`
    DisplayName string `form:"display_name" json:"display_name"`
}

var once sync.Once
var engine *(xorm.Engine)
var reqMethodArr []VueSelectOps
var typeArr []VueSelectOps

// init 函数在程序启动时执行，后面不会再执行。
func init(){
    engine = mysqldb.GetEngine()
    log.Println("GetEngine complete")
}

// 请求类型
var ReqMehdArr = map[int]string{
    1: "GET",
    2: "POST",
    3: "PATCH",
    4: "DELETE",
    5: "OPTIONS",
}




/**
 * 得到Request Method的ops数组
 */
func ReqMethodOps() ([]VueSelectOps){
    once.Do(func() {
        reqMethodArr = []VueSelectOps{
            VueSelectOps{
                Key: 1,
                DisplayName: ReqMehdArr[1],
            },
            VueSelectOps{
                Key: 2,
                DisplayName: ReqMehdArr[2],
            },
            VueSelectOps{
                Key: 3,
                DisplayName: ReqMehdArr[3],
            },
            VueSelectOps{
                Key: 4,
                DisplayName: ReqMehdArr[4],
            },
            VueSelectOps{
                Key: 5,
                DisplayName: ReqMehdArr[5],
            },
        }
    })
    return reqMethodArr
}

/**
 * 根据用户的级别，通过own_id字段进行数据的过滤
 */
func OwnIdQueryFilter(c *gin.Context, whereParam mysqldb.XOrmWhereParam) (mysqldb.XOrmWhereParam, error){
    customerType := helper.GetCurrentCustomerType(c)
    if customerType == helper.AdminCommonType {
        whereParam["own_id"] = helper.GetCurrentCustomerId(c)
    } else if customerType != helper.AdminSuperType {
        return  nil, errors.New("you donot have role")
    }
    return whereParam, nil
}



/**
 * 得到customer type 对应的name
 */
func GetCustomerTypeName() ([]VueSelectOps){
    once.Do(func() {
        typeArr = []VueSelectOps{
            VueSelectOps{
                Key: int64(helper.AdminSuperType),
                DisplayName: "Super Admin",
            },
            VueSelectOps{
                Key: int64(helper.AdminCommonType),
                DisplayName: "Common Admin",
            },
            VueSelectOps{
                Key: int64(helper.AdminChildType),
                DisplayName: "Child Account",
            },
        }
    })
    return typeArr
}