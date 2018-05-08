package fec

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/url"
    // "errors"
    "log"
    "encoding/json"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "github.com/fecshopsoft/fec-go/helper"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
)



// trace info 用于从Get中获取数据。
type TraceGetInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    Uuid string `binding:"required" form:"uuid" json:"uuid" bson:"uuid"`
    WebsiteId string `binding:"required" form:"website_id" json:"website_id" bson:"website_id"` 
    Devide string `form:"devide" json:"devide" bson:"devide"`
    UserAgent string `form:"user_agent" json:"user_agent" bson:"user_agent"`
    BrowserName string `form:"browser_name" json:"browser_name" bson:"browser_name"`
    BrowserVersion string `form:"browser_version" json:"browser_version" bson:"browser_version"`
    BrowserDate string `form:"browser_date" json:"browser_date" bson:"browser_date"`
    BrowserLang string `form:"browser_lang" json:"browser_lang" bson:"browser_lang"`
    Operate string `form:"operate" json:"operate" bson:"operate"`
    OperateRelase string `form:"operate_relase" json:"operate_relase" bson:"operate_relase"`
    Url string `form:"url" json:"url" bson:"url"`
    Domain string `form:"domain" json:"domain" bson:"domain"`
    Title string `form:"title" json:"title" bson:"title"`
    ReferUrl string `form:"refer_url" json:"refer_url" bson:"refer_url"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    FirstReferrerUrl string `form:"first_referrer_url" json:"first_referrer_url" bson:"first_referrer_url"`
    ClActivity string `form:"cl_activity" json:"cl_activity" bson:"cl_activity"`
    ClActivityChild string `form:"cl_activity_child" json:"cl_activity_child" bson:"cl_activity_child"`
    IsReturn string `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage string `form:"first_page" json:"first_page" bson:"first_page"`
    DevicePixelRatio string `form:"device_pixel_ratio" json:"device_pixel_ratio" bson:"device_pixel_ratio"`
    Resolution string `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth string `form:"color_depth" json:"color_depth" bson:"color_depth"`
    /**
     * fid  广告id
     * fec_source   渠道
     * fec_medium   子渠道
     * fec_campaign 活动 
     * fec_content  推广员
     * fec_design   广告设计员
     */
    Fid string `form:"fid" json:"fid" bson:"fid"`
    FecSource string `form:"fec_source" json:"fec_source" bson:"fec_source"`
    FecMedium string `form:"fec_medium" json:"fec_medium" bson:"fec_medium"`
    FecCampaign string `form:"fec_campaign" json:"fec_campaign" bson:"fec_campaign"`
    FecContent string `form:"fec_content" json:"fec_content" bson:"fec_content"`
    FecDesign string `form:"fec_design" json:"fec_design" bson:"fec_design"`
    
    FecStore string `form:"fec_store" json:"fec_store" bson:"fec_store"`
    FecLang string `form:"fec_lang" json:"fec_lang" bson:"fec_lang"`
    FecApp string `form:"fec_app" json:"fec_app" bson:"fec_app"`
    FecCurrency string `form:"fec_currency" json:"fec_currency" bson:"fec_currency"`
    
    
    
    Category string `form:"category" json:"category" bson:"category"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    
    
    
    // Cart string `form:"cart" json:"cart" bson:"cart"`
    Search string `form:"search" json:"search" bson:"search"`
}




