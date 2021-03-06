package customer

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "crypto/md5"
    "errors"
    // "time"
    "unicode/utf8"
    "encoding/hex"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    "github.com/fecshopsoft/fec-go/helper"
    //"fmt"
)

type CustomerToken struct {
    Id int64 `form:"id" json:"id"`
    Type int `form:"type" json:"type"`
    Username string `form:"username" json:"username" binding:"required"`
}

type CustomerUsername struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    JobType int `form:"job_type" json:"job_type"`
}

type Customer struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    Status int `form:"status" json:"status"`
    MarketGroupId int64 `form:"market_group_id" json:"market_group_id"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    Type int `form:"type" json:"type"`
	JobType int `form:"job_type" json:"job_type"`
}

type CustomerPassword struct {
    Password string `form:"password" json:"password" binding:"required"`
    NewPassword string `form:"new_password" json:"new_password" binding:"required"`
    ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}


type CustomerLogin struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
}

type CustomerAdd struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    JobType int `form:"job_type" json:"job_type"`
	MarketGroupId int64 `form:"market_group_id" json:"market_group_id"`
    Status int `form:"status" json:"status"`
    Type int `form:"type" json:"type" binding:"required"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
}

type CustomerUpdate struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    JobType int `form:"job_type" json:"job_type"`
	MarketGroupId int64 `form:"market_group_id" json:"market_group_id"`
    Status int `form:"status" json:"status" binding:"required"`
    Type int `form:"type" json:"type" binding:"required"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
   
    Password string `xorm:"varchar(200)" form:"password" json:"password"`
}

type CMarketGroup struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    OwnId int64 `form:"own_id" json:"own_id"`
}
func (cMarketGroup CMarketGroup) TableName() string {
    return "base_market_group"
}

// 账户激活状态的值
var statusEnable int = 1
// 密码最小长度
var PasswordMinLen int = 6
 

func (customerUpdate CustomerUpdate) TableName() string {
    return "customer"
}

func (customerAdd CustomerAdd) TableName() string {
    return "customer"
}

func (customerLogin CustomerLogin) TableName() string {
    return "customer"
}

func (customerToken CustomerToken) TableName() string {
    return "customer"
}

func (customerUsername CustomerUsername) TableName() string {
    return "customer"
}


/**
 * 增加一条记录
 */
func CustomerAddOne(c *gin.Context){
    var customer CustomerAdd
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
func CustomerUpdateById(c *gin.Context){
    var customer CustomerUpdate
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
    
    affected, err := engine.Where("id = ?", customer.Id).MustCols("parent_id").Update(&customer)
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
/*
func CustomerUpdatePayInfoById(c *gin.Context){
    var customer CustomerUpdatePayInfo
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    affected, err := engine.Where("id = ?", customer.Id).Update(&customer)
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
*/

/**
 * 更新密码的格式等信息是否满足要求
 */
func checkPassFormat(customer CustomerPassword) (bool, error){
    if utf8.RuneCountInString(customer.NewPassword) < PasswordMinLen {
        return false, errors.New("new password length must >= " + strconv.Itoa(PasswordMinLen))
    }
    if customer.ConfirmPassword != customer.NewPassword {
        return false, errors.New("New Password and confirmation password must be the same")
    }
    return true, nil
}
/**
 * 更新密码
 */
func CustomerUpdatePassword(c *gin.Context){
    var customer CustomerPassword
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 格式检查
    _, err = checkPassFormat(customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 更新密码
    encryptionPassStr, err := encryptionPass(customer.Password)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    newEncryptionPassStr, err := encryptionPass(customer.NewPassword)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    username := helper.GetCurrentCustomerUsername(c)
    var customerLogin CustomerLogin
    customerLogin.Username = username
    customerLogin.Password = newEncryptionPassStr
    affected, err := engine.Update(&customerLogin, &CustomerLogin{Username:username,Password:encryptionPassStr})
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "username":username,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 对密码进行处理，得到加密后的密码字符串
 */
func getCustomerPassword(password string) (string, error){
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
func CustomerDeleteById(c *gin.Context){
    var customer Customer
    var id DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&customer)
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
func CustomerDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var customer Customer
    var ids DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&customer)
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
 * 用户登录
 */
func CustomerAccountLogin(c *gin.Context){
    var customerToken CustomerToken
    var customerLogin CustomerLogin
    err := c.ShouldBindJSON(&customerLogin)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    encryptionPassStr, err := encryptionPass(customerLogin.Password)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    //has, err := engine.Where(" password = ? ", encryptionPassStr).And(" username = ?", customer.Username).Get(&customer)
    has, err := engine.Where("username = ? and password = ? and status = ? ", customerLogin.Username, encryptionPassStr, statusEnable).Get(&customerToken)
    
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    if has != true{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("该账户不存在或密码错误"))
        return  
    }
    // 如果用户类型是 helper.AdminChildType,则判断他的主账户是否是有效的
	/*
    if customerToken.Type == helper.AdminChildType {
        parentId := customerToken.ParentId
        var customerParentToken CustomerToken
        //has, err := engine.Where(" password = ? ", encryptionPassStr).And(" username = ?", customer.Username).Get(&customer)
        has, err := engine.Where("id = ?  and status = ? ", parentId, statusEnable).Get(&customerParentToken)
        
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return  
        }
        if has != true{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("该账户对应的主账户已经被禁用，因此您无法登录"))
            return  
        }
    }
	*/
    
    token, err := security.JwtSignToken(customerToken)
    result := util.BuildSuccessResult(gin.H{
        "token": token,
    })
    c.JSON(http.StatusOK, result)    
}
/**
 * 账户中心，获取当前用户的数据
 */
func CustomerAccountIndex(c *gin.Context){
    var roles []string
    customerType := helper.GetCurrentCustomerType(c)
    if customerType == helper.AdminType {
        roles = append(roles, helper.VueUserRoles[helper.AdminType])
    } else  {
        roles = append(roles, helper.VueUserRoles[helper.CommonType])
    } 
	
    // roles = append(roles, "editor")
    // cCustomer := currentCustomer.(map[string]interface{})
    cCustomer := helper.GetCurrentCustomer(c)
    result := util.BuildSuccessResult(gin.H{
        "name": cCustomer["username"],
        "avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
        "roles": roles,
    })
    c.JSON(http.StatusOK, result)
    
}

/**
 * 密码加密，得到字符串md5加密后的字符串
 */
func encryptionPass(password string) (string, error){
    if password == "" {
        return "", errors.New("password is empty")
    }
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(password))
    cipherStr := md5Ctx.Sum(nil)
    return hex.EncodeToString(cipherStr), nil
}
/**
 * 列表查询
 */
func CustomerList(c *gin.Context){
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
    //parent_id, _    := strconv.Atoi(c.DefaultQuery("parent_id", ""))
    status, _    := strconv.Atoi(c.DefaultQuery("status", ""))
    accountType, _    := strconv.Atoi(c.DefaultQuery("type", ""))
    username := c.DefaultQuery("username", "")
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    //whereParam["parent_id"] = parent_id
    whereParam["status"] = status
    whereParam["type"] = accountType
    if username != "" {
        whereParam["username"] = []string{"like", username}
    }
    //whereParam["age"] = []string{"scope","2","20"}
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
    // 
    //commonAdminAccount, err := GetCustomerOwnIdOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
	marketGroupOps, err := GetCurrentMarketGroupOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    //customerType := helper.GetCurrentCustomerType(c)
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": customers,
        "total":counts,
        "typeOps": GetCustomerTypeName(),
		"marketGroupOps": marketGroupOps,
    //    "commonAdminOps": commonAdminAccount,
        //"customerType": customerType,
		
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

func GetCurrentMarketGroupOps(c *gin.Context) ([]VueSelectOps, error){
    var MGArr []VueSelectOps
    var cMarketGroups []CMarketGroup
    var err error
    err = engine.Find(&cMarketGroups) 
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
func GetCustomerUsernameByIds(ids []int64) ([]CustomerUsername, error){
    // 得到结果数据
    var customers []CustomerUsername
    engine := mysqldb.GetEngine()
    err := engine.In("id", ids).Find(&customers) 
    if err != nil{
        return nil, err 
    }
    return customers, nil
}


/**
 * 通过ids，查询customer表，得到
 */
func GetActiveCustomers() ([]Customer, error){
    // newTime := time.Now().Unix()
    enableStatus := 1
    // 得到结果数据
    var customers []Customer
    engine := mysqldb.GetEngine()
    err := engine.Where("status = ? ", enableStatus).Find(&customers) 
    if err != nil{
        return nil, err 
    }
    return customers, nil
}


/**
 * 通过ids，查询customer表，得到
 */
/*
func GetPaymentActiveCustomers() ([]Customer, error){
    // newTime := time.Now().Unix()
    enableStatus := 1
    // 得到结果数据
    var customers []Customer
    engine := mysqldb.GetEngine()
    err := engine.Where("status = ? ", enableStatus).Find(&customers) 
    if err != nil{
        return nil, err 
    }
    return customers, nil
}
*/


/**
 * 通过id查询一条记录
 */
func GetCustomerOneById(id int64) (Customer, error){
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

// 找到own_id下的美工员工 job_type = 2

func GetDesigiPerson() ([]CustomerUsername, error) {
    var customers []CustomerUsername
    enableStatus := 1
    jobType := 2
    err := engine.
        Where("job_type = ? and status = ?", jobType, enableStatus).
        Find(&customers)
    if err != nil {
        return customers, err
    } 
    
    return customers, nil
}



/**
 * 得到enable的common  admin 
 */
/*
func GetAllEnableCommonCustomer() ([]CustomerUsername, error){
    // 得到结果数据
    var customers []CustomerUsername
    engine := mysqldb.GetEngine()
    err := engine.Where("type = ? and status = ? ", helper.AdminCommonType, statusEnable).Find(&customers) 
    if err != nil{
        return nil, err 
    }
    return customers, nil
}
*/

/**
 * 通过id查询customer
 */
/*
func CustomerOneById(c *gin.Context){
    engine := mysqldb.GetEngine()
    customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    var customer Customer
    if customerId == 0 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("customer id is empty"))
        return
    }
    has, err := engine.Where("id = ?", customerId).Get(&customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    if has != true {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("customer is not exist"))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "customer":customer,
    })
    c.JSON(http.StatusOK, result)
}
*/
/**
 * 通过username查询customer
 */
/*
func CustomerOneByUsername(c *gin.Context){
    username := c.Param("username")
    var customer Customer
    if username == "" {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("username is empty"))
        return
    }
    has, err := engine.Where("username = ?", username).Get(&customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    if has != true {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("customer is not exist"))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "customer": customer,
    })
    c.JSON(http.StatusOK, result)
}
*/
/**
 * 得到用户总数
 */
/*
func CustomerCount(c *gin.Context){
    var customer Customer
    
    counts, err := engine.Count(&customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "counts":counts,
    })
    c.JSON(http.StatusOK, result)
}
*/
/**
 * 用户注册
 */
/* 
func CustomerAccountRegister(c *gin.Context){
    var customer CustomerSave
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    encryptionPass, err := encryptionPass(customer.Password)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    customer.Password = encryptionPass
    customer.CreatedAt = time.Now().Unix();
    customer.UpdatedAt  = time.Now().Unix();
    affected, err := engine.Insert(&customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    if affected < 1 {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("注册用户插入数据失败"))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "status": "success",
    })
    c.JSON(http.StatusOK, result)
}
*/