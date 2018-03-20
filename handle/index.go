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
type MapStrInterface map[string]interface{}

var engine *(xorm.Engine)
var listDefaultPage string = "1"
var listPageCount   string = "20"
// 当前的用户
var cCustomer MapStrInterface
// 当前的用户id
var cCustomerId int
// 当前的用户username
var cCustomerUsername string
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
func GetCurrentCustomerId() (int, error){
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
    cCustomerId =  int(customerIdFloat64)
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



