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
    //"fmt"
)

type Customer struct {
    Id int64
    Username string `form:"username" json:"username" binding:"required"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
    CreatedAt int64
    UpdatedAt int64
}

/**
 * 通过id查询customer
 */
func CustomerOneById(c *gin.Context, engine *xorm.Engine){
    customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var customer Customer
    if customerId != 0 {
        has, err := engine.Where("id = ?", customerId).Get(&customer)
        if err == nil && has == true {
            c.JSON(http.StatusOK, gin.H{"result":customer})
        }
    }
}
/**
 * 通过username查询customer
 */
func CustomerOneByUsername(c *gin.Context, engine *xorm.Engine){
    username := c.Param("username")
    var customer Customer
    if username != "" {
        has, err := engine.Where("username = ?", username).Get(&customer)
        if err == nil && has == true {
            c.JSON(http.StatusOK, gin.H{"result":customer})
        }
    }
}

/**
 * 列表查询
 */
func CustomerList(c *gin.Context, engine *xorm.Engine){
    var customers []Customer
     _ = engine.Find(&customers)
    c.JSON(http.StatusOK, gin.H{"result":customers})
}

/**
 * 增加一条记录
 */
func CustomerAdd(c *gin.Context, engine *xorm.Engine){
    var customer Customer
    if err := c.ShouldBindJSON(&customer); err == nil {
        affected, err := engine.Insert(&customer)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusOK, gin.H{
                "affected":affected,
                "customer":customer,
            })
        }
    }else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}


/**
 * 更新一条记录
 */
func CustomerUpdateById(c *gin.Context, engine *xorm.Engine){
    var customer Customer
    if err := c.ShouldBindJSON(&customer); err == nil {
        affected, err := engine.Update(&customer, &Customer{Id:customer.Id})
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusOK, gin.H{
                "affected":affected,
                "customer":customer,
            })
        }
    }else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}



/**
 * 删除一条记录
 */
func CustomerDeleteById(c *gin.Context, engine *xorm.Engine){
    var customer Customer
    customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    affected, err := engine.Where("id = ?",customerId).Delete(&customer)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    } 
    c.JSON(http.StatusOK, gin.H{
        "affected":affected,
        "id":customerId,
    })
}


/**
 * 删除一条记录
 */
func CustomerCount(c *gin.Context, engine *xorm.Engine){
    var customer Customer
    
    counts, err := engine.Count(&customer)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    } 
    c.JSON(http.StatusOK, gin.H{
        "counts":counts,
    })
}



/**
 * 用户注册
 */
func CustomerAccountRegister(c *gin.Context, engine *xorm.Engine){
    var customer Customer
    if err := c.ShouldBindJSON(&customer); err == nil {
        encryptionPass, err := encryptionPass(customer.Password)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        customer.Password = encryptionPass
        customer.CreatedAt = time.Now().Unix();
        customer.UpdatedAt  = time.Now().Unix();
        affected, err := engine.Insert(&customer)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        } else if affected > 0 {
            c.JSON(http.StatusOK, gin.H{"status": "success"})
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}



type CustomerLogin struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `xorm:"varchar(200)" form:"password" json:"password" binding:"required"`
}


/**
 * 用户登录
 */
func CustomerAccountLogin(c *gin.Context, engine *xorm.Engine){
    var customer Customer
    var customerLogin CustomerLogin
    if err := c.ShouldBindJSON(&customerLogin); err == nil {
        encryptionPassStr, err := encryptionPass(customerLogin.Password)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        //has, err := engine.Where(" password = ? ", encryptionPassStr).And(" username = ?", customer.Username).Get(&customer)
        has, err := engine.Where("username = ? and password = ?", customerLogin.Username, encryptionPassStr).Get(&customer)
        if err == nil && has == true {
            accesToken, err := security.JwtSignToken(customer)
            // 通过token，得到保存的值。
            data, expired, err := security.JwtParse(accesToken);
            if err != nil{
                c.JSON(http.StatusOK, gin.H{
                    "loginStatus":"fail",
                    "error": err.Error(),
                    
                })
            } else {
                c.JSON(http.StatusOK, gin.H{
                    "loginStatus": "success",
                    "accesToken": accesToken,
                    "data":data,
                    "expired":expired,
                })
            }
        } else if err != nil{
            c.JSON(http.StatusOK, gin.H{
                "loginStatus":"fail",
                "error": err.Error(),
            })
        } else {
            c.JSON(http.StatusOK, gin.H{
                "loginStatus":"fail",
                "error": "customer account is not exist",
            })
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
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




