package router

import(
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/fecshopsoft/fec-go/handle"
    //"fmt"
)

func Listen(listenIp string, engine *xorm.Engine) { 
    r := gin.Default()
    v1 := r.Group("/v1")
    {
        v1.POST("/customer/account/register", func(c *gin.Context) {
            handle.CustomerAccountRegister(c, engine);
        })
        
        
        v1.POST("/customer/account/login", func(c *gin.Context) {
            handle.CustomerAccountLogin(c, engine);
        })
        
        
        
        
        
        
        
        
        v1.GET("/customers/id/:id", func(c *gin.Context) {
            handle.CustomerOneById(c, engine);
        })
        
        v1.GET("/customers/username/:username", func(c *gin.Context) {
            handle.CustomerOneByUsername(c, engine);
        })
        
        v1.GET("/customers", func(c *gin.Context) {
            handle.CustomerList(c, engine);
        })
        
        v1.POST("/customers", func(c *gin.Context) {
            handle.CustomerAdd(c, engine);
        })
        
        
        v1.PATCH("/customers", func(c *gin.Context) {
            handle.CustomerUpdateById(c, engine);
        })
        
        
        v1.DELETE("/customers/id/:id", func(c *gin.Context) {
            handle.CustomerDeleteById(c, engine);
        })
        
        v1.GET("/customers/count", func(c *gin.Context) {
            handle.CustomerCount(c, engine);
        })
        
        
        /*
        v1.POST("/users", func(c *gin.Context) {
            data := testMysql.AddOne(mysqlDB, c);
            c.JSON(http.StatusOK, data)
        })
        v1.PATCH("/users/:id", func(c *gin.Context) {
            data := testMysql.UpdateById(mysqlDB, c);
            c.JSON(http.StatusOK, data)
        })
        v1.DELETE("/users/:id", func(c *gin.Context) {
            data := testMysql.DeleteById(mysqlDB, c);
            c.JSON(http.StatusOK, data)
        })
        */
    }
    r.Run(listenIp) // 这里改成您的ip和端口
}