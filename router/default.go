package router

import(
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/handle"
    //"fmt"
    //"time"
    "net/http"
)

func Listen(listenIp string) { 
    r := gin.Default()
    r.NoRoute(handle.NotFound)
    r.Use(handle.CORSMiddleware())
    
    //mi := router.Group("/mi", handler.ApiGlobal, handler.AdminCheckLogin)
    v1 := r.Group("/v1")
    {
        /*
        v1.POST("/customer/account/login", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"code":20000,"data":gin.H{"token":"admin"}})
        })
        v1.GET("/customer/account/index", func(c *gin.Context) {
            var roles,role []string
            roles = append(roles, "admin44")
            role  = append(role, "admin44s")
            c.JSON(http.StatusOK, gin.H{
                "code":20000,
                "data":gin.H{
                    "roles":roles,
                    "role":role,
                    "name":"admin",
                    "avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
                },
            })
        })
        */
        // 退出登录
        v1.POST("/customer/account/logout", handle.PermissionAdmin, func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"code":20000,"data":gin.H{"token":"admin"}})
        })
        // 登录
        v1.POST("/customer/account/login", handle.CustomerAccountLogin)
        // 得到账户信息
        v1.GET("/customer/account/index",  handle.PermissionAdmin, handle.CustomerAccountIndex)
        // 得到customer 列表
        v1.GET("/customer/list", handle.PermissionAdmin, handle.CustomerList)
        // 增加一个用户
        v1.POST("/customer/addone", handle.PermissionAdmin, handle.CustomerAddOne)
        // 更新一个用户
        v1.PATCH("/customer/updateone", handle.PermissionAdmin, handle.CustomerUpdateById)
        // 更新一个用户
        v1.PATCH("/customer/updatepassword", handle.PermissionAdmin, handle.CustomerUpdatePassword)
        
        // 删除一个用户
        v1.DELETE("/customers/id/:id", handle.PermissionAdmin, handle.CustomerDeleteById)
        // 批量删除用户
        v1.DELETE("/customer/batch", handle.PermissionAdmin, handle.CustomerDeleteByIds)
        
        /*
        v1.GET("/customer/list", func(c *gin.Context) {
            item := []gin.H{}
            item = append(item, gin.H{
                        "id":22,
                        "date":"xxxxxxxxx",
                        "title":"zzzzzzzzz",
                        "author":22,
                        "reviewer":"reviewer",
                        "importance":"3",
                        "readings":"readings",
                        "timestamp":time.Now().Unix(),
                        "forecast":"SSSSS",
                        "type":[]string{"CN"},
                        "status":"published",
                        "pageviews":"XXX",
                    })
            c.JSON(http.StatusOK, gin.H{
                "code":20000,
                "data":gin.H{
                    "items":item, 
                    "total":500,
                },
            })
        })
        */
        
        
        /*
        v1.POST("/customer/account/register", handle.CustomerAccountRegister) 
        
        
        v1.GET("/customers/id/:id", handle.CustomerOneById)
        
        v1.GET("/customers/username/:username", handle.CustomerOneByUsername)
        
        v1.GET("/customers", handle.CustomerList)
        
        
        
        
        v1.PATCH("/customers", handle.CustomerUpdateById)
        
        
        
        
        v1.GET("/customers/count", handle.CustomerCount)
        
        */
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