package handle

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "errors"
    "github.com/go-xorm/xorm"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/config"
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

type VueSelectOps struct{
    Key int64 `form:"key" json:"key"`
    DisplayName string `form:"display_name" json:"display_name"`
}

var engine *(xorm.Engine)
var listDefaultPage string = "1"
var listPageCount   string = "20"
// 当前的用户
var cCustomer MapStrInterface
// 当前的用户id
var cCustomerId int64
// 当前的用户username
var cCustomerUsername string
var cCustomerType int
func init(){
    engine = mysqldb.GetEngine()
    if listDefaultPage == "" && config.Get("listDefaultPage") != "" {
        listDefaultPage = config.Get("listDefaultPage")
    }
    if listPageCount == "" && config.Get("listPageCount") != "" {
        listPageCount = config.Get("listPageCount")
    }
}

func NotFound(c *gin.Context) {
	//if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		
	//}
    c.AbortWithStatusJSON(http.StatusNotFound, util.BuildFailResult("未知资源"))
}
// 得到当前的customerId
func GetCurrentCustomerId() (int64, error){
    if cCustomerId != 0 {
        return cCustomerId, nil
    }
    if cCustomer == nil {
        var err error
        cCustomer, err = GetCurrentCustomer()
        if err != nil {
            return 0, err
        }
    }
    customerIdFloat64, ok := cCustomer["id"].(float64)
    if ok == false {
        return 0, errors.New("get current id fail")
    } 
    cCustomerId =  int64(customerIdFloat64)
    if cCustomerId == 0 {
        return 0, errors.New("get current id fail")
    }
    return cCustomerId, nil
}   
// 得到当前的customer
func GetCurrentCustomer() (MapStrInterface, error){
    var ok bool = true
    if cCustomer == nil {
        cCustomer, ok = currentCustomer.(map[string]interface{})
    }
    if ok == false {
        return nil, errors.New("var currentCustomer type convert error")
    }
    if cCustomer == nil {
        return nil, errors.New("you must relogin your account")
    }
    return cCustomer, nil
}  
// 得到当前的username
func GetCurrentCustomerUsername() (string, error){
    if cCustomerUsername != "" {
        return cCustomerUsername, nil
    }
    if cCustomer == nil {
        var err error
        cCustomer, err = GetCurrentCustomer()
        if err != nil {
            return "", err
        }
    }
    cCustomerUsername, ok := cCustomer["username"].(string)
    if ok == false {
        return "", errors.New("var cCustomer username type convert error")
    } 
    if cCustomerUsername == "" {
        return "", errors.New("get current username fail")
    }
    return cCustomerUsername, nil
}  

// 得到当前用户的type
func GetCurrentCustomerType() (int, error){
    if cCustomerType != 0 {
        return cCustomerType, nil
    }
    if cCustomer == nil {
        var err error
        cCustomer, err = GetCurrentCustomer()
        if err != nil {
            return 0, err
        }
    }
    cCustomerTypeFloat32, ok := cCustomer["type"].(float32)
    if ok == false {
        return 0, errors.New("var cCustomer type type convert error")
    } 
    if cCustomerTypeFloat32 == 0 {
        return 0, errors.New("get current username fail")
    }
    cCustomerType = int(cCustomerTypeFloat32)
    return cCustomerType, nil
}  



/**
 * 得到Request Method的ops数组
 */
func ReqMethodOps() ([]VueSelectOps){
    groupArr := []VueSelectOps{
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
    return groupArr
}

/**
 * 通过前台传递的own_id，得到当前的own_id
 * 如果当前用户type == 2,则own_id = 当前用户的customerId
 * 如果创建人的type == 1,则own_id = 前台传递的own_id，另外需要检查传递的own_id的合法性，数据库中是否存在，并且type是否 == 2
 * 其他的判定为不合法
 */
func GetCustomerOwnId(ownId int64) (int64, error){
    // 添加创建人
    customerId, err := GetCurrentCustomerId()
    if err != nil {
        return 0, err
    }
    customerType, err := GetCurrentCustomerType()
    if err != nil {
        return 0, err
    }
    if customerType == AdminCommonType {
        return customerId, nil
    }
    if customerType == AdminSuperType {
        customerOwn, err := GetCustomerOneById(ownId)
        if err != nil {
            return 0, err
        }
        if customerOwn.Type != AdminCommonType { 
            return 0, errors.New("error: own id account type error")
        }
        return ownId, nil
    }
    return 0, errors.New("you not hava role operate it")
}

// 得到当前可用的own_id数组，用于role编辑部分
// common admin账户只能选择当前用户的id
// super admin账户可以选择所有的common admin账户
func OwnIdOps() ([]VueSelectOps, error){
    var ownIdArr []VueSelectOps
    customerType, err := GetCurrentCustomerType()
    if err != nil{
        return nil, err 
    }
    customerId,err := GetCurrentCustomerId()
    if err != nil{
        return nil, err 
    }
    customerUsername,err := GetCurrentCustomerUsername()
    if err != nil{
        return nil, err 
    }
    if customerType == AdminCommonType {
        ownIdArr = append(ownIdArr, VueSelectOps{Key: customerId, DisplayName: customerUsername})
        return ownIdArr, nil
    }
    if customerType == AdminSuperType {
        customers, err := GetAllEnableCommonCustomer()
        if err != nil{
            return nil, err 
        }
        for i:=0; i<len(customers); i++ {
            customer := customers[i]
            ownIdArr = append(ownIdArr, VueSelectOps{Key: customer.Id, DisplayName: customer.Username})
        }
        return ownIdArr, nil
    }
    return nil, errors.New("you do not have role to operate own_id")
}

/*
func ReqMethodArr() {
    return map[int]string{
        1: "GET",
        2: "POST",
        3: "PATCH",
        4: "DELETE",
        5: "OPTIONS",
    }
}
*/
