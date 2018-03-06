package router

import(
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/handle"
    //"fmt"
)

func Listen(listenIp string) { 
    r := gin.Default()
    r.NoRoute(handle.NotFound)
    v1 := r.Group("/v1")
    {
        v1.POST("/customer/account/register", handle.CustomerAccountRegister)
        
        v1.POST("/customer/account/login", handle.CustomerAccountLogin)
        
        v1.POST("/customer/account/index", handle.PermissionAdmin, handle.CustomerAccountIndex)
        
        v1.GET("/customers/id/:id", handle.CustomerOneById)
        
        v1.GET("/customers/username/:username", handle.CustomerOneByUsername)
        
        v1.GET("/customers", handle.CustomerList)
        
        v1.POST("/customers", handle.CustomerAdd)
        
        
        v1.PATCH("/customers", handle.CustomerUpdateById)
        
        
        v1.DELETE("/customers/id/:id", handle.CustomerDeleteById)
        
        v1.GET("/customers/count", handle.CustomerCount)
        
        
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