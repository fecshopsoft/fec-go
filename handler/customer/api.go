package customer

import(
    //"github.com/gin-gonic/gin"
   // "github.com/fecshopsoft/fec-go/helper"
   // "errors"
    //"fmt"
)

// 得到当前的 customerParentId
/*
func GetCurrentCustomerParentId(c *gin.Context) (int64, error){
    parentId := c.GetInt64("currentCustomerParentId")
    if parentId != 0 {
        return parentId, nil
    }
    customerId := helper.GetCurrentCustomerId(c)
    customerOne, err := GetCustomerOneById(customerId)
    if err != nil{
        return 0, err
    }
    c.Set("currentCustomerParentId", customerOne.ParentId)
    return customerOne.ParentId, nil
} 
*/

/**
 * 对于3种级别的用户，在列表中，根据传入的own_id，返回合法的own_id，进行数据的过滤
 * 1.super admin, 当 own_id 为0，则返回0，当own_id不为0，则返回传递的值
 * 2.common admin，直接将当前的customer_id 作为own_id返回
 * 3.common admin  child，将parent_id返回。
 * 该函数一般用于列表数据查询where条件中对own_id的过滤
 */
 /*
func Get3OwnId(c *gin.Context, own_id int64) (int64, error){
    var currentOwnId int64
    var err error
    cType := helper.GetCurrentCustomerType(c)
    if own_id == 0 {
        if cType == helper.AdminSuperType {
            // 超级账户
            return int64(0), nil
        } else if cType == helper.AdminCommonType {
            currentOwnId = helper.GetCurrentCustomerId(c)
        } else if cType == helper.AdminChildType {
            currentOwnId, err = GetCurrentCustomerParentId(c)
            if err != nil {
                return 0, err
            }
        }
    } else {
        if cType == helper.AdminSuperType {
            // 超级账户
            return own_id, nil
        } else if cType == helper.AdminCommonType {
            currentOwnId = helper.GetCurrentCustomerId(c)
        } else if cType == helper.AdminChildType {
            currentOwnId, err = GetCurrentCustomerParentId(c)
            if err != nil {
                return 0, err
            }
        }
    }
    if currentOwnId == 0 {
        return 0, errors.New("current own id is 0")
    }
    return currentOwnId, nil
}
*/


/**
 * 得到当前用户的有效的own信息。
 * 主要用于页面渲染过程中，own_name部分的渲染。。
 */
/*
func Get3OwnNameOps(c *gin.Context) ([]helper.VueSelectOps, error){
    var ids []int64
    var groupArr []helper.VueSelectOps
    customerType := helper.GetCurrentCustomerType(c)
    customerId := helper.GetCurrentCustomerId(c)
    // 如果type == AdminSuperType，则返回所有的type = 2 的用户的own信息
    if customerType == helper.AdminSuperType {
        cCustomers, err := GetAllEnableCommonCustomer()
        if err != nil{
            return nil, err
        }
        for i:=0; i<len(cCustomers); i++ {
            commonCustomer := cCustomers[i]
            groupArr = append(groupArr, helper.VueSelectOps{Key: commonCustomer.Id, DisplayName: commonCustomer.Username})
        }
        return groupArr, nil
    }
    
    if customerType == helper.AdminChildType {
        parentId, err := GetCurrentCustomerParentId(c)
        if err != nil{
            return nil, err
        }
        ids = append(ids, parentId)
    } else if customerType == helper.AdminCommonType {
        ids = append(ids, customerId)
    }
    customers, err := GetCustomerUsernameByIds(ids)
    if err != nil{
        return nil, err
    }
    for i:=0; i<len(customers); i++ {
        customerOne := customers[i]
        groupArr = append(groupArr, helper.VueSelectOps{Key: customerOne.Id, DisplayName: customerOne.Username})
    }
    return groupArr, nil
    
}
*/


/**
 * 如果用户的级别=1，则直接返回own_id
 * 如果用户的级别=2，则直接返回其本身的id
 * 如果用户的级别=3，则返回他的parent_id
 */
 /*
func Get3SaveDataOwnId(c *gin.Context, own_id int64) (int64, error){
    // 添加创建人
    customerId := helper.GetCurrentCustomerId(c)
    customerType := helper.GetCurrentCustomerType(c)
    if customerType == helper.AdminSuperType {
        if own_id == 0 {
            return 0, errors.New("own id is empty")
        }
        return own_id, nil
    }
    if customerType == helper.AdminCommonType {
        return customerId, nil
    }
    if customerType == helper.AdminChildType {
        customer, err := GetCustomerOneById(customerId)
        if err != nil {
            return 0, err
        }
        parent_id := customer.ParentId
        if parent_id == 0 {
            return 0, errors.New("customer parent id is 0")
        }
        return parent_id, nil
    }
    return 0, errors.New("GetThirdUserOwnId error: customer account type error")
}
*/