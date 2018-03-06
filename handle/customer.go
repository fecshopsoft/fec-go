package handle

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "crypto/md5"
    "errors"
    "time"
    "encoding/hex"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    //"fmt"
)

type Customer struct {
    Id int64
    Username string `form:"username" json:"username" binding:"required"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
    CreatedAt int64
    UpdatedAt int64
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
 * 列表查询
 */
func CustomerList(c *gin.Context){
    var customers []Customer
     _ = engine.Find(&customers)
    result := util.BuildSuccessResult(gin.H{
        "customers": customers,
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
    accesToken, err := security.JwtSignToken(customer)
    // 通过token，得到保存的值。
    /*
    data, expired, err := security.JwtParse(accesToken);
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    "data":data,
    "expired":expired,
    */
    result := util.BuildSuccessResult(gin.H{
        "accesToken": accesToken,
    })
    c.JSON(http.StatusOK, result)
}


func CustomerAccountIndex(c *gin.Context){
    
    result := util.BuildSuccessResult(gin.H{
        "status": "success",
        "currentCustomer": currentCustomer,
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