// trace info 用于将数据保存到数据库
type TraceInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    Uuid string `binding:"required" form:"uuid" json:"uuid" bson:"uuid"`
    
    Ip string `form:"ip" json:"ip" bson:"ip"`
    CountryCode string `form:"country_code" json:"country_code" bson:"country_code"`
    CountryName string `form:"country_name" json:"country_name" bson:"country_name"`
    StateName string `form:"state_name" json:"state_name" bson:"state_name"`
    CityName string `form:"city_name" json:"city_name" bson:"city_name"`
    
    WebsiteId string `binding:"required" form:"website_id" json:"website_id" bson:"website_id"` 
    Devide string `form:"devide" json:"devide" bson:"devide"`
    UserAgent string `form:"user_agent" json:"user_agent" bson:"user_agent"`
    BrowserName string `form:"browser_name" json:"browser_name" bson:"browser_name"`
    BrowserVersion string `form:"browser_version" json:"browser_version" bson:"browser_version"`
    BrowserDate string `form:"browser_date" json:"browser_date" bson:"browser_date"`
    BrowserLang string `form:"browser_lang" json:"browser_lang" bson:"browser_lang"`
    Operate string `form:"operate" json:"operate" bson:"operate"`
    OperateRelase string `form:"operate_relase" json:"operate_relase" bson:"operate_relase"`
    Url string `form:"url" json:"url" bson:"url"`
    Domain string `form:"domain" json:"domain" bson:"domain"`
    Title string `form:"title" json:"title" bson:"title"`
    ReferUrl string `form:"refer_url" json:"refer_url" bson:"refer_url"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    FirstReferrerUrl string `form:"first_referrer_url" json:"first_referrer_url" bson:"first_referrer_url"`
    ClActivity string `form:"cl_activity" json:"cl_activity" bson:"cl_activity"`
    ClActivityChild string `form:"cl_activity_child" json:"cl_activity_child" bson:"cl_activity_child"`
    IsReturn string `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage int `form:"first_page" json:"first_page" bson:"first_page"`
    DevicePixelRatio string `form:"device_pixel_ratio" json:"device_pixel_ratio" bson:"device_pixel_ratio"`
    Resolution string `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth string `form:"color_depth" json:"color_depth" bson:"color_depth"`
    /**
     * fid  广告id
     * fec_source   渠道
     * fec_medium   子渠道
     * fec_campaign 活动 
     * fec_content  推广员
     * fec_design   广告设计员
     */
    Fid string `form:"fid" json:"fid" bson:"fid"`
    FecSource string `form:"fec_source" json:"fec_source" bson:"fec_source"`
    FecMedium string `form:"fec_medium" json:"fec_medium" bson:"fec_medium"`
    FecCampaign string `form:"fec_campaign" json:"fec_campaign" bson:"fec_campaign"`
    FecContent string `form:"fec_content" json:"fec_content" bson:"fec_content"`
    FecDesign string `form:"fec_design" json:"fec_design" bson:"fec_design"`
    
    FecStore string `form:"fec_store" json:"fec_store" bson:"fec_store"`
    FecLang string `form:"fec_lang" json:"fec_lang" bson:"fec_lang"`
    FecApp string `form:"fec_app" json:"fec_app" bson:"fec_app"`
    FecCurrency string `form:"fec_currency" json:"fec_currency" bson:"fec_currency"`
    
    // 附加字段 service_timestamp
    // 服务器接收数据的时间戳
    ServiceTimestamp int64 `form:"service_timestamp" json:"service_timestamp" bson:"service_timestamp"`
    // 服务器接收数据, 格式：Y-m-d H:i:s
    ServiceDatetime string `form:"service_datetime" json:"service_datetime" bson:"service_datetime"`
    // 服务器接收数据, 格式：Y-m-d
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    // 页面停留时间
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    // 由于按照时间分库，站点分表，查询当前表，是否存在uuid，如果不存在，则 uuid_first_page = 1，否则 uuid_first_page = 0
    UuidFirstPage int `form:"uuid_first_page" json:"uuid_first_page" bson:"uuid_first_page"`
    // Ip First Page ，类似上面的 uuid_first_page
    IpFirstPage int `form:"ip_first_page" json:"ip_first_page" bson:"ip_first_page"`
    // uuid 
    UuidFirstCategory int `form:"uuid_first_category" json:"uuid_first_category" bson:"uuid_first_category"`
    //
    IpFirstCategory int `form:"ip_first_category" json:"ip_first_category" bson:"ip_first_category"`
    // 去掉某些参数后的url
    UrlNew string `form:"url_new" json:"url_new" bson:"url_new"`
    // 登录后访问搜索页面的用户
    SearchLoginEmail int `form:"search_login_email" json:"search_login_email" bson:"search_login_email"`
    // 首次访问某个url的时候，标记为1
    FirstVisitThisUrl int `form:"first_visit_this_url" json:"first_visit_this_url" bson:"first_visit_this_url"`
    
    
    Category string `form:"category" json:"category" bson:"category"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    SearchSkuClick int `form:"search_sku_click" json:"search_sku_click" bson:"search_sku_click"`
    // SearchSkuCart map[string]map[string]int  `form:"search_sku_cart" json:"search_sku_cart" bson:"search_sku_cart"`
    
    // Cart []CartItem `form:"cart" json:"cart" bson:"cart"`
    Search SearchInfo `form:"search" json:"search" bson:"search"`
}

