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
        
        v1.POST("/customer/account/logout", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"code":20000,"data":gin.H{"token":"admin"}})
        })
        
        v1.POST("/customer/account/login", handle.CustomerAccountLogin)
        
        v1.GET("/customer/account/index",  handle.PermissionAdmin, handle.CustomerAccountIndex)
        
        v1.GET("/customer/list", handle.CustomerList)
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
        
        v1.POST("/customers", handle.CustomerAdd)
        
        
        v1.PATCH("/customers", handle.CustomerUpdateById)
        
        
        v1.DELETE("/customers/id/:id", handle.CustomerDeleteById)
        
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