package middleware
/**
 * 该部分的验证，必须在 PermissionLoginToken 验证后，才能使用
 * 根据用户的type等级，来决定用户是否有权限，如果没有权限则中止
 */
import(
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/fecshopsoft/fec-go/util"
    customerH "github.com/fecshopsoft/fec-go/handler/customer"
)

/**
 * 【超级admin】权限验证，只有【超级admin】账户才可以通过
 */
func SuperAdminRole(c *gin.Context){
    customerType := helper.GetCurrentCustomerType(c)
    if customerType != helper.AdminSuperType {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
        return
    }
}

/**
 * 【普通admin】权限验证，【超级admin】和【普通admin】都可以通过
 */
func CommonAdminRole(c *gin.Context){
    customerType := helper.GetCurrentCustomerType(c)
    if customerType != helper.AdminSuperType && customerType != helper.AdminCommonType {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
        return
    }
}

/**
 * 【普通admin子账户】权限验证，【超级admin】，【普通admin】，都可以通过
 * 【普通admin子账户】需要进行数据库查询权限认证。
 */
func CommonAdminChildRole(c *gin.Context){
    // 当前用户信息
    customer_id := helper.GetCurrentCustomerId(c)
    resources, ok := customerResourceCache.Get(customer_id)
    if ok == false {
        // 如果是 【超级admin】和【普通admin】，直接返回
        customerType := helper.GetCurrentCustomerType(c)
        if customerType == helper.AdminSuperType || customerType == helper.AdminCommonType {
            return
        }
        // 如果还不是 AdminChildType ， 则说明该账户的类型有问题。
        if customerType != helper.AdminChildType {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
            return
        }
        customer, err := customerH.GetCustomerOneById(customer_id)
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        parent_id := customer.ParentId
        own_id := parent_id
        // 根据用户的customer_id 得到用户的role_ids
        
        role_ids, err := customerH.GetRoleIdsByCustomerOwnId(own_id, customer_id)
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        // 根据用户的role_ids ， 得到用户的可用的resource_ids
        resource_ids, err := customerH.GetResourceIdsByRoleOwnIds(own_id, role_ids)
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        // 根据用户的resource_ids，得到想用的resources
        resources, err = customerH.GetResourcesByIds(resource_ids)
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        // 将 resources 保存到channels里面。
        customerResourceCache.Set(customer_id, resources)
    }
    // 计算权限
    /**
     * 得到当前的URI PATH 和 RequestMethod
     * url_path := r.URL.Path   // "path":"/v1/customer/list",
     * request_method := r.Method // "request_method":"GET",
     */
    r := c.Request
    role_access := false
    if r.URL.Path != "" && r.Method != "" {  // 如果不为空
        for i:=0; i<len(resources); i++ {
            resource := resources[i]
            ReqMehdStr := customerH.ReqMehdArr[resource.RequestMethod]
            if resource.UrlKey == r.URL.Path && ReqMehdStr == r.Method {
                role_access = true
                break
            }
        }
    }
    //c.AbortWithStatusJSON(http.StatusOK, gin.H{
    //    "rUrlPath": r.URL.Path,
    //    "rMethod": r.Method,
    //})
    //return
    // 没有权限则返回，不允许访问。
    if role_access == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
        return
    }
    
    return 
}

