package middleware

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/helper"
    //"fmt"
)
/**
 * 验证用户是否登录
 * 如果登录成功，将 currentCustomer, currentCustomerId, currentCustomerType, currentCustomerUsername 添加到上下文
 */
func PermissionLoginToken(c *gin.Context){
    //c.AbortWithStatusJSON(http.StatusOK, c.Request.Header)
    access_token := helper.GetHeader(c, "X-Token");
    if  access_token == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("access_token can not empty"))
        return
	}
	data, logined, expired, err := security.JwtParse(access_token);
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    /*
    c.AbortWithStatusJSON(http.StatusOK,gin.H{
        "data":data,
        "logined":logined,
        "expired":expired,
    })
    */
    now := time.Now().Unix()
    if logined != 1 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("用户未登录，请先登录"))
        return
    }
    if expired < now {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("token 已经过期，您需要重新登录"))
        return
    }
    
    // 设置currentCustomer
    if data == nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current customer fail"))
        return
    }
    c.Set("currentCustomer", data)
    customer := c.GetStringMap("currentCustomer");
    
    // 设置currentCustomerId
    customerIdFloat64, ok := customer["id"].(float64)
    if ok == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current customer id fail"))
        return
    } 
    customerId :=  int64(customerIdFloat64)
    if customerId == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current customer id fail"))
        return
    }
    c.Set("currentCustomerId", customerId)
    
    // 设置currentCustomerUsername
    customerUsername, ok := customer["username"].(string)
    if ok == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current customer username fail"))
        return
    } 
    if customerUsername == "" {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current customer username fail"))
        return
    }
    c.Set("currentCustomerUsername", customerUsername)
    
    // 设置currentCustomerType
    customerTypeFloat64, ok := customer["type"].(float64)
    if ok == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current customer type fail"))
        return
    } 
    customerType :=  int(customerTypeFloat64)
    if customerType == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("get current id fail"))
        return
    }
    c.Set("currentCustomerType", customerType)
    //currentCustomer = data
}




// 验证登录用户，是否有权限访问当前的资源

func PermissionRole(c *gin.Context){
    /*
    cCustomer, ok := currentCustomer.(map[string]interface{})
    if ok == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("you must relogin your account"))
        return
    }
    username := cCustomer["username"].(string)
    customerType := cCustomer["type"].(float64)
    parentId := cCustomer["parent_id"].(float64)
    */
    /*
    r := c.Request
    // url path
    path := r.URL.Path
    // url request method
    requestMethod := r.Method
    */
    // 验证当前访问的url，是否存在访问权限。
    // c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("you do not have permission to visit this url"))
    // return
    
}

// cors中间件，跨域请求加入相关参数。
/*
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, PUT, DELETE")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}
*/