package customerModel

import(
    "github.com/globalsign/mgo/bson"
)

// type MapStringInt64 map[string]int64
type UuidCustomer struct{
    Id_  bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    CustomerId string `form:"customer_id" json:"customer_id" bson:"customer_id"`
    Uuids []string `form:"uuids" json:"uuids" bson:"uuids"`
    Emails []string `form:"emails" json:"emails" bson:"emails"`
    UpdatedAt int64 `form:"updated_at" json:"updated_at" bson:"updated_at"`
}

type UuidCustomerEmail struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value UuidCustomerEmailValue `form:"value" json:"value" bson:"value"`
}

type UuidCustomerEmailValue struct{
    Email string `form:"email" json:"email" bson:"email"`
    Count string `form:"count" json:"count" bson:"count"`
}


type CustomerUuid struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value CustomerUuidValue `form:"value" json:"value" bson:"value"`
}

type CustomerUuidValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    
    Uuid string `form:"uuid" json:"uuid" bson:"uuid"`
    CustomerId string `form:"customer_id" json:"customer_id" bson:"customer_id"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    
    RegisterEmail string `form:"register_email" json:"register_email" bson:"register_email"`
    LoginEmail string `form:"login_email" json:"login_email" bson:"login_email"`
    CustomerEmail []string `form:"customer_email" json:"customer_email" bson:"customer_email"`
    
    Fid map[string]int64 `form:"fid" json:"fid" bson:"fid"`
    FecContent map[string]int64 `form:"fec_content" json:"fec_content" bson:"fec_content"`
    FecMarketGroup map[string]int64 `form:"fec_market_group" json:"fec_market_group" bson:"fec_market_group"`
    FecCampaign map[string]int64 `form:"fec_campaign" json:"fec_campaign" bson:"fec_campaign"`
    FecSource map[string]int64 `form:"fec_source" json:"fec_source" bson:"fec_source"`
    FecMedium map[string]int64 `form:"fec_medium" json:"fec_medium" bson:"fec_medium"`
    FecDesign map[string]int64 `form:"fec_design" json:"fec_design" bson:"fec_design"`
    
    FidMain string `form:"fid_main" json:"fid_main" bson:"fid_main"`
    FecContentMain string `form:"fec_content_main" json:"fec_content_main" bson:"fec_content_main"`
    FecMarketGroupMain string `form:"fec_market_group_main" json:"fec_market_group_main" bson:"fec_market_group_main"`
    FecCampaignMain string `form:"fec_campaign_main" json:"fec_campaign_main" bson:"fec_campaign_main"`
    FecSourceMain string `form:"fec_source_main" json:"fec_source_main" bson:"fec_source_main"`
    FecMediumMain string `form:"fec_medium_main" json:"fec_medium_main" bson:"fec_medium_main"`
    FecDesignMain string `form:"fec_design_main" json:"fec_design_main" bson:"fec_design_main"`
    
    
    Sku map[string]int64 `form:"sku" json:"sku" bson:"sku"`
    SkuCart map[string]int64 `form:"sku_cart" json:"sku_cart" bson:"sku_cart"`
    SkuOrder map[string]int64 `form:"sku_order" json:"sku_order" bson:"sku_order"`
    SkuOrderSuccess map[string]int64 `form:"sku_order_success" json:"sku_order_success" bson:"sku_order_success"`
    Category map[string]int64 `form:"category" json:"category" bson:"category"`
    
    Search map[string]int64 `form:"search" json:"search" bson:"search"`
    
    Cart []CartItem `form:"cart" json:"cart" bson:"cart"`
    Order []OrderInfo `form:"order" json:"order" bson:"order"`
    
    VisitPageSku int64 `form:"visit_page_sku" json:"visit_page_sku" bson:"visit_page_sku"`
    VisitPageCategory int64 `form:"visit_page_category" json:"visit_page_category" bson:"visit_page_category"`
    VisitPageSearch int64 `form:"visit_page_search" json:"visit_page_search" bson:"visit_page_search"`
    VisitPageCart int64 `form:"visit_page_cart" json:"visit_page_cart" bson:"visit_page_cart"`
    VisitPageOrder int64 `form:"visit_page_order" json:"visit_page_order" bson:"visit_page_order"`
    VisitPageOrderAmount float64 `form:"visit_page_order_amount" json:"visit_page_order_amount" bson:"visit_page_order_amount"`
    VisitPageOrderProcessing int64 `form:"visit_page_order_processing" json:"visit_page_order_processing" bson:"visit_page_order_processing"`
    VisitPageOrderProcessingAmount float64 `form:"visit_page_order_processing_amount" json:"visit_page_order_processing_amount" bson:"visit_page_order_processing_amount"`
    VisitPageOrderPending int64 `form:"visit_page_order_pending" json:"visit_page_order_pending" bson:"visit_page_order_pending"`
    VisitPageOrderPendingAmount float64 `form:"visit_page_order_pending_amount" json:"visit_page_order_pending_amount" bson:"visit_page_order_pending_amount"`
    
    Ip map[string]int64 `form:"ip" json:"ip" bson:"ip"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
    FecApp map[string]int64 `form:"fec_app" json:"fec_app" bson:"fec_app"`
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    Language map[string]int64 `form:"language" json:"language" bson:"language"`
    
    IpMain string `form:"ip_main" json:"ip_main" bson:"ip_main"`
    BrowserNameMain string `form:"browser_name_main" json:"browser_name_main" bson:"browser_name_main"`
    DevideMain string `form:"devide_main" json:"devide_main" bson:"devide_main"`
    CountryCodeMain string `form:"country_code_main" json:"country_code_main" bson:"country_code_main"`
    OperateMain string `form:"operate_main" json:"operate_main" bson:"operate_main"`
    FecAppMain string `form:"fec_app_main" json:"fec_app_main" bson:"fec_app_main"`
    ResolutionMain string `form:"resolution_main" json:"resolution_main" bson:"resolution_main"`
    ColorDepthMain string `form:"color_depth_main" json:"color_depth_main" bson:"color_depth_main"`
    LanguageMain string `form:"language_main" json:"language_main" bson:"language_main"`
    
    Domain string `form:"domain" json:"domain" bson:"domain"`
    
    ReferUrl string `form:"refer_url" json:"refer_url" bson:"refer_url"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    IsReturn int `form:"is_return" json:"is_return" bson:"is_return"`
    
    FirstPageUrl string `form:"first_page_url" json:"first_page_url" bson:"first_page_url"`
    OutPage string `form:"out_page" json:"out_page" bson:"out_page"`
    DevicePixelRatio string `form:"device_pixel_ratio" json:"device_pixel_ratio" bson:"device_pixel_ratio"`
    Data []CustomerUuidData `form:"data" json:"data" bson:"data"`
}

type CustomerUuidData struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Ip string `form:"ip" json:"ip" bson:"ip"`
    CountryCode string `form:"country_code" json:"country_code" bson:"country_code"`
    CountryName string `form:"country_name" json:"country_name" bson:"country_name"`
    ServiceDatetime string `form:"service_datetime" json:"service_datetime" bson:"service_datetime"`
    Devide string `form:"devide" json:"devide" bson:"devide"`
    Uuid string `form:"uuid" json:"uuid" bson:"uuid"`
    Fid string `form:"fid" json:"fid" bson:"fid"`
    FecContent string `form:"fec_content" json:"fec_content" bson:"fec_content"`
    FecMarketGroup string `form:"fec_market_group" json:"fec_market_group" bson:"fec_market_group"`
    FecCampaign string `form:"fec_campaign" json:"fec_campaign" bson:"fec_campaign"`
    FecSource string `form:"fec_source" json:"fec_source" bson:"fec_source"`
    FecMedium string `form:"fec_medium" json:"fec_medium" bson:"fec_medium"`
    FecDesign string `form:"fec_design" json:"fec_design" bson:"fec_design"`
    
    DevicePixelRatio string `form:"device_pixel_ratio" json:"device_pixel_ratio" bson:"device_pixel_ratio"`
    IsReturn int `form:"is_return" json:"is_return" bson:"is_return"`
    UserAgent string `form:"user_agent" json:"user_agent" bson:"user_agent"`
    
    BrowserName string `form:"browser_name" json:"browser_name" bson:"browser_name"`
    BrowserVersion string `form:"browser_version" json:"browser_version" bson:"browser_version"`
    BrowserDate string `form:"browser_date" json:"browser_date" bson:"browser_date"`
    BrowserLang string `form:"browser_lang" json:"browser_lang" bson:"browser_lang"`
    Operate string `form:"operate" json:"operate" bson:"operate"`
    OperateRelase string `form:"operate_relase" json:"operate_relase" bson:"operate_relase"`
    Domain string `form:"domain" json:"domain" bson:"domain"`
    Url string `form:"url" json:"url" bson:"url"`
    Title string `form:"title" json:"title" bson:"title"`
    ReferUrl string `form:"refer_url" json:"refer_url" bson:"refer_url"`
    
    FecApp string `form:"fec_app" json:"fec_app" bson:"fec_app"`
    Language string `form:"language" json:"language" bson:"language"`
    
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    Resolution string `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth string `form:"color_depth" json:"color_depth" bson:"color_depth"`
    FirstPage string `form:"first_page" json:"first_page" bson:"first_page"`
    UrlNew string `form:"url_new" json:"url_new" bson:"url_new"`
    LoginEmail string `form:"login_email" json:"login_email" bson:"login_email"`
    RegisterEmail string `form:"register_email" json:"register_email" bson:"register_email"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    Category string `form:"category" json:"category" bson:"category"`
    Search SearchInfo `form:"search" json:"search" bson:"search"`
    Cart []CartItem `form:"cart" json:"cart" bson:"cart"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    Order OrderInfo `form:"order" json:"order" bson:"order"`
}



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


// search
type SearchInfo struct{
    Text string `form:"text" json:"text" bson:"text" json:"text"`
    ResultQty int64 `form:"result_qty" json:"result_qty" bson:"result_qty" json:"result_qty"`
}

// cart
type CartItem struct{
    Sku string `form:"sku" json:"sku" bson:"sku"`
    Qty int64 `form:"qty" json:"qty" bson:"qty"`
    Price float64 `form:"price" json:"price" bson:"price"`
}
