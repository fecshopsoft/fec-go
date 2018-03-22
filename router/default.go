package router

import(
    "github.com/gin-gonic/gin"
    customerHandler "github.com/fecshopsoft/fec-go/handler/customer"
    "github.com/fecshopsoft/fec-go/middleware"
    "github.com/fecshopsoft/fec-go/config"
    // "github.com/fecshopsoft/fec-go/initialization"
    //"fmt"
    //"time"
    "os"
    "io"
    // "log"
    // "time"
    "path/filepath"
    "net/http"
)

func Listen(listenIp string) { 
    
    // gin.DisableConsoleColor()
    // log.Println("------333：" + time.Now().String())
    initLog()
    // log.Println("------444：" + time.Now().String())
    r := gin.Default()
    r.NoRoute(middleware.NotFound)
    // 初始化cors
    r.Use(middleware.CORS())
    // 初始化上下文中的全局变量
    r.Use(middleware.InitContext)
    //mi := router.Group("/mi", handler.ApiGlobal, handler.AdminCheckLogin)
    // ##：【middleware.PermissionLoginToken】 验证是否有token，以及token是否有效
    // ##：【customerHandler.PermissionRole】 验证用户是否有权限访问该资源
    v1 := r.Group("/v1")
    {
        // #### customer login ####
        // 登录账户
        v1.POST("/customer/account/login",  customerHandler.CustomerAccountLogin)
        // 得到账户信息
        v1.GET("/customer/account/index",   middleware.PermissionLoginToken, customerHandler.CustomerAccountIndex)
        // 退出登录
        v1.POST("/customer/account/logout", middleware.PermissionLoginToken, func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"code":20000,"data":gin.H{"token":"admin"}})
        })
        
        // #### customer account ####
        // 得到 customer 列表
        v1.GET("/customer/list",                middleware.PermissionLoginToken, middleware.PermissionRole, customerHandler.CustomerList)
        // 增加一个用户
        v1.POST("/customer/addone",             middleware.PermissionLoginToken, customerHandler.CustomerAddOne)
        // 更新一个用户
        v1.PATCH("/customer/updateone",         middleware.PermissionLoginToken, customerHandler.CustomerUpdateById)
        // 更新一个用户
        v1.PATCH("/customer/updatepassword",    middleware.PermissionLoginToken, customerHandler.CustomerUpdatePassword)
        // 删除一个用户
        v1.DELETE("/customer/deleteone",        middleware.PermissionLoginToken, customerHandler.CustomerDeleteById)
        // 批量删除用户
        v1.DELETE("/customer/deletebatch",      middleware.PermissionLoginToken, customerHandler.CustomerDeleteByIds)
        
        // #### customer resource group ####
        // 得到 resource group 列表
        v1.GET("/customer/resourcegroup/list",              middleware.PermissionLoginToken, customerHandler.ResourceGroupList)
        // 增加一个resource group
        v1.POST("/customer/resourcegroup/addone",           middleware.PermissionLoginToken, customerHandler.ResourceGroupAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resourcegroup/updateone",       middleware.PermissionLoginToken, customerHandler.ResourceGroupUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resourcegroup/deleteone",      middleware.PermissionLoginToken, customerHandler.ResourceGroupDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resourcegroup/deletebatch",    middleware.PermissionLoginToken, customerHandler.ResourceGroupDeleteByIds)
        
        // #### customer resource ####
        // 得到 resource 列表
        v1.GET("/customer/resource/list",           middleware.PermissionLoginToken, customerHandler.ResourceList)
        // 增加一个resource group
        v1.POST("/customer/resource/addone",        middleware.PermissionLoginToken, customerHandler.ResourceAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resource/updateone",    middleware.PermissionLoginToken, customerHandler.ResourceUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resource/deleteone",   middleware.PermissionLoginToken, customerHandler.ResourceDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resource/deletebatch", middleware.PermissionLoginToken, customerHandler.ResourceDeleteByIds)
        
        // #### customer role ####
        // 得到 role 列表
        v1.GET("/customer/role/list",               middleware.PermissionLoginToken, customerHandler.RoleList)
        // 增加一个resource group
        v1.POST("/customer/role/addone",            middleware.PermissionLoginToken, customerHandler.RoleAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/role/updateone",        middleware.PermissionLoginToken, customerHandler.RoleUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/role/deleteone",       middleware.PermissionLoginToken, customerHandler.RoleDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/role/deletebatch",     middleware.PermissionLoginToken, customerHandler.RoleDeleteByIds)
        
        
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
        v1.POST("/customer/account/register", customerHandler.CustomerAccountRegister) 
        
        
        v1.GET("/customers/id/:id", customerHandler.CustomerOneById)
        
        v1.GET("/customers/username/:username", customerHandler.CustomerOneByUsername)
        
        v1.GET("/customers", customerHandler.CustomerList)
        
        
        
        
        v1.PATCH("/customers", customerHandler.CustomerUpdateById)
        
        
        
        
        v1.GET("/customers/count", customerHandler.CustomerCount)
        
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
// 设置 gin.DefaultWriter 和 gin.DefaultErrorWriter
func initLog() {
	if "false" == config.Get("output_log") {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	// Logging to a file.
	//gin_file, _ := os.Create("gin.log")
	routerInfoLogUrl := config.Get("router_info_log")
	if routerInfoLogUrl == "" {
		routerInfoLogUrl = "logs/router_info.log"
	}
	path := filepath.Dir(routerInfoLogUrl)
	os.MkdirAll(path, 0777)

	routerErrorLogUrl := config.Get("router_error_log")
	if routerErrorLogUrl == "" {
		routerErrorLogUrl = "logs/router_error.log"
	}
	path = filepath.Dir(routerErrorLogUrl)
	os.MkdirAll(path, 0777)
	infoFile, err := os.OpenFile(routerInfoLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err2 := os.OpenFile(routerErrorLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err2 != nil {
		panic(err2)
	}
	//ginLogUrl := "gin.log"
	//var gin_file *os.File
	//if _,err := os.Stat(ginLogUrl);err!=nil{
	//	gin_file, _ = os.Create(ginLogUrl)
	//}else{
	//	gin_file, _ = os.OpenFile("gin.log", os.O_RDWR|os.O_APPEND, 0666)
	//}
	//gin_error_file, _ := os.Create(routerErrorLogUrl)
	gin.DefaultWriter = io.MultiWriter(infoFile)
	gin.DefaultErrorWriter = errorFile
	//gin.RecoveryWithWriter(gin_error_file)
}