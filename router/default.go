package router
/**
 * 监听实现部分
 * 1.initialization初始化数据
 * 2.设置404 NotFound Router
 * 3.初始化cors
 * 4.初始化上下文中的全局变量
 * 5.设置传递和接收数据的router
 * 6.刷新全局变量的：/fec/trace/cronssss
 * 7.vue端的请求的router
 * 8.监听ip。
 */
import(
    "github.com/gin-gonic/gin"
    customerHandler "github.com/fecshopsoft/fec-go/handler/customer"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
    fecHandler "github.com/fecshopsoft/fec-go/handler/fec"
    testHandler "github.com/fecshopsoft/fec-go/handler/test"
    cronHandler "github.com/fecshopsoft/fec-go/handler/cron"
    wholeHandler "github.com/fecshopsoft/fec-go/handler/whole"
    advertiseHandler "github.com/fecshopsoft/fec-go/handler/advertise"
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
    // 初始化全局变量
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
        v1.GET("/customer/list",                middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerList)
        // 增加一个用户
        v1.POST("/customer/addone",             middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerAddOne)
        // 更新一个用户
        v1.PATCH("/customer/updateone",         middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerUpdateById)
        // 更新一个用户的付费信息
        //v1.PATCH("/customer/updateonepayinfo",         middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerUpdatePayInfoById)
        v1.GET("/customer/role/allandselected", middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerRoleAllAndSelect)
        v1.PATCH("/customer/role/updaterelate",    middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerRoleUpdate)
        
        // 删除一个用户
        v1.DELETE("/customer/deleteone",        middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerDeleteById)
        // 批量删除用户
        v1.DELETE("/customer/deletebatch",      middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.CustomerDeleteByIds)
        
		
        // 得到 customer 列表  role: 1,2,3
        //v1.GET("/customer/child/list",                middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerChildList)
        // 增加一个用户 role: 1,2
        //v1.POST("/customer/child/addone",             middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerChildAddOne)
        // 更新一个用户 role: 1,2
        //v1.PATCH("/customer/child/updateone",         middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerChildUpdateById)
        // 删除一个用户 role: 1,2
        //v1.DELETE("/customer/child/deleteone",        middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerChildDeleteById)
        // 批量删除用户 role: 1,2
        //v1.DELETE("/customer/child/deletebatch",      middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerChildDeleteByIds)
        
        //v1.GET("/customer/child/role/allandselected", middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.CustomerRoleAllAndSelect)
        // 更新customer 的 roles
        
        
        // #### customer resource group #### 权限：1 ，只有super admin才能访问
        // 得到 resource group 列表
        v1.GET("/customer/resourcegroup/list",              middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceGroupList)
        // 增加一个resource group
        v1.POST("/customer/resourcegroup/addone",           middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceGroupAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resourcegroup/updateone",       middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceGroupUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resourcegroup/deleteone",      middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceGroupDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resourcegroup/deletebatch",    middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceGroupDeleteByIds)
        
        // #### customer resource #### 权限：1 ，只有super admin才能访问
        // 得到 resource 列表
        v1.GET("/customer/resource/list",           middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceList)
        // 增加一个resource group
        v1.POST("/customer/resource/addone",        middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceAddOne)
        // 更新一个resource group
        v1.PATCH("/customer/resource/updateone",    middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceUpdateById)
        // 删除一个resource group
        v1.DELETE("/customer/resource/deleteone",   middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceDeleteById)
        // 批量删除resource group
        v1.DELETE("/customer/resource/deletebatch", middleware.PermissionLoginToken, middleware.AdminRole, customerHandler.ResourceDeleteByIds)
        
        // #### customer role #### 权限：1，2 
        // 得到 role 列表
        v1.GET("/customer/role/list",               middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleList)
        // 增加一个role
        v1.POST("/customer/role/addone",            middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleAddOne)
        // 更新一个role
        v1.PATCH("/customer/role/updateone",        middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleUpdateById)
        // 删除一个role
        v1.DELETE("/customer/role/deleteone",       middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleDeleteById)
        // 批量删除role
        v1.DELETE("/customer/role/deletebatch",     middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleDeleteByIds)
        // 权限资源编辑页面的数据获取，得到全部的可用资源，以及勾选的权限
        v1.GET("/customer/role/resource/allandselected",   middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleResourceAllAndSelect)
        // 权限资源信息的更新
        v1.PATCH("/customer/role/resource/updateone",      middleware.PermissionLoginToken, middleware.CommonRole, customerHandler.RoleResourceUpdate)
        
        
        // #### Common Market Group #### 权限：1, 2, 3
        // 得到 marketGroup 列表
        v1.GET("/common/marketgroup/list",           middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.MarketGroupList)
        // 增加一个 marketGroup
        v1.POST("/common/marketgroup/addone",        middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.MarketGroupAddOne)
        // 更新一个 marketGroup
        v1.PATCH("/common/marketgroup/updateone",    middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.MarketGroupUpdateById)
        // 删除一个 marketGroup
        v1.DELETE("/common/marketgroup/deleteone",   middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.MarketGroupDeleteById)
        // 批量删除 marketGroup
        v1.DELETE("/common/marketgroup/deletebatch", middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.MarketGroupDeleteByIds)
        
        
        // #### Common Market Group #### 权限：1, 2, 3
        // 得到 marketGroup 列表
        v1.GET("/common/channel/list",           middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.ChannelList)
        // 增加一个 marketGroup
        v1.POST("/common/channel/addone",        middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.ChannelAddOne)
        // 更新一个 marketGroup
        v1.PATCH("/common/channel/updateone",    middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.ChannelUpdateById)
        // 删除一个 marketGroup
        v1.DELETE("/common/channel/deleteone",   middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.ChannelDeleteById)
        // 批量删除 marketGroup
        v1.DELETE("/common/channel/deletebatch", middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.ChannelDeleteByIds)
        
        
        // #### Common Market Group #### 权限：1, 2, 3
        // 得到 marketGroup 列表
        v1.GET("/common/website/list",           middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.WebsiteList)
        // 增加一个 marketGroup
        v1.POST("/common/website/addone",        middleware.PermissionLoginToken, middleware.AdminRole, commonHandler.WebsiteAddOne)
        // 更新一个 marketGroup
        v1.PATCH("/common/website/updateone",    middleware.PermissionLoginToken, middleware.AdminRole, commonHandler.WebsiteUpdateById)
        // 删除一个 marketGroup
        v1.DELETE("/common/website/deleteone",   middleware.PermissionLoginToken, middleware.AdminRole, commonHandler.WebsiteDeleteById)
        // 批量删除 marketGroup
        v1.DELETE("/common/website/deletebatch", middleware.PermissionLoginToken, middleware.AdminRole, commonHandler.WebsiteDeleteByIds)
        
        // 得到 marketGroup 列表
        v1.GET("/common/website/jscode",         middleware.PermissionLoginToken, middleware.CommonRole, commonHandler.WebsiteJsCode)
        
        // #### Whole Site
        v1.GET("/whole/site/list",              middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SiteList)
        v1.GET("/whole/site/fetchtrendinfo",    middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SiteTrendInfo)
        // #### Whole Browser
        v1.GET("/whole/browser/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.BrowserList)
        v1.GET("/whole/browser/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.BrowserTrendInfo)
        // #### Whole Refer
        v1.GET("/whole/refer/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.ReferList)
        v1.GET("/whole/refer/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.ReferTrendInfo)
        // #### Whole Country
        v1.GET("/whole/country/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.CountryList)
        v1.GET("/whole/country/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.CountryTrendInfo)
        // #### Whole Devide
        v1.GET("/whole/devide/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.DevideList)
        v1.GET("/whole/devide/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.DevideTrendInfo)
        // #### Whole Sku
        v1.GET("/whole/sku/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SkuList)
        v1.GET("/whole/sku/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SkuTrendInfo)
        // #### Whole Sku Refer
        v1.GET("/whole/skurefer/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SkuReferList)
        v1.GET("/whole/skurefer/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SkuReferTrendInfo)
        // #### Whole Search
        v1.GET("/whole/search/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SearchList)
        v1.GET("/whole/search/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SearchTrendInfo)
        // #### Whole Search Lang
        v1.GET("/whole/searchlang/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SearchLangList)
        v1.GET("/whole/searchlang/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.SearchLangTrendInfo)
        
        // #### Whole Url
        v1.GET("/whole/url/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.UrlList)
        v1.GET("/whole/url/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.UrlTrendInfo)
        // #### Whole First Url
        v1.GET("/whole/firsturl/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.FirstUrlList)
        v1.GET("/whole/firsturl/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.FirstUrlTrendInfo)
        // #### Whole Category
        v1.GET("/whole/category/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.CategoryList)
        v1.GET("/whole/category/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.CategoryTrendInfo)
        // #### Whole App
        v1.GET("/whole/app/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.AppList)
        v1.GET("/whole/app/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.AppTrendInfo)
        // #### Whole Advertise
        v1.GET("/whole/advertise/init",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.AdvertiseInit)
        v1.GET("/whole/advertise/generateurl",     middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.AdvertiseGenerateUrl)
        v1.GET("/whole/advertise/list",            middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.AdvertiseList)
        v1.GET("/whole/advertise/download/mutilxlsx", wholeHandler.AdvertiseDownloadMutilXlsx)
        v1.POST("/whole/advertise/generatemutiladvertise", middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.GenerateMutilAdvertise)
        v1.POST("/whole/advertise/generatemutillinkadvertise", middleware.PermissionLoginToken, middleware.CommonRole, wholeHandler.GenerateMutilLinkAdvertise)
        
        // #### Advertise Fid
        v1.GET("/advertise/fid/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.FidList)
        v1.GET("/advertise/fid/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.FidTrendInfo)
        
        // #### Advertise Content
        v1.GET("/advertise/content/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.ContentList)
        v1.GET("/advertise/content/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.ContentTrendInfo)
        
        // #### Advertise MarketGroup
        v1.GET("/advertise/marketgroup/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.MarketGroupList)
        v1.GET("/advertise/marketgroup/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.MarketGroupTrendInfo)
        
        // #### Advertise Design
        v1.GET("/advertise/design/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.DesignList)
        v1.GET("/advertise/design/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.DesignTrendInfo)
        
        // #### Advertise Campaign
        v1.GET("/advertise/campaign/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.CampaignList)
        v1.GET("/advertise/campaign/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.CampaignTrendInfo)
        
        // #### Advertise Medium
        v1.GET("/advertise/medium/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.MediumList)
        v1.GET("/advertise/medium/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.MediumTrendInfo)
        
        // #### Advertise Source
        v1.GET("/advertise/source/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.SourceList)
        v1.GET("/advertise/source/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.SourceTrendInfo)
        
        // #### Advertise Source
        v1.GET("/advertise/edm/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.EdmList)
        v1.GET("/advertise/edm/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.EdmTrendInfo)
        
        // #### Customer Uuid
        v1.GET("/customer/uuid/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.UuidList)
        v1.GET("/customer/uuid/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.UuidTrendInfo)
        v1.GET("/customer/uuid/one",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.UuidOne)
        
        // #### Whole Url
        v1.GET("/advertise/eid/list",            middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.EidList)
        v1.GET("/advertise/eid/fetchtrendinfo",  middleware.PermissionLoginToken, middleware.CommonRole, advertiseHandler.EidTrendInfo)
       
       
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

















