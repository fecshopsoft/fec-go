package router

import(
    "github.com/gin-gonic/gin"
    customerHandler "github.com/fecshopsoft/fec-go/handler/customer"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
    fecHandler "github.com/fecshopsoft/fec-go/handler/fec"
    testHandler "github.com/fecshopsoft/fec-go/handler/test"
    cronHandler "github.com/fecshopsoft/fec-go/handler/cron"
    wholeHandler "github.com/fecshopsoft/fec-go/handler/whole"
    "github.com/fecshopsoft/fec-go/middleware"
    "github.com/fecshopsoft/fec-go/config"
    "github.com/fecshopsoft/fec-go/initialization"
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
    err := initialization.InitWebsiteInfo()
    if err != nil {
        panic(err)
    }
    // log.Println("------444：" + time.Now().String())
    r := gin.Default()
    r.NoRoute(middleware.NotFound)
    // 初始化cors
    r.Use(middleware.CORS())
    // 初始化上下文中的全局变量
    r.Use(middleware.InitContext)
    r.GET("/fec/trace", fecHandler.PermisstionWebsiteId, fecHandler.SaveJsData)
    // r.GET("/fec/ip", fecHandler.Iptest)
    r.POST("/fec/trace/api", fecHandler.PermissionAccessToken, fecHandler.SaveApiData)
    r.GET("/test/mgo", testHandler.MgoFind)
    r.GET("/test/es", testHandler.EsFind)
    r.GET("/test/mgo/mapreduce", testHandler.MgoMapReduce)
    // 这个函数是为了更新上面的 initialization.InitWebsiteInfo()，需要在cron中1分钟调用一次刷新。
    r.GET("/fec/trace/cronssss", cronHandler.UpdateSite)
    
    //mi := router.Group("/mi", handler.ApiGlobal, handler.AdminCheckLogin)
    // ##：【middleware.PermissionLoginToken】 验证是否有token，以及token是否有效
    // ##：【customerHandler.PermissionRole】 验证用户是否有权限访问该资源
    v1 := r.Group("/v1")
    {
        
        
        
        // #### customer login ####
        // 登录账户 - 权限：1,2,3
        v1.POST("/customer/account/login",  customerHandler.CustomerAccountLogin)
        // 得到账户信息 - 权限：1,2,3
        v1.GET("/customer/account/index",   middleware.PermissionLoginToken, customerHandler.CustomerAccountIndex)
        // 退出登录 - 权限：1,2,3
        v1.POST("/customer/account/logout", middleware.PermissionLoginToken, func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"code":20000,"data":gin.H{"token":"admin"}})
        })
        
        // 更新账户密码 权限：1，2,3
        v1.PATCH("/customer/updatepassword",    middleware.PermissionLoginToken, customerHandler.CustomerUpdatePassword)
        
        // #### customer account #### 权限：1 ，只有super admin才能访问
        // 得到 customer 列表  
        v1.GET("/customer/list",                middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.CustomerList)
        // 增加一个用户
        v1.POST("/customer/addone",             middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.CustomerAddOne)
        // 更新一个用户
        v1.PATCH("/customer/updateone",         middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.CustomerUpdateById)
        // 更新一个用户的付费信息
        v1.PATCH("/customer/updateonepayinfo",         middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.CustomerUpdatePayInfoById)
        
        // 删除一个用户
        v1.DELETE("/customer/deleteone",        middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.CustomerDeleteById)
        // 批量删除用户
        v1.DELETE("/customer/deletebatch",      middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.CustomerDeleteByIds)
        
        // 得到 customer 列表  role: 1,2,3
        v1.GET("/customer/child/list",                middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerChildList)
        // 增加一个用户 role: 1,2
        v1.POST("/customer/child/addone",             middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerChildAddOne)
        // 更新一个用户 role: 1,2
        v1.PATCH("/customer/child/updateone",         middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerChildUpdateById)
        // 删除一个用户 role: 1,2
        v1.DELETE("/customer/child/deleteone",        middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerChildDeleteById)
        // 批量删除用户 role: 1,2
        v1.DELETE("/customer/child/deletebatch",      middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerChildDeleteByIds)
        
        v1.GET("/customer/child/role/allandselected", middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerRoleAllAndSelect)
        // 更新customer 的 roles
        v1.PATCH("/customer/child/role/updateone",    middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.CustomerRoleUpdate)
        
        
        // #### customer resource group #### 权限：1 ，只有super admin才能访问
        // 得到 resource group 列表
        v1.GET("/customer/resourcegroup/list",              middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceGroupList)
        // 增加一个resource group
        v1.POST("/customer/resourcegroup/addone",           middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceGroupAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resourcegroup/updateone",       middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceGroupUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resourcegroup/deleteone",      middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceGroupDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resourcegroup/deletebatch",    middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceGroupDeleteByIds)
        
        // #### customer resource #### 权限：1 ，只有super admin才能访问
        // 得到 resource 列表
        v1.GET("/customer/resource/list",           middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceList)
        // 增加一个resource group
        v1.POST("/customer/resource/addone",        middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resource/updateone",    middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resource/deleteone",   middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resource/deletebatch", middleware.PermissionLoginToken, middleware.SuperAdminRole, customerHandler.ResourceDeleteByIds)
        
        // #### customer role #### 权限：1，2 
        // 得到 role 列表
        v1.GET("/customer/role/list",               middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleList)
        // 增加一个role
        v1.POST("/customer/role/addone",            middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleAddOne)
        // 更新一个role
        v1.PATCH("/customer/role/updateone",        middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleUpdateById)
        // 删除一个role
        v1.DELETE("/customer/role/deleteone",       middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleDeleteById)
        // 批量删除role
        v1.DELETE("/customer/role/deletebatch",     middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleDeleteByIds)
        // 权限资源编辑页面的数据获取，得到全部的可用资源，以及勾选的权限
        v1.GET("/customer/role/resource/allandselected",   middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleResourceAllAndSelect)
        // 权限资源信息的更新
        v1.PATCH("/customer/role/resource/updateone",      middleware.PermissionLoginToken, middleware.CommonAdminRole, customerHandler.RoleResourceUpdate)
        
        
        // #### Common Market Group #### 权限：1, 2, 3
        // 得到 marketGroup 列表
        v1.GET("/common/marketgroup/list",           middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.MarketGroupList)
        // 增加一个 marketGroup
        v1.POST("/common/marketgroup/addone",        middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.MarketGroupAddOne)
        // 更新一个 marketGroup
        v1.PATCH("/common/marketgroup/updateone",    middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.MarketGroupUpdateById)
        // 删除一个 marketGroup
        v1.DELETE("/common/marketgroup/deleteone",   middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.MarketGroupDeleteById)
        // 批量删除 marketGroup
        v1.DELETE("/common/marketgroup/deletebatch", middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.MarketGroupDeleteByIds)
        
        
        // #### Common Market Group #### 权限：1, 2, 3
        // 得到 marketGroup 列表
        v1.GET("/common/channel/list",           middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.ChannelList)
        // 增加一个 marketGroup
        v1.POST("/common/channel/addone",        middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.ChannelAddOne)
        // 更新一个 marketGroup
        v1.PATCH("/common/channel/updateone",    middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.ChannelUpdateById)
        // 删除一个 marketGroup
        v1.DELETE("/common/channel/deleteone",   middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.ChannelDeleteById)
        // 批量删除 marketGroup
        v1.DELETE("/common/channel/deletebatch", middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.ChannelDeleteByIds)
        
        
        // #### Common Market Group #### 权限：1, 2, 3
        // 得到 marketGroup 列表
        v1.GET("/common/website/list",           middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.WebsiteList)
        // 增加一个 marketGroup
        v1.POST("/common/website/addone",        middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.WebsiteAddOne)
        // 更新一个 marketGroup
        v1.PATCH("/common/website/updateone",    middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.WebsiteUpdateById)
        // 删除一个 marketGroup
        v1.DELETE("/common/website/deleteone",   middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.WebsiteDeleteById)
        // 批量删除 marketGroup
        v1.DELETE("/common/website/deletebatch", middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.WebsiteDeleteByIds)
        
        // 得到 marketGroup 列表
        v1.GET("/common/website/jscode",         middleware.PermissionLoginToken, middleware.CommonAdminChildRole, commonHandler.WebsiteJsCode)
        
        // #### Basestics
        v1.GET("/whole/browser/list",            middleware.PermissionLoginToken, middleware.CommonAdminChildRole, wholeHandler.BrowserList)
        v1.GET("/whole/browser/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonAdminChildRole, wholeHandler.BrowserTrendInfo)
       
        
        
        
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
    // r.RunTLS(listenIp, "/etc/letsencrypt/live/fecshop.appfront.fancyecommerce.com/fullchain.pem", "/etc/letsencrypt/live/fecshop.appfront.fancyecommerce.com/privkey.pem") // 这里改成您的ip和端口
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