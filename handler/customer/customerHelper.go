package customer

import(
    "github.com/gin-gonic/gin"
    "errors"
)

/**
 * 通过前台传递的own_id，得到合法的own_id
 * 如果当前用户type == 2,则own_id = 当前用户的customerId
 * 如果创建人的type == 1,则own_id = 前台传递的own_id，另外需要检查传递的own_id的合法性，数据库中是否存在，并且type是否 == 2
 * 其他的判定为不合法
 */
func GetCustomerOwnId(c *gin.Context, ownId int64) (int64, error){
    // 添加创建人
    customerId := GetCurrentCustomerId(c)
    customerType := GetCurrentCustomerType(c)
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
func GetCustomerOwnIdOps(c *gin.Context) ([]VueSelectOps, error){
    var ownIdArr []VueSelectOps
    customerType := GetCurrentCustomerType(c)
    customerId := GetCurrentCustomerId(c)
    customerUsername := GetCurrentCustomerUsername(c)
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

