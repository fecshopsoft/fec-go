package customer

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "errors"
    // "time"
    "unicode/utf8"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    "github.com/fecshopsoft/fec-go/helper"
    //"fmt"
)

type CustomerChildAdd struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    Email string `form:"email" json:"email"`
    Sex int `form:"sex" json:"sex"`
    Name string `form:"name" json:"name"`
    Telephone string `form:"telephone" json:"telephone"`
    Type int `form:"type" json:"type" `
    ParentId int64 `form:"parent_id" json:"parent_id" `
    Remark string `form:"remark" json:"remark"`
    MarketGroupId int `form:"market_group_id" json:"market_group_id"`
    JobType int `form:"job_type" json:"job_type"`
    Status int `form:"status" json:"status"`
    Age int `form:"age" json:"age"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    BirthDate  int64 `form:"birth_date" json:"birth_date"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
}

type CustomerChildUpdate struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    Email string `form:"email" json:"email"`
    Sex int `form:"sex" json:"sex"`
    Name string `form:"name" json:"name" binding:"required"`
    Telephone string `form:"telephone" json:"telephone"`
    Type int `form:"type" json:"type" binding:"required"`
    ParentId int64 `form:"parent_id" json:"parent_id"`
    Remark string `form:"remark" json:"remark"`
    MarketGroupId int `form:"market_group_id" json:"market_group_id"`
    JobType int `form:"job_type" json:"job_type"`
    Status int `form:"status" json:"status" binding:"required"`
    Age int `form:"age" json:"age"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    BirthDate  int64 `form:"birth_date" json:"birth_date"`
    Password string `xorm:"varchar(200)" form:"password" json:"password"`
}

func (customerChildAdd CustomerChildAdd) TableName() string {
    return "customer"
}

func (customerChildUpdate CustomerChildUpdate) TableName() string {
    return "customer"
}
/**
 * 增加一条记录
 */
func CustomerChildAddOne(c *gin.Context){
    var customer CustomerChildAdd
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 处理密码
    passwordEncry, err := getCustomerPassword(customer.Password)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    customer.Password = passwordEncry
    
    customer_id := helper.GetCurrentCustomerId(c)
    if customer_id == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("current customer id is empty"))
        return
    }
    customer_type := helper.GetCurrentCustomerType(c)
    if customer_type != helper.AdminCommonType {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("only common admin account can add child account"))
        return
    }
    customer.ParentId = customer_id
    customer.Type = helper.AdminChildType
    // 插入
    affected, err := engine.Insert(&customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "customer":customer,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过id为条件，更新一条记录
 */
func CustomerChildUpdateById(c *gin.Context){
    var customer CustomerChildUpdate
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    passwordEncry, err := getCustomerPassword(customer.Password)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    customer.Password = passwordEncry
    
    customer_id := helper.GetCurrentCustomerId(c)
    if customer_id == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("current customer id is empty"))
        return
    }
    customer_type := helper.GetCurrentCustomerType(c)
    if customer_type != helper.AdminCommonType {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("only common admin account can add child account"))
        return
    }
    // 在条件和更改中都加入parent_id 和 type，防止数据被恶意篡改。
    customer.ParentId = customer_id
    customer.Type = helper.AdminChildType
    
    affected, err := engine.Update(&customer, &Customer{Id:customer.Id, ParentId:customer_id, Type: helper.AdminChildType})
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "customer":customer,
    })
    c.JSON(http.StatusOK, result)
}

/**
 * 对密码进行处理，得到加密后的密码字符串
 */
func getCustomerChildPassword(password string) (string, error){
    if password == "" {
        return "", nil
    }
    if utf8.RuneCountInString(password) < PasswordMinLen {
        return "", errors.New("password length must >= " + strconv.Itoa(PasswordMinLen) )
    }
    encryptionPassStr, err := encryptionPass(password)
    if err != nil {
        return "", err
    }
    return encryptionPassStr, nil
}
/**
 * 删除一条记录
 */
func CustomerChildDeleteById(c *gin.Context){
    var customer Customer
    var id DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    customer_id := helper.GetCurrentCustomerId(c)
    if customer_id == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("current customer id is empty"))
        return
    }
    customer_type := helper.GetCurrentCustomerType(c)
    if customer_type != helper.AdminCommonType {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("only common admin account can add child account"))
        return
    }
    // 在条件和更改中都加入parent_id 和 type，防止数据被恶意篡改。
    // customer.ParentId = customer_id
    // customer.Type = helper.AdminChildType
    affected, err := engine.Where("id = ? and parent_id = ? and type = ?",id.Id, customer_id, helper.AdminChildType).Delete(&customer)
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
func CustomerChildDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var customer Customer
    var ids DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    customer_id := helper.GetCurrentCustomerId(c)
    if customer_id == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("current customer id is empty"))
        return
    }
    customer_type := helper.GetCurrentCustomerType(c)
    if customer_type != helper.AdminCommonType {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("only common admin account can add child account"))
        return
    }
    // 在条件和更改中都加入parent_id 和 type，防止数据被恶意篡改。
    // customer.ParentId = customer_id
    // customer.Type = helper.AdminChildType
    affected, err := engine.In("id", ids.Ids).Where("parent_id = ? and type = ?", customer_id, helper.AdminChildType).Delete(&customer)
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
func CustomerChildList(c *gin.Context){
    // log.Println(&CustomerToken{Username:"xxxx"})
    // log.Println(CustomerToken{Username:"xxxx"})
    // params := c.Request.URL.Query()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    status, _    := strconv.Atoi(c.DefaultQuery("status", ""))
    sex, _   := strconv.Atoi(c.DefaultQuery("sex", ""))
    username := c.DefaultQuery("username", "")
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    whereParam["status"] = status
    whereParam["sex"] = sex
    if username != "" {
        whereParam["username"] = []string{"like", username}
    }
    whereParam["type"] = helper.AdminChildType
    whereParam["parent_id"] = helper.GetCurrentCustomerId(c)
    whereParam["created_at"] = []string{"scope", created_at_begin, created_at_end}
    whereStr, whereVal := mysqldb.GetXOrmWhere(whereParam)
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
    var customer Customer
    counts, err := engine.Where(whereStr, whereVal...).Count(&customer)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var customers []Customer
    err = query.Find(&customers) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    marketGroupOps, err := GetCurrentMarketGroupOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": customers,
        "total":counts,
        "marketGroupOps": marketGroupOps,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}



type CMarketGroup struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    OwnId int64 `form:"own_id" json:"own_id"`
}
func (cMarketGroup CMarketGroup) TableName() string {
    return "base_market_group"
}

// 2,3用户对应的marketGroups
func GetCurrentMarketGroupOps(c *gin.Context) ([]VueSelectOps, error){
    var own_id int64
    var MGArr []VueSelectOps
    currentCustomerId := helper.GetCurrentCustomerId(c)
    GetCurrentCustomerType := helper.GetCurrentCustomerType(c)
    if GetCurrentCustomerType == helper.AdminCommonType {
        own_id = currentCustomerId
    } else if GetCurrentCustomerType == helper.AdminChildType {
        parent_id, err := GetCurrentCustomerParentId(c)
        if err != nil{
            return nil, err
        }
        own_id = parent_id
    }
    if GetCurrentCustomerType != helper.AdminSuperType && own_id == 0 {
        return nil, errors.New(" current own id is empty ")
    }
    
    var cMarketGroups []CMarketGroup
    var err error
    if own_id != 0 {
        err = engine.Where("own_id = ?", own_id).Find(&cMarketGroups) 
    } else {
        err = engine.Find(&cMarketGroups) 
    }
    if err != nil{
        return nil, err  
    }
    for i:=0; i<len(cMarketGroups); i++ {
        cMarketGroup := cMarketGroups[i]
        MGArr = append(MGArr, VueSelectOps{Key: cMarketGroup.Id, DisplayName: cMarketGroup.Name})
    }
    return MGArr, nil
}


/**
 * 通过ids，查询customer表，得到
 */
/*
func GetCustomerChildUsernameByIds(ids []int64) ([]CustomerUsername, error){
    // 得到结果数据
    var customers []CustomerUsername
    engine := mysqldb.GetEngine()
    err := engine.In("id", ids).Find(&customers) 
    if err != nil{
        return nil, err 
    }
    return customers, nil
}
*/

/**
 * 通过id查询一条记录
 */
func GetCustomerChildOneById(id int64) (Customer, error){
    var customer Customer
    has, err := engine.Where("id = ?", id).Get(&customer)
    if err != nil {
        return customer, err
    } 
    if has == false {
        return customer, errors.New("get customer by id error, empty data.")
    }
    return customer, nil
}

/**
 * 得到enable的common  admin 
 */
func GetAllEnableCommonCustomerChild() ([]CustomerUsername, error){
    // 得到结果数据
    var customers []CustomerUsername
    engine := mysqldb.GetEngine()
    err := engine.Where("type = ? and status = ? ", helper.AdminCommonType, statusEnable).Find(&customers) 
    if err != nil{
        return nil, err 
    }
    return customers, nil
}
