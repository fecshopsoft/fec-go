package fec

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/url"
    "encoding/json"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
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
    
    Category string `form:"category" json:"category" bson:"category"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    
    
    Cart string `form:"cart" json:"cart" bson:"cart"`
    Search string `form:"search" json:"search" bson:"search"`
}

// trace info 用于将数据保存到数据库
type TraceInfo struct{
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
    
    
    Category string `form:"category" json:"category" bson:"category"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    
    Cart []CartItem `form:"cart" json:"cart" bson:"cart"`
    Search SearchInfo `form:"search" json:"search" bson:"search"`
}
// cart
type CartItem struct{
    Sku string `form:"sku" json:"sku" bson:"sku"`
    Qty int64 `form:"qty" json:"qty" bson:"qty"`
    Price float64 `form:"price" json:"price" bson:"price"`
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
    var cartItems []CartItem
    traceGetInfo.Cart, _ = url.QueryUnescape(traceGetInfo.Cart)
    err = json.Unmarshal([]byte(traceGetInfo.Cart), &cartItems)
    traceInfo.Cart = cartItems
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
    traceInfo.FirstPage, _ = url.QueryUnescape(traceGetInfo.FirstPage)
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
    
    
    // 进行保存。
    err = mongodb.MC(traceInfo.TableName(), func(coll *mgo.Collection) error {
        // c.Find(bson.M{"_id": id}).One(traceInfo)
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



/*

*/