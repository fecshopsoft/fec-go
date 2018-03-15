package handle

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/util"
    //"fmt"
)

var currentCustomer interface{}

func getHeader(c *gin.Context, key string) string{
    if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
/**
 * 通过id查询customer
 */
func PermissionAdmin(c *gin.Context){
    //c.AbortWithStatusJSON(http.StatusOK, c.Request.Header)
    access_token := getHeader(c, "X-Token");
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
    currentCustomer = data
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}





