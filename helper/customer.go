package helper

import(
    "github.com/gin-gonic/gin"
)

// admin
var AdminType int = 1 
// common
var CommonType int = 2

var VueUserRoles = map[int]string{
    AdminType: "super_admin",
    CommonType: "common_admin",
}


// 得到当前的customerId
func GetCurrentCustomerId(c *gin.Context) (int64){
    return c.GetInt64("currentCustomerId");
}  

// 得到当前的customerType
func GetCurrentCustomerType(c *gin.Context) (int){
    return c.GetInt("currentCustomerType");
} 

func IsAdmin(c *gin.Context) bool{
    customerType := GetCurrentCustomerType(c)
    if customerType == AdminType {
        return true
    } else {
        return false
    }
}
 
// 得到当前的customer
func GetCurrentCustomer(c *gin.Context) (MapStrInterface){
    return c.GetStringMap("currentCustomer");
}  
// 得到当前的 customerUsername
func GetCurrentCustomerUsername(c *gin.Context) (string){
    return c.GetString("currentCustomerUsername");
}  


