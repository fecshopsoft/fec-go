package helper

import(
    "github.com/gin-gonic/gin"
)

// superAdmin
var AdminSuperType int = 1 
// superAdmin
var AdminCommonType int = 2 
// superAdmin
var AdminChildType int = 3

var VueUserRoles = map[int]string{
    AdminSuperType: "admin",
    AdminCommonType: "common_admin",
    AdminChildType: "common_admin_child",
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


