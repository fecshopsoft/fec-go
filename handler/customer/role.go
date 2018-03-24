package customer

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "log"
    // "errors"
    // "time"
    // "log"
    "unicode/utf8"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    //"fmt"
)

type Role struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    OwnId int64 `form:"own_id" json:"own_id"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
}

type RoleUpdate struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    OwnId int64 `form:"own_id" json:"own_id"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
}

func (role Role) TableName() string {
    return "role_info"
}
func (roleUpdate RoleUpdate) TableName() string {
    return "role_info"
}

/**
 * 增加一条记录
 */
func RoleAddOne(c *gin.Context){
    var role Role
    err := c.ShouldBindJSON(&role);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 添加创建人
    customerId := GetCurrentCustomerId(c)
    
    role.CreatedCustomerId = customerId
    // 添加own_id
    ownId, err := GetCustomerOwnId(c, role.OwnId)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    role.OwnId = ownId
    // 插入
    affected, err := engine.Insert(&role)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "role":role,
    })
    c.JSON(http.StatusOK, result)
}

/**
 * 通过id为条件，更新一条记录
 */
func RoleUpdateById(c *gin.Context){
    var role RoleUpdate
    err := c.ShouldBindJSON(&role);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 添加own_id
    ownId, err := GetCustomerOwnId(c, role.OwnId)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    role.OwnId = ownId
    // 根据用户级别，得到更新的条件。
    roleUpdate := &RoleUpdate{Id:role.Id}
    customerType := GetCurrentCustomerType(c)
    if customerType != AdminSuperType {
        roleUpdate.OwnId = ownId
    }
    // 进行更新。
    affected, err := engine.Update(&role, roleUpdate)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    if affected == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("该数据不存在，或您没有权限编辑该数据"))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "role":role,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 删除一条记录
 */
func RoleDeleteById(c *gin.Context){
    var role Role
    var id DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&role)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "id":id.Id,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过ids批量删除数据
 */
func RoleDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var role Role
    var ids DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&role)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected": affected,
        "ids": ids.Ids,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 列表查询
 */
func RoleList(c *gin.Context){
    // params := c.Request.URL.Query()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    name     := c.DefaultQuery("name", "")
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    if name != "" {
        whereParam["name"] = []string{"like", name}
    }
    whereParam["created_at"] = []string{"scope", created_at_begin, created_at_end}
    // 根据用户的级别，通过own_id字段进行数据的过滤
    
    whereParam, err := OwnIdQueryFilter(c, whereParam)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    //whereParam["own_id"] = 93
    log.Println(whereParam)
    
    whereStr, whereVal := mysqldb.GetXOrmWhere(whereParam)
    log.Println(whereStr)
    log.Println(whereVal)
    // 进行查询
    query := engine.Limit(limit, (page-1)*limit)
    if whereStr != "" {
        query = query.Where(whereStr, whereVal...)
    }
    // 排序
    if sortD == "+" && sortColumns != "" {
        query = query.Asc(sortColumns)
    } else if sortD == "-" && sortColumns != "" {
        query = query.Desc(sortColumns)
    }
    // 得到查询count数
    var role Role
    counts, err := engine.Where(whereStr, whereVal...).Count(&role)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var roles []Role
    err = query.Find(&roles) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    ownIdOps, err := GetCustomerOwnIdOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    createdCustomerOps, err := GetRoleCreatedCustomerOps(roles)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    customerUsername := GetCurrentCustomerUsername(c)
    customerType := GetCurrentCustomerType(c)
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": roles,
        "total": counts,
        "createdCustomerOps": createdCustomerOps,
        "ownIdOps": ownIdOps,
        "customerUsername": customerUsername,
        "customerType": customerType,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

// vue将created_customer_id 渲染成 created customer username 所需要的slice
func GetRoleCreatedCustomerOps(roles []Role) ([]VueSelectOps, error){
    var groupArr []VueSelectOps
    var ids []int64
    for i:=0; i<len(roles); i++ {
        role := roles[i]
        ids = append(ids, role.CreatedCustomerId)
    }
    customers, err := GetCustomerUsernameByIds(ids)
    if err != nil{
        return nil, err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        groupArr = append(groupArr, VueSelectOps{Key: customer.Id, DisplayName: customer.Username})
    }
    // if groupArr == nil {
    //     return nil, errors.New("customer ids is empty")
    // }
    return groupArr, nil
}
