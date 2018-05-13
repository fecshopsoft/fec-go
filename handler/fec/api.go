package fec

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/gin-gonic/gin"
    "net/http"
    "errors"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "github.com/fecshopsoft/fec-go/initialization"
)


// trace info
type TraceApiInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    
    Uuid string `form:"uuid" json:"uuid" bson:"uuid"`
    ClActivity string `form:"cl_activity" json:"cl_activity" bson:"cl_activity"`
    ClActivityChild string `form:"cl_activity_child" json:"cl_activity_child" bson:"cl_activity_child"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    FirstPage string `form:"first_page" json:"first_page" bson:"first_page"`
    FirstReferrerUrl string `form:"first_referrer_url" json:"first_referrer_url" bson:"first_referrer_url"`
    IsReturn string `form:"is_return" json:"is_return" bson:"is_return"`
    WebsiteId string `binding:"required" form:"website_id" json:"website_id" bson:"website_id"` 
    
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
    FecMarketGroup string `form:"fec_market_group" json:"fec_market_group" bson:"fec_market_group"`
    
    FecStore string `form:"fec_store" json:"fec_store" bson:"fec_store"`
    FecLang string `form:"fec_lang" json:"fec_lang" bson:"fec_lang"`
    FecApp string `form:"fec_app" json:"fec_app" bson:"fec_app"`
    FecCurrency string `form:"fec_currency" json:"fec_currency" bson:"fec_currency"`
    
    LoginEmail string `form:"login_email" json:"login_email" bson:"login_email"`
    RegisterEmail string `form:"register_email" json:"register_email" bson:"register_email"`
    PaymentPendingOrder OrderInfo `form:"payment_pending_order" json:"payment_pending_order" bson:"payment_pending_order"`
    PaymentSuccessOrder OrderInfo `form:"payment_success_order" json:"payment_success_order" bson:"payment_success_order"`
    
    Cart []CartItem `form:"cart" json:"cart" bson:"cart"`
}

// cart
type CartItem struct{
    Sku string `form:"sku" json:"sku" bson:"sku"`
    Qty int64 `form:"qty" json:"qty" bson:"qty"`
    Price float64 `form:"price" json:"price" bson:"price"`
}

func (traceApiDbInfo TraceApiDbInfo) TableName() string {
    return "trace_info"
}
// trace info
type TraceApiDbInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    
    Uuid string `form:"uuid" json:"uuid" bson:"uuid"`
    ClActivity string `form:"cl_activity" json:"cl_activity" bson:"cl_activity"`
    ClActivityChild string `form:"cl_activity_child" json:"cl_activity_child" bson:"cl_activity_child"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    FirstPage int `form:"first_page" json:"first_page" bson:"first_page"`
    FirstReferrerUrl string `form:"first_referrer_url" json:"first_referrer_url" bson:"first_referrer_url"`
    IsReturn string `form:"is_return" json:"is_return" bson:"is_return"`
    WebsiteId string `binding:"required" form:"website_id" json:"website_id" bson:"website_id"` 
    
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
    FecMarketGroup string `form:"fec_market_group" json:"fec_market_group" bson:"fec_market_group"`
    
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
    
    
    LoginEmail string `form:"login_email" json:"login_email" bson:"login_email"`
    RegisterEmail string `form:"register_email" json:"register_email" bson:"register_email"`
    Order OrderInfo `form:"order" json:"order" bson:"order"`
    Cart []CartItem `form:"cart" json:"cart" bson:"cart"`
    Ip string `form:"ip" json:"ip" bson:"ip"`
    CountryCode string `form:"country_code" json:"country_code" bson:"country_code"`
    CountryName string `form:"country_name" json:"country_name" bson:"country_name"`
    StateName string `form:"state_name" json:"state_name" bson:"state_name"`
    CityName string `form:"city_name" json:"city_name" bson:"city_name"`
    
    Devide string `form:"devide" json:"devide" bson:"devide"`
    UserAgent string `form:"user_agent" json:"user_agent" bson:"user_agent"`
    BrowserName string `form:"browser_name" json:"browser_name" bson:"browser_name"`
    BrowserVersion string `form:"browser_version" json:"browser_version" bson:"browser_version"`
    BrowserDate string `form:"browser_date" json:"browser_date" bson:"browser_date"`
    BrowserLang string `form:"browser_lang" json:"browser_lang" bson:"browser_lang"`
    Operate string `form:"operate" json:"operate" bson:"operate"`
    OperateRelase string `form:"operate_relase" json:"operate_relase" bson:"operate_relase"`
    
    Domain string `form:"domain" json:"domain" bson:"domain"`
    ReferUrl string `form:"refer_url" json:"refer_url" bson:"refer_url"`

    DevicePixelRatio string `form:"device_pixel_ratio" json:"device_pixel_ratio" bson:"device_pixel_ratio"`
    Resolution string `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth string `form:"color_depth" json:"color_depth" bson:"color_depth"`
    
    SearchSkuCart map[string]map[string]int  `form:"search_sku_cart" json:"search_sku_cart" bson:"search_sku_cart"`
    SearchSkuOrder map[string]map[string]int  `form:"search_sku_order" json:"search_sku_order" bson:"search_sku_order"`
    // SearchSkuOrderSuccess map[string]map[string]int  `form:"search_sku_order_success" json:"search_sku_order_success" bson:"search_sku_order_success"`
    
}

// 对于api传递的 payment_pending_order 和 payment_success_order ，来设定支付状态
// order
type OrderInfo struct{
    Invoice string `form:"invoice" json:"invoice" bson:"invoice"`
    OrderType string `form:"order_type" json:"order_type" bson:"order_type"`
    // 未支付：payment_pending，  已支付：payment_confirmed  
    PaymentStatus string `form:"payment_status" json:"payment_status" bson:"payment_status"`
    PaymentType string `form:"payment_type" json:"payment_type" bson:"payment_type"`
    Amount float64 `form:"amount" json:"amount" bson:"amount"`
    Shipping float64 `form:"shipping" json:"shipping" bson:"shipping"`
    Discount_amount float64 `form:"discount_amount" json:"discount_amount" bson:"discount_amount"`
    Coupon string `form:"coupon" json:"coupon" bson:"coupon"`
    City string `form:"city" json:"city" bson:"city"`
    Email string `form:"email" json:"email" bson:"email"`
    FirstName string `form:"first_name" json:"first_name" bson:"first_name"`
    LastName string `form:"last_name" json:"last_name" bson:"last_name"`
    Zip string `form:"zip" json:"zip" bson:"zip"`
    CountryCode string `form:"country_code" json:"country_code" bson:"country_code"`
    StateCode string `form:"state_code" json:"state_code" bson:"state_code"`
    // StateCode string `form:"state_code" json:"state_code" bson:"state_code"`
    CreatedAt int64 `form:"created_at" json:"created_at" bson:"created_at"`
    CountryName string `form:"country_name" json:"country_name" bson:"country_name"`
    StateName string `form:"state_name" json:"state_name" bson:"state_name"`
    Address1 string `form:"address1" json:"address1" bson:"address1"`
    Address2 string `form:"address2" json:"address2" bson:"address2"`
    Products []OrderProduct `form:"products" json:"products" bson:"products"`
    // pending   complete 两种支付状态。
    // PaymentStatus string `form:"payment_status" json:"payment_status" bson:"payment_status"`
    
}
// order products
type OrderProduct struct{
    Sku string `form:"sku" json:"sku" bson:"sku"`
    Name string `form:"name" json:"name" bson:"name"`
    Qty int64 `form:"qty" json:"qty" bson:"qty"`
    Price float64 `form:"price" json:"price" bson:"price"`
}

func SaveApiData(c *gin.Context){
    var traceApiInfo TraceApiInfo
    err := c.ShouldBindJSON(&traceApiInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 从access-token解析出来的website_id 和 参数传递的website_id做比较，如果不一致，则报错。
    token_website_id := helper.GetCurrentWebsiteId(c)
    if traceApiInfo.WebsiteId != token_website_id {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("access-token(" + token_website_id + ") is not Inconsistent with website_id(" + traceApiInfo.WebsiteId + ")"))
        return
    }
     
    if  traceApiInfo.Uuid == "" {
        // uuid为空的情况，只有订单状态更新的情况，如果下面的三个字段，有一个为空，则退出
        if traceApiInfo.PaymentSuccessOrder.Invoice == "" ||  traceApiInfo.PaymentSuccessOrder.CreatedAt == 0 || traceApiInfo.PaymentSuccessOrder.PaymentStatus == "" {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(" uuid can not empty (if you update order payment status, [Invoice, CreatedAt, PaymentStatus] is required)"))
            return
        } else if traceApiInfo.PaymentSuccessOrder.PaymentStatus != "payment_pending" &&  traceApiInfo.PaymentSuccessOrder.PaymentStatus != "payment_confirmed" {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(" order payment status must in ['payment_pending', 'payment_confirmed']"))
            return
        }
        
    }
    
    // traceApiInfo.PaymentSuccessOrder.Invoice
    var traceApiDbInfo TraceApiDbInfo
    
    // 得到db name
    var dbName string
    if traceApiInfo.PaymentSuccessOrder.Invoice == "" {
        dbName = helper.GetTraceDbName()
    } else {
        OrderCreatedAt := traceApiInfo.PaymentSuccessOrder.CreatedAt
        dateStr := helper.GetDateTimeUtcByTimestamps(OrderCreatedAt)
        dbName = helper.GetTraceDbNameByDate(dateStr)
    }
    // 得到collection name
    collName := helper.GetTraceDataCollName(traceApiInfo.WebsiteId)
    
    if traceApiInfo.Uuid != "" && traceApiInfo.PaymentSuccessOrder.Invoice == "" {
        // 除了订单状态更新，其他的都是插入数据，也就是下面的赋值
        traceApiDbInfo.Id_ = bson.NewObjectId()
        traceApiDbInfo.Uuid = traceApiInfo.Uuid
        traceApiDbInfo.ClActivity = traceApiInfo.ClActivity
        traceApiDbInfo.ClActivityChild = traceApiInfo.ClActivityChild
        traceApiDbInfo.FirstReferrerDomain = traceApiInfo.FirstReferrerDomain
        firstPage := traceApiInfo.FirstPage
        traceApiDbInfo.FirstPage, _ = helper.Int(firstPage)
        
        traceApiDbInfo.FirstReferrerUrl = traceApiInfo.FirstReferrerUrl
        
        traceApiDbInfo.IsReturn = traceApiInfo.IsReturn
        traceApiDbInfo.WebsiteId = traceApiInfo.WebsiteId
        traceApiDbInfo.Fid = traceApiInfo.Fid
        traceApiDbInfo.FecSource = traceApiInfo.FecSource
        traceApiDbInfo.FecMedium = traceApiInfo.FecMedium
        traceApiDbInfo.FecCampaign = traceApiInfo.FecCampaign
        traceApiDbInfo.FecContent = traceApiInfo.FecContent
        traceApiDbInfo.FecDesign = traceApiInfo.FecDesign
        
        // traceApiDbInfo.FecMarketGroup = traceApiInfo.FecMarketGroup
        customerIdWithMarketGroup := initialization.CustomerIdWithMarketGroup
        customerId, err := helper.Int64(traceApiDbInfo.FecContent)
        if err == nil && customerId != 0 {
            if mgId, ok := customerIdWithMarketGroup[customerId]; ok {
                traceApiDbInfo.FecMarketGroup = helper.Str64(mgId)
            }
        }
        
        traceApiDbInfo.FecStore = traceApiInfo.FecStore
        traceApiDbInfo.FecLang = traceApiInfo.FecLang
        traceApiDbInfo.FecApp = traceApiInfo.FecApp
        traceApiDbInfo.FecCurrency = traceApiInfo.FecCurrency
        
        traceApiDbInfo.LoginEmail = traceApiInfo.LoginEmail
        traceApiDbInfo.RegisterEmail = traceApiInfo.RegisterEmail
        traceApiDbInfo.Order = traceApiInfo.PaymentPendingOrder
        traceApiDbInfo.Cart = traceApiInfo.Cart
        
        //  ##############
        
        traceApiDbInfo.ServiceTimestamp = helper.DateTimestamps()
        traceApiDbInfo.ServiceDatetime = helper.DateTimeUTCStr()
        traceApiDbInfo.ServiceDateStr = helper.DateUTCStr()
        
        // 计算出来的属性
        // StaySeconds
        preTraceInfo, err := updatePreStaySecondsAndReturn(dbName, collName, traceApiDbInfo.Uuid, traceApiDbInfo.ServiceTimestamp)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        
        // UuidFirstPage
        traceApiDbInfo.UuidFirstPage, err = getUuidFirstPage(dbName, collName, traceApiDbInfo.Uuid)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return
        }
        
        // 其他的字符，通过查询最近的数据，如果有数据，则将上一条的数据，附加到当前
        if preTraceInfo.Uuid != "" {
            // 不为空，则说明有数据，保存
            traceApiDbInfo.Ip = preTraceInfo.Ip
            traceApiDbInfo.CountryCode = preTraceInfo.CountryCode
            traceApiDbInfo.CountryName = preTraceInfo.CountryName
            traceApiDbInfo.StateName = preTraceInfo.StateName
            traceApiDbInfo.CityName = preTraceInfo.CityName
            traceApiDbInfo.Devide = preTraceInfo.Devide
            traceApiDbInfo.UserAgent = preTraceInfo.UserAgent
            traceApiDbInfo.BrowserName = preTraceInfo.BrowserName
            traceApiDbInfo.BrowserVersion = preTraceInfo.BrowserVersion
            traceApiDbInfo.BrowserDate = preTraceInfo.BrowserDate
            traceApiDbInfo.BrowserLang = preTraceInfo.BrowserLang
            traceApiDbInfo.Operate = preTraceInfo.Operate
            traceApiDbInfo.OperateRelase = preTraceInfo.OperateRelase
            traceApiDbInfo.Domain = preTraceInfo.Domain
            traceApiDbInfo.ReferUrl = preTraceInfo.ReferUrl
            traceApiDbInfo.FirstReferrerDomain = preTraceInfo.FirstReferrerDomain
            traceApiDbInfo.FirstReferrerUrl = preTraceInfo.FirstReferrerUrl
            traceApiDbInfo.DevicePixelRatio = preTraceInfo.DevicePixelRatio
            traceApiDbInfo.Resolution = preTraceInfo.Resolution
            traceApiDbInfo.ColorDepth = preTraceInfo.ColorDepth
        }
        // 如果是订单数据，那么更新 SearchSkuOrder
        
        if traceApiDbInfo.Order.Invoice != "" &&  len(traceApiDbInfo.Order.Products) > 0 {
            products := traceApiDbInfo.Order.Products
            searchSkuOrder := make(map[string]map[string]int)
            for i:=0; i<len(products); i++ {
                product := products[i]
                sku := product.Sku
                searchInfo, _ := getBeforeCartOne(dbName, collName, traceApiDbInfo.Uuid, sku)
                if searchInfo.Text != "" {
                    skuQ := make(map[string]int)
                    searchText := ClearDianChar(searchInfo.Text)
                    
                    if _,ok := searchSkuOrder[searchText]; ok {
                        skuQ = searchSkuOrder[searchText]
                    } else {
                        skuQ = make(map[string]int)
                    }
                    skuQ[sku] = 1
                    searchSkuOrder[searchText] = skuQ
                    
                }
            }
            traceApiDbInfo.SearchSkuOrder = searchSkuOrder
        }
        
        // 如果是购物车页面
        if len(traceApiDbInfo.Cart) > 0 {
            cartData := traceApiDbInfo.Cart
            searchSkuCart := make(map[string]map[string]int)
            for i:=0; i<len(cartData); i++ {
                item := cartData[i]
                sku := item.Sku
                searchInfo, _ := getBeforeCartOne(dbName, collName, traceApiDbInfo.Uuid, sku)
                if searchInfo.Text != "" {
                    var skuQ  map[string]int
                    searchText := ClearDianChar(searchInfo.Text)
                    if _,ok := searchSkuCart[searchText]; ok {
                        skuQ = searchSkuCart[searchText]
                    } else {
                        skuQ = make(map[string]int)
                    }
                    skuQ[sku] = 1
                    searchSkuCart[searchText] = skuQ
                }
            }
            traceApiDbInfo.SearchSkuCart = searchSkuCart
        }
    }
    
    //  ##############
    
            
            
    err = mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        // 如果传递了订单，那么将订单保存到这个变量中
        // var orderInfo OrderInfo
        var err error
        // 如果是成功订单，那么只更新订单的支付状态，其他的不变
        if traceApiInfo.PaymentSuccessOrder.Invoice != "" {
            invoice := traceApiInfo.PaymentSuccessOrder.Invoice
            payment_status := traceApiInfo.PaymentSuccessOrder.PaymentStatus
            if invoice == "" {
                return errors.New("invoice can not empty")
            }
            // invoice两个作为条件查询
            selector := bson.M{"order.invoice": invoice}
            updateData := bson.M{"$set": bson.M{"order.payment_status": payment_status}}
            err = coll.Update(selector, updateData)
            return err
        } else {
            if traceApiInfo.Uuid == "" {
                return errors.New("uuid can not empty")
            } 
            
            err = coll.Insert(traceApiDbInfo)
            return err
        }
        // 进行赋值。 // bsonObjectID := bson.ObjectIdHex("573ce4451e02f4bae78788aa")
    })
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "traceApiInfo": traceApiInfo,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}



