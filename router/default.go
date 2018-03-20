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
    // ##：【handle.PermissionLoginToken】 验证是否有token，以及token是否有效
    // ##：【handle.PermissionRole】 验证用户是否有权限访问该资源
    v1 := r.Group("/v1")
    {
        // #### customer login ####
        // 登录账户
        v1.POST("/customer/account/login", handle.CustomerAccountLogin)
        // 得到账户信息
        v1.GET("/customer/account/index",  handle.PermissionLoginToken, handle.CustomerAccountIndex)
        // 退出登录
        v1.POST("/customer/account/logout", handle.PermissionLoginToken, func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"code":20000,"data":gin.H{"token":"admin"}})
        })
        
        // #### customer account ####
        // 得到customer 列表
        v1.GET("/customer/list", handle.PermissionLoginToken, handle.PermissionRole, handle.CustomerList)
        // 增加一个用户
        v1.POST("/customer/addone", handle.PermissionLoginToken, handle.CustomerAddOne)
        // 更新一个用户
        v1.PATCH("/customer/updateone", handle.PermissionLoginToken, handle.CustomerUpdateById)
        // 更新一个用户
        v1.PATCH("/customer/updatepassword", handle.PermissionLoginToken, handle.CustomerUpdatePassword)
        // 删除一个用户
        v1.DELETE("/customer/deleteone", handle.PermissionLoginToken, handle.CustomerDeleteById)
        // 批量删除用户
        v1.DELETE("/customer/deletebatch", handle.PermissionLoginToken, handle.CustomerDeleteByIds)
        
        // #### customer resource group ####
        // 得到resource group 列表
        v1.GET("/customer/resourcegroup/list", handle.PermissionLoginToken, handle.ResourceGroupList)
        // 增加一个resource group
        v1.POST("/customer/resourcegroup/addone", handle.PermissionLoginToken, handle.ResourceGroupAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resourcegroup/updateone", handle.PermissionLoginToken, handle.ResourceGroupUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resourcegroup/deleteone", handle.PermissionLoginToken, handle.ResourceGroupDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resourcegroup/deletebatch", handle.PermissionLoginToken, handle.ResourceGroupDeleteByIds)
        
        
        
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
    }
    r.Run(listenIp) // 这里改成您的ip和端口
}