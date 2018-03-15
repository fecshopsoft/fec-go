package handle

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "crypto/md5"
    "errors"
    "time"
    "unicode/utf8"
    "encoding/hex"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    //"fmt"
)

type Customer struct {
    Id int64 `form:"id" json:"id"`
    Username string `form:"username" json:"username" binding:"required"`
    Email string `form:"email" json:"email"`
    Sex int `form:"sex" json:"sex"`
    Name string `form:"name" json:"name"`
    Telephone string `form:"telephone" json:"telephone"`
    Remark string `form:"remark" json:"remark"`
    Status int `form:"status" json:"status"`
    Age int `form:"age" json:"age"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
    CreatedAt int64 `form:"created_at" json:"created_at"`
    UpdatedAt int64 `form:"updated_at" json:"updated_at"`
    BirthDate  int64 `form:"birth_date" json:"birth_date"`
    
}

var engine *(xorm.Engine)

func init(){
    engine = mysqldb.GetEngine()
}
/**
 * 通过id查询customer
 */
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
/**
 * 通过username查询customer
 */
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



/**
 * 增加一条记录
 */
func CustomerAdd(c *gin.Context){
    var customer Customer
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
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
 * 更新一条记录
 */
func CustomerUpdateById(c *gin.Context){
    var customer Customer
    err := c.ShouldBindJSON(&customer);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Update(&customer, &Customer{Id:customer.Id})
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
 * 删除一条记录
 */
func CustomerDeleteById(c *gin.Context){
    var customer Customer
    customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    affected, err := engine.Where("id = ?",customerId).Delete(&customer)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "id":customerId,
    })
    c.JSON(http.StatusOK, result)
}


/**
 * 删除一条记录
 */
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
/**
 * 用户注册
 */
func CustomerAccountRegister(c *gin.Context){
    var customer Customer
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


type CustomerLogin struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
}


/**
 * 用户登录
 */
func CustomerAccountLogin(c *gin.Context){
    var customer Customer
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
    has, err := engine.Where("username = ? and password = ?", customerLogin.Username, encryptionPassStr).Get(&customer)
    
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    if has != true{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("该账户不存在或密码错误"))
        return  
    }
    customer.Password = ""
    token, err := security.JwtSignToken(customer)
    result := util.BuildSuccessResult(gin.H{
        "token": token,
    })
    c.JSON(http.StatusOK, result)    
    
}

/**
 * 账户中心
 */
func CustomerAccountIndex(c *gin.Context){
    var roles []string
    roles = append(roles, "admin")
    //roles = append(roles, "editor")
    cCustomer := currentCustomer.(map[string]interface{})
    result := util.BuildSuccessResult(gin.H{
        "name": cCustomer["username"],
        "avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
        "roles": roles,
    })
    c.JSON(http.StatusOK, result)
    
}

/**
 * 密码加密
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
    // params := c.Request.URL.Query()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    page, _  := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
    id, _    := strconv.Atoi(c.DefaultQuery("id", ""))
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
    whereParam["id"] = id
    whereParam["sex"] = sex
    whereParam["username"] = []string{"like", username}
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
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": customers,
        "total":counts,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}



