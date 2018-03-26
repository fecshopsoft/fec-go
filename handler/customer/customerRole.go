package customer

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "errors"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/helper"
    //"fmt"
)

type CustomerRole struct {
    Id int64 `form:"id" json:"id"`
    OwnId int64 `form:"own_id" json:"own_id" `
    CustomerId int64 `form:"customer_id" json:"customer_id" `
    RoleId int64 `form:"role_id" json:"role_id" `
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId int64 `form:"created_customer_id" json:"created_customer_id" `
}


// 查询customer 对应的所有的role信息，以及选中的信息
func CustomerRoleAllAndSelect(c *gin.Context){
    // 得到选中的role
    // var resourceChecked []int64
    customerRoles, err := GetCurrentCustomerRoleAll(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    var vueMutilSelect []MapStrInterface
    // 得到所有的resource列表
    // var allResource []MapStrInterface
    own_id := helper.GetCurrentCustomerId(c)
    roles, err := GetRolesByOwnId(own_id)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    for i:=0; i<len(roles); i++ {
        role := roles[i];
        checked := false
        for j:=0; j<len(customerRoles); j++ {
            customerRole := customerRoles[j];
            if customerRole.RoleId == role.Id {
                checked = true
            }
        }
        vueMutilSelect = append(vueMutilSelect, MapStrInterface{
            "id": role.Id,
            "name": role.Name,
            "checked": checked,
        })
        
    }
    result := util.BuildSuccessResult(gin.H{
        "allRole": vueMutilSelect,
    })
    c.JSON(http.StatusOK, result)
}



// 得到当前用户，对应的所有的权限组。一对多的关系。
func GetCurrentCustomerRoleAll(c *gin.Context) ([]CustomerRole, error){
    var customerRoles []CustomerRole
    // own_id, err := strconv.Atoi(c.Param("own_id"))
    customer_id, err := strconv.Atoi(c.DefaultQuery("customer_id", ""))
    if err != nil{
        return  customerRoles, err
    }
    if customer_id == 0{
        return  customerRoles, errors.New("customer_id is empty")
    }
    // 查看该customer对应的parent_id ， 是否是当前用户？如果不是，则不允许
    customer, err := GetCustomerOneById(int64(customer_id))
    if err != nil{
        return  customerRoles, err
    }
    parent_id := customer.ParentId
    currentCustomerId := helper.GetCurrentCustomerId(c)
    if currentCustomerId == 0 {
        return  customerRoles, errors.New("current customer id is empty")
    }
    if parent_id != currentCustomerId {
        return  customerRoles, errors.New("You do not have permission to edit user role")
    }
    // 当前用户，就是own_id
    own_id := currentCustomerId
    
    // 得到当前用户对应的父账户所有的权限列表
    err = engine.Where("customer_id = ? and own_id = ?", customer_id, own_id).Find(&customerRoles) 
    if err != nil{
        return customerRoles, err 
    }
    // log.Println( customerRoles)
    return customerRoles, nil
}

// 根据customer_id 得到该用户的role_ids
func GetRoleIdsByCustomerOwnId(own_id int64, customer_id int64) ([]int64, error){ 
    var customerRoles []CustomerRole
    var role_ids []int64
    err := engine.Where("customer_id = ? and own_id = ?", customer_id, own_id).Find(&customerRoles)
    if err != nil{
        return role_ids, err 
    }
    for i:=0; i<len(customerRoles); i++ {
        customerRole := customerRoles[i]
        role_ids = append(role_ids,customerRole.RoleId)
    }
    return role_ids, nil 
}

// 接收更新role resource的类型
type UpdateCustomerRole struct{
    CustomerId int64 `form:"customer_id" json:"customer_id" binding:"required"`
    Roles []int64 `form:"roles" json:"roles" binding:"required"`
}

// 更新 customer 对应的 roles
func CustomerRoleUpdate(c *gin.Context){
    var updateCustomerRole UpdateCustomerRole
    err := c.ShouldBindJSON(&updateCustomerRole);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    customer_id := updateCustomerRole.CustomerId
    roles := updateCustomerRole.Roles
    
    // 查看该customer对应的parent_id ， 是否是当前用户？如果不是，则不允许
    customer, err := GetCustomerOneById(int64(customer_id))
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    parent_id := customer.ParentId
    currentCustomerId := helper.GetCurrentCustomerId(c)
    if currentCustomerId == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("current customer id is empty"))
        return
    }
    if parent_id != currentCustomerId {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("You do not have permission to edit user role"))
        return
    }
    own_id := currentCustomerId
    
    // 删除 在RoleResource表中role_id 和 own_id对应的所有的资源
    var customerRole CustomerRole
    _, err = engine.Where("customer_id = ? and own_id = ? ", customer_id, own_id).Delete(&customerRole)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // 将获取的数据插入。
    for i:=0; i<len(roles); i++ {
        var cr CustomerRole
        cr.CreatedCustomerId = currentCustomerId
        cr.OwnId = own_id
        cr.CustomerId = customer_id
        cr.RoleId = roles[i]
        _, err = engine.Insert(&cr)
    }
    result := util.BuildSuccessResult(gin.H{
        "updateCustomerRole":updateCustomerRole,
    })
    c.JSON(http.StatusOK, result)
}

