package fec

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
)




// trace info
type TraceApiInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    
    Uuid string `binding:"required" form:"uuid" json:"uuid" bson:"uuid"`
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
    
    LoginEmail string `form:"login_email" json:"login_email" bson:"login_email"`
    RegisterEmail string `form:"register_email" json:"register_email" bson:"register_email"`
    PaymentPendingOrder OrderInfo `form:"payment_pending_order" json:"payment_pending_order" bson:"payment_pending_order"`
    PaymentSuccessOrder OrderInfo `form:"payment_success_order" json:"payment_success_order" bson:"payment_success_order"`
    
}

func (traceApiDbInfo TraceApiDbInfo) TableName() string {
    return "trace_info"
}
// trace info
type TraceApiDbInfo struct{
    Id_ bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    
    Uuid string `binding:"required" form:"uuid" json:"uuid" bson:"uuid"`
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
    
    LoginEmail string `form:"login_email" json:"login_email" bson:"login_email"`
    RegisterEmail string `form:"register_email" json:"register_email" bson:"register_email"`
    Order OrderInfo `form:"order" json:"order" bson:"order"`
    
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
    
    // traceApiInfo.PaymentSuccessOrder.Invoice
    var traceApiDbInfo TraceApiDbInfo
    
    // 进行数据的保存。
    err = mongodb.MC(traceApiDbInfo.TableName(), func(coll *mgo.Collection) error {
        // 如果传递了订单，那么将订单保存到这个变量中
        // var orderInfo OrderInfo
        var err error
        /*
        is_insert := 1
        if traceApiInfo.PaymentPendingOrder.Invoice != "" {
            invoice := traceApiInfo.PaymentPendingOrder.Invoice
            coll.Find(bson.M{"order.invoice": invoice}).One(&traceApiDbInfo)
            orderInfo = traceApiInfo.PaymentPendingOrder
        }
        */
        // 如果是成功订单，那么只更新订单的支付状态，其他的不变
        if traceApiInfo.PaymentSuccessOrder.Invoice != "" {
            invoice := traceApiInfo.PaymentSuccessOrder.Invoice
            payment_status := traceApiInfo.PaymentSuccessOrder.PaymentStatus
            
            selector := bson.M{"order.invoice": invoice}
            updateData := bson.M{"$set": bson.M{"order.payment_status": payment_status}}
            err = coll.Update(selector, updateData)
            return err
        } else {
            // 其他的则为插入
            traceApiDbInfo.Id_ = bson.NewObjectId()
            traceApiDbInfo.Uuid = traceApiInfo.Uuid
            traceApiDbInfo.ClActivity = traceApiInfo.ClActivity
            traceApiDbInfo.ClActivityChild = traceApiInfo.ClActivityChild
            traceApiDbInfo.FirstReferrerDomain = traceApiInfo.FirstReferrerDomain
            traceApiDbInfo.FirstPage = traceApiInfo.FirstPage
            traceApiDbInfo.FirstReferrerUrl = traceApiInfo.FirstReferrerUrl
            
            traceApiDbInfo.IsReturn = traceApiInfo.IsReturn
            traceApiDbInfo.WebsiteId = traceApiInfo.WebsiteId
            traceApiDbInfo.Fid = traceApiInfo.Fid
            traceApiDbInfo.FecSource = traceApiInfo.FecSource
            traceApiDbInfo.FecMedium = traceApiInfo.FecMedium
            traceApiDbInfo.FecCampaign = traceApiInfo.FecCampaign
            traceApiDbInfo.FecContent = traceApiInfo.FecContent
            traceApiDbInfo.FecDesign = traceApiInfo.FecDesign
            
            traceApiDbInfo.LoginEmail = traceApiInfo.LoginEmail
            traceApiDbInfo.RegisterEmail = traceApiInfo.RegisterEmail
            traceApiDbInfo.Order = traceApiInfo.PaymentPendingOrder
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