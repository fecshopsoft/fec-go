package model

import(
    // "github.com/fecshopsoft/fec-go/shell/whole"
    // "github.com/fecshopsoft/fec-go/util"
    handlerFec "github.com/fecshopsoft/fec-go/handler/fec"
    // "github.com/globalsign/mgo/bson"
)

// type MapStringInt64 map[string]int64

type AdvertiseFid struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value AdvertiseFidValue `form:"value" json:"value" bson:"value"`
}

type AdvertiseFidValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Fid string `form:"fid" json:"fid" bson:"fid"`
    FecMarketGroup string `form:"fec_market_group" json:"fec_market_group" bson:"fec_market_group"`
    FecContent string `form:"fec_content" json:"fec_content" bson:"fec_content"`
    FecSource string `form:"fec_source" json:"fec_source" bson:"fec_source"`
    FecDesign string `form:"fec_design" json:"fec_design" bson:"fec_design"`
    
    FecCampaign map[string]int64 `form:"fec_campaign" json:"fec_campaign" bson:"fec_campaign"`
    FecMedium map[string]int64 `form:"fec_medium" json:"fec_medium" bson:"fec_medium"`
    
    FecMediumMain string `form:"fec_medium_main" json:"fec_medium_main" bson:"fec_medium_main"`
    SuccessOrderCAllUvRate float64 `form:"success_order_c_all_uv_rate" json:"success_order_c_all_uv_rate" bson:"success_order_c_all_uv_rate"`
    SuccessOrderCSuccessUvRate float64 `form:"success_order_c_success_uv_rate" json:"success_order_c_success_uv_rate" bson:"success_order_c_success_uv_rate"`
    
    
    
    FirstReferrerDomain map[string]int64 `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    SkuVisitInfo map[string]int64 `form:"sku_visit_info" json:"sku_visit_info" bson:"sku_visit_info"`
    CategoryVisitInfo map[string]int64 `form:"category_visit_info" json:"category_visit_info" bson:"category_visit_info"`
    SearchVisitInfo map[string]int64 `form:"search_visit_info" json:"search_visit_info" bson:"search_visit_info"`
    
    CartSkuInfo map[string]int64 `form:"cart_sku_info" json:"cart_sku_info" bson:"cart_sku_info"`
    OrderSkuInfo map[string]int64 `form:"order_sku_info" json:"order_sku_info" bson:"order_sku_info"`
    SuccessOrderSkuInfo map[string]int64 `form:"success_order_sku_info" json:"success_order_sku_info" bson:"success_order_sku_info"`
    OrderIncrementId []string `form:"order_increment_id" json:"order_increment_id" bson:"order_increment_id"`
    FailOrderIncrementId map[string]int64 `form:"fail_order_increment_id" json:"fail_order_increment_id" bson:"fail_order_increment_id"`
    SuccessOrderIncrementId map[string]int64 `form:"success_order_increment_id" json:"success_order_increment_id" bson:"success_order_increment_id"`
    SuccessOrderInfo []handlerFec.OrderInfo `form:"success_order_info" json:"success_order_info" bson:"success_order_info"`
    FailOrderInfo []handlerFec.OrderInfo `form:"fail_order_info" json:"fail_order_info" bson:"fail_order_info"`
    
    RegisterCount int64 `form:"register_count" json:"register_count" bson:"register_count"`
    LoginCount int64 `form:"login_count" json:"login_count" bson:"login_count"`
    CategoryCount int64 `form:"category_count" json:"category_count" bson:"category_count"`
    SkuCount int64 `form:"sku_count" json:"sku_count" bson:"sku_count"`
    SearchCount int64 `form:"search_count" json:"search_count" bson:"search_count"`    
    
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    CartCount int64 `form:"cart_count" json:"cart_count" bson:"cart_count"`
    OrderCount int64 `form:"order_count" json:"order_count" bson:"order_count"`
    OrderNoCount int64 `form:"order_no_count" json:"order_no_count" bson:"order_no_count"`
    OrderPaymentRate float64 `form:"order_payment_rate" json:"order_payment_rate" bson:"order_payment_rate"`
    OrderAmount float64 `form:"order_amount" json:"order_amount" bson:"order_amount"`
    SuccessOrderAmount float64 `form:"success_order_amount" json:"success_order_amount" bson:"success_order_amount"`
    
    SuccessOrderCount int64 `form:"success_order_count" json:"success_order_count" bson:"success_order_count"`
    SuccessOrderNoCount int64 `form:"success_order_no_count" json:"success_order_no_count" bson:"success_order_no_count"`
    
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    Language map[string]int64 `form:"language" json:"language" bson:"language"`
    IpCount int64 `form:"ip_count" json:"ip_count" bson:"ip_count"`
    
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
    FecApp map[string]int64 `form:"fec_app" json:"fec_app" bson:"fec_app"`
    IsReturn int64 `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage int64 `form:"first_page" json:"first_page" bson:"first_page"`
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    IsReturnRate float64 `form:"is_return_rate" json:"is_return_rate" bson:"is_return_rate"`
    StaySecondsRate float64 `form:"stay_seconds_rate" json:"stay_seconds_rate" bson:"stay_seconds_rate"`
    JumpOutRate float64 `form:"jump_out_rate" json:"jump_out_rate" bson:"jump_out_rate"`
    DropOutRate float64 `form:"drop_out_rate" json:"drop_out_rate" bson:"drop_out_rate"`
    SkuSaleRate float64 `form:"sku_sale_rate" json:"sku_sale_rate" bson:"sku_sale_rate"`
    PvRate float64 `form:"pv_rate" json:"pv_rate" bson:"pv_rate"`
}