// 为了中间变量的生成，而进行查询
type TraceMiddInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    Uuid string `binding:"required" form:"uuid" json:"uuid" bson:"uuid"`
    Ip string `binding:"required" form:"ip" json:"ip" bson:"ip"`
    // 附加字段 service_timestamp
    // 服务器接收数据的时间戳
    ServiceTimestamp int64 `form:"service_timestamp" json:"service_timestamp" bson:"service_timestamp"`
    // 服务器接收数据, 格式：Y-m-d H:i:s
    ServiceDatetime string `form:"service_datetime" json:"service_datetime" bson:"service_datetime"`
    // 服务器接收数据, 格式：Y-m-d
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    // 页面停留时间
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    // 由于按照时间分库，站点分表，查询当前表，是否存在uuid，如果不存在，则 uuid_first_page = 1，否则 uuid_first_page = 0
    UuidFirstPage int `form:"uuid_first_page" json:"uuid_first_page" bson:"uuid_first_page"`
    // Ip First Page ，类似上面的 uuid_first_page
    IpFirstPage int `form:"ip_first_page" json:"ip_first_page" bson:"ip_first_page"`
    // uuid 
    UuidFirstCategory int `form:"uuid_first_category" json:"uuid_first_category" bson:"uuid_first_category"`
    //
    IpFirstCategory int `form:"ip_first_category" json:"ip_first_category" bson:"ip_first_category"`
    // 去掉某些参数后的url
    UrlNew string `form:"url_new" json:"url_new" bson:"url_new"`
    // 登录后访问搜索页面的用户
    SearchLoginEmail int `form:"search_login_email" json:"search_login_email" bson:"search_login_email"`
    // 首次访问某个url的时候，标记为1
    FirstVisitThisUrl int `form:"first_visit_this_url" json:"first_visit_this_url" bson:"first_visit_this_url"`
    Url string `form:"url" json:"url" bson:"url"`
    Search SearchInfo `form:"search" json:"search" bson:"search"`
    SearchSkuClick int `form:"search_sku_click" json:"search_sku_click" bson:"search_sku_click"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
}


// search
type SearchInfo struct{
    Text string `form:"text" json:"text" bson:"text" json:"text"`
    ResultQty int64 `form:"result_qty" json:"result_qty" bson:"result_qty" json:"result_qty"`
}

func (traceInfo TraceInfo) TableName() string {
    return "trace_info"
}

func SaveJsData(c *gin.Context){
    var traceGetInfo TraceGetInfo
    // query := c.Request.URL.Query()
    err := c.ShouldBindQuery(&traceGetInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    var traceInfo TraceInfo
    // 赋值 将 traceGetInfo 的值 赋值给 traceInfo
    // search - 搜索
    var searchInfo SearchInfo
    traceGetInfo.Search, _ = url.QueryUnescape(traceGetInfo.Search)
    err = json.Unmarshal([]byte(traceGetInfo.Search), &searchInfo)
    // url decode
    // searchInfo.Text, _ = url.QueryUnescape(searchInfo.Text)
    traceInfo.Search = searchInfo
    // cart - 购物车
    // var cartItems []CartItem
    // traceGetInfo.Cart, _ = url.QueryUnescape(traceGetInfo.Cart)
    // err = json.Unmarshal([]byte(traceGetInfo.Cart), &cartItems)
    // traceInfo.Cart = cartItems
    // 其他变量赋值
    traceInfo.Uuid, _ = url.QueryUnescape(traceGetInfo.Uuid)
    traceInfo.WebsiteId, _ = url.QueryUnescape(traceGetInfo.WebsiteId)
    traceInfo.Devide, _ = url.QueryUnescape(traceGetInfo.Devide)
    traceInfo.UserAgent, _ = url.QueryUnescape(traceGetInfo.UserAgent)
    traceInfo.BrowserName, _ = url.QueryUnescape(traceGetInfo.BrowserName)
    traceInfo.BrowserVersion, _ = url.QueryUnescape(traceGetInfo.BrowserVersion)
    traceInfo.BrowserDate, _ = url.QueryUnescape(traceGetInfo.BrowserDate)
    traceInfo.BrowserLang, _ = url.QueryUnescape(traceGetInfo.BrowserLang)
    traceInfo.Operate, _ = url.QueryUnescape(traceGetInfo.Operate)
    traceInfo.OperateRelase, _ = url.QueryUnescape(traceGetInfo.OperateRelase)
    traceInfo.Url, _ = url.QueryUnescape(traceGetInfo.Url)
    traceInfo.Domain, _ = url.QueryUnescape(traceGetInfo.Domain)
    traceInfo.Title, _ = url.QueryUnescape(traceGetInfo.Title)
    traceInfo.ReferUrl, _ = url.QueryUnescape(traceGetInfo.ReferUrl)
    traceInfo.FirstReferrerDomain, _ = url.QueryUnescape(traceGetInfo.FirstReferrerDomain)
    traceInfo.FirstReferrerUrl, _ = url.QueryUnescape(traceGetInfo.FirstReferrerUrl)
    traceInfo.ClActivity, _ = url.QueryUnescape(traceGetInfo.ClActivity)
    traceInfo.ClActivityChild, _ = url.QueryUnescape(traceGetInfo.ClActivityChild)
    traceInfo.IsReturn, _ = url.QueryUnescape(traceGetInfo.IsReturn)
    firstPage, _ := url.QueryUnescape(traceGetInfo.FirstPage)
    traceInfo.FirstPage, _ = helper.Int(firstPage)
    
    traceInfo.DevicePixelRatio, _ = url.QueryUnescape(traceGetInfo.DevicePixelRatio)
    traceInfo.Resolution, _ = url.QueryUnescape(traceGetInfo.Resolution)
    traceInfo.ColorDepth, _ = url.QueryUnescape(traceGetInfo.ColorDepth)
    
    traceInfo.Fid, _ = url.QueryUnescape(traceGetInfo.Fid)
    traceInfo.FecSource, _ = url.QueryUnescape(traceGetInfo.FecSource)
    traceInfo.FecMedium, _ = url.QueryUnescape(traceGetInfo.FecMedium)
    traceInfo.FecCampaign, _ = url.QueryUnescape(traceGetInfo.FecCampaign)
    traceInfo.FecContent, _ = url.QueryUnescape(traceGetInfo.FecContent)
    traceInfo.FecDesign, _ = url.QueryUnescape(traceGetInfo.FecDesign)
    // category
    traceInfo.Category, _ = url.QueryUnescape(traceGetInfo.Category)
    
    // sku
    traceInfo.Sku, _ = url.QueryUnescape(traceGetInfo.Sku)
    
    traceInfo.FecStore, _ = url.QueryUnescape(traceGetInfo.FecStore)
    traceInfo.FecLang, _ = url.QueryUnescape(traceGetInfo.FecLang)
    traceInfo.FecApp, _ = url.QueryUnescape(traceGetInfo.FecApp)
    traceInfo.FecCurrency, _ = url.QueryUnescape(traceGetInfo.FecCurrency)
    
    ipStr := c.ClientIP()
    countryCode, countryName, stateName, cityName, err := helper.GetIpInfo(ipStr) 
    if err == nil {
        // 如果获取国家报错，仅仅丢弃这几个字段，其它的还是要接收
        traceInfo.CountryCode = countryCode
        traceInfo.CountryName = countryName
        traceInfo.StateName   = stateName
        traceInfo.CityName    = cityName
    }
    traceInfo.Ip = ipStr
    
    // 得到 dbName 和 collName
    dbName := helper.GetTraceDbName()
    collName := helper.GetTraceDataCollName(traceInfo.WebsiteId)
    
    traceInfo.ServiceTimestamp = helper.DateTimestamps()
    traceInfo.ServiceDatetime = helper.DateTimeUTCStr()
    traceInfo.ServiceDateStr = helper.DateUTCStr()
    // 计算出来的属性
    // StaySeconds 查找最近的一次访问，时间差，就是最近一次访问的停留时间
    err = updatePreStaySeconds(dbName, collName, traceInfo.Uuid, traceInfo.ServiceTimestamp)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // UuidFirstPage
    traceInfo.UuidFirstPage, err = getUuidFirstPage(dbName, collName, traceInfo.Uuid)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // IpFirstPage
    traceInfo.IpFirstPage, err = getIpFirstPage(dbName, collName, ipStr)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    if traceInfo.Category != "" {
        // UuidFirstCategory
        traceInfo.UuidFirstCategory, err = getUuidFirstCategory(dbName, collName, traceInfo.Uuid, traceInfo.Category)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        // IpFirstCategory
        traceInfo.IpFirstCategory, err = getIpFirstCategory(dbName, collName, ipStr, traceInfo.Category)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        
    }
    // UrlNew
    traceInfo.UrlNew = getUrlNew(traceInfo.Url)
    // SearchLoginEmail
    if traceInfo.Search.Text != "" {
        // SearchLoginEmail
        traceInfo.SearchLoginEmail, err = getSearchLoginEmail(dbName, collName, traceInfo.Uuid)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
    }
    // FirstVisitThisUrl - first_visit_this_url
    traceInfo.FirstVisitThisUrl, err = getFirstVisitThisUrl(dbName, collName, traceInfo.Uuid, traceInfo.UrlNew)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    traceInfo.ReferUrl,_ = url.QueryUnescape(traceInfo.ReferUrl)
    log.Println("ReferUrl:")
    log.Println(traceInfo.ReferUrl);
    // 如果是产品页面，
    if traceInfo.Sku != "" {
        referUrl := traceInfo.ReferUrl
        
        log.Println("referUrl:" + referUrl)
        if (referUrl != "" && helper.StrContains(referUrl, "/catalogsearch/index?q=")) || (referUrl !="" && helper.StrContains(referUrl, "/#/search/")) {
            log.Println("if success")
            searchInfo, _ := getBeforeSearchOne(dbName, collName, traceInfo.Uuid, traceInfo.ReferUrl)
            log.Println("if success searchInfo")
            log.Println(searchInfo.Text)
            if searchInfo.Text != "" {
                log.Println("if success searchInfo ###")
                log.Println(searchInfo)
                
                traceInfo.Search = searchInfo
                traceInfo.SearchSkuClick = 1
            }
        }
    }
    
    // 进行保存。
    
    err = mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        traceInfo.Id_ = bson.NewObjectId()
        err := coll.Insert(traceInfo)
        return err
    })
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "traceInfo": traceInfo,
        // "query": query,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}
    
    
// 初始化mongodb表，以及表索引。
func InitTraceDataCollIndex() error{
    // 得到所有的 active websiteId
    dateInt64 := helper.DateTimestamps()
    var daySec int64 = 24 * 3600
    var dInt64 int64
    var err error
    // websiteId := "9b17f5b4-b96f-46fd-abe6-a579837ccdd9"
    // 得到今天以及未来15天的日期
    var j int64
    websiteInfos, err := commonHandler.GetAllActiveWebsiteId()
    if err != nil {
        return err
    }
    // var websiteIds []int64
    for i:=0; i<len(websiteInfos); i++ {
        websiteInfo := websiteInfos[i]
        websiteId := websiteInfo.SiteUid 
        // 最长五天。
        for j=0; j<5; j++ {
            dInt64 = dateInt64 + j * daySec
            dateStr := helper.GetDateTimeUtcByTimestamps(dInt64)
            err = createIndex(dateStr, websiteId)
            if err != nil {
                return err
            }
        }
    }
    
    return err
}

// 创建表索引
func createIndex(dateStr string, websiteId string) error{
    dbName := helper.GetTraceDbNameByDate(dateStr)
    collName := helper.GetTraceDataCollName(websiteId)
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var err error
        // 下面的每一个子项，就是一个索引。
        mgoIndex := [][]string{
            []string{"order.invoice"},
            []string{"uuid", "_id"},
            []string{"ip"},
            []string{"uuid", "service_timestamp"},
        }
        for i:=0; i<len(mgoIndex); i++ {
            index := mgo.Index{
                Key: mgoIndex[i],
                // Unique: true,
                // DropDups: true,
                Background: true, // See notes.
                // Sparse: true,
            }
            err = coll.EnsureIndex(index)
            if err != nil {
                return err
            }
        }
        // lastIndexes, err := coll.Indexes() // 查看表索引
        // if err != nil {
        //     return err
        // }
        // log.Println(lastIndexes)
        return err
    })
    return err
}

