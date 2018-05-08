package model

import(
    // "github.com/fecshopsoft/fec-go/shell/whole"
    // "github.com/fecshopsoft/fec-go/util"
   
    // "github.com/globalsign/mgo/bson"
)



// type MapStringInt64 map[string]int64

type WholeBrowser struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeBrowserValue `form:"value" json:"value" bson:"value"`
}

type WholeBrowserValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName string `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
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


type WholeAll struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeAllValue `form:"value" json:"value" bson:"value"`
}

type WholeAllValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    CartCount int64 `form:"cart_count" json:"cart_count" bson:"cart_count"`
    OrderCount int64 `form:"order_count" json:"order_count" bson:"order_count"`
    OrderNoCount int64 `form:"order_no_count" json:"order_no_count" bson:"order_no_count"`
    OrderPaymentRate float64 `form:"order_payment_rate" json:"order_payment_rate" bson:"order_payment_rate"`
    
    SuccessOrderCount int64 `form:"success_order_count" json:"success_order_count" bson:"success_order_count"`
    SuccessOrderNoCount int64 `form:"success_order_no_count" json:"success_order_no_count" bson:"success_order_no_count"`
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
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    
    IpCount int64 `form:"ip_count" json:"ip_count" bson:"ip_count"`
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    Language map[string]int64 `form:"language" json:"language" bson:"language"`
    
    LoginEmailCount int64 `form:"login_email_count" json:"login_email_count" bson:"login_email_count"`
    RegisterEmailCount int64 `form:"register_email_count" json:"register_email_count" bson:"register_email_count"`
    
    OrderAmount float64 `form:"order_amount" json:"order_amount" bson:"order_amount"`
    SuccessOrderAmount float64 `form:"success_order_amount" json:"success_order_amount" bson:"success_order_amount"`
    
    CategoryCount int64 `form:"category_count" json:"category_count" bson:"category_count"`
    ProductCount int64 `form:"product_count" json:"product_count" bson:"product_count"`
    SearchCount int64 `form:"search_count" json:"search_count" bson:"search_count"`
    
}

type WholeRefer struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeReferValue `form:"value" json:"value" bson:"value"`
}
// first_referrer_domain
type WholeReferValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
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


type WholeCountry struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeCountryValue `form:"value" json:"value" bson:"value"`
}
// Country
type WholeCountryValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    CountryCode string `form:"country_code" json:"country_code" bson:"country_code"`
    CountryName string `form:"country_name" json:"country_name" bson:"country_name"`
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


type WholeDevide struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeDevideValue `form:"value" json:"value" bson:"value"`
}
// devide
type WholeDevideValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Devide string `form:"devide" json:"devide" bson:"devide"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
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


type WholeSku struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeSkuValue `form:"value" json:"value" bson:"value"`
}
// sku
type WholeSkuValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
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
    
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
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


type WholeSkuRefer struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeSkuReferValue `form:"value" json:"value" bson:"value"`
}
// sku
type WholeSkuReferValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Sku string `form:"sku" json:"sku" bson:"sku"`
    FirstReferrerDomain string `form:"first_referrer_domain" json:"first_referrer_domain" bson:"first_referrer_domain"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
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
    
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
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




type WholeSearch struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeSearchValue `form:"value" json:"value" bson:"value"`
}
// search_text
type WholeSearchValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    Language map[string]int64 `form:"language" json:"language" bson:"language"`
    IpCount int64 `form:"ip_count" json:"ip_count" bson:"ip_count"`
    
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
    FecApp map[string]int64 `form:"fec_app" json:"fec_app" bson:"fec_app"`
    IsReturn int64 `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage int64 `form:"first_page" json:"first_page" bson:"first_page"`
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    IsReturnRate float64 `form:"is_return_rate" json:"is_return_rate" bson:"is_return_rate"`
    StaySecondsRate float64 `form:"stay_seconds_rate" json:"stay_seconds_rate" bson:"stay_seconds_rate"`
    JumpOutRate float64 `form:"jump_out_rate" json:"jump_out_rate" bson:"jump_out_rate"`
    DropOutRate float64 `form:"drop_out_rate" json:"drop_out_rate" bson:"drop_out_rate"`
    PvRate float64 `form:"pv_rate" json:"pv_rate" bson:"pv_rate"`
    
    SearchText string `form:"search_text" json:"search_text" bson:"search_text"`
    SearchSkuClick int64 `form:"search_sku_click" json:"search_sku_click" bson:"search_sku_click"`
    SearchLoginEmail int64 `form:"search_login_email" json:"search_login_email" bson:"search_login_email"`
    SearchSkuCart int64 `form:"search_sku_cart" json:"search_sku_cart" bson:"search_sku_cart"`
    SearchSkuOrder int64 `form:"search_sku_order" json:"search_sku_order" bson:"search_sku_order"`
    SearchSkuOrderSuccess int64 `form:"search_sku_order_success" json:"search_sku_order_success" bson:"search_sku_order_success"`
    SearchQty int64 `form:"search_qty" json:"search_qty" bson:"search_qty"`
    
    SearchSkuClickRate float64 `form:"search_sku_click_rate" json:"search_sku_click_rate" bson:"search_sku_click_rate"`
    SearchSaleRate float64 `form:"search_sale_rate" json:"search_sale_rate" bson:"search_sale_rate"`
     
}




type WholeSearchLang struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeSearchLangValue `form:"value" json:"value" bson:"value"`
}
// search_text
type WholeSearchLangValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Language string `form:"language" json:"language" bson:"language"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    
    IpCount int64 `form:"ip_count" json:"ip_count" bson:"ip_count"`
    
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
    FecApp map[string]int64 `form:"fec_app" json:"fec_app" bson:"fec_app"`
    IsReturn int64 `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage int64 `form:"first_page" json:"first_page" bson:"first_page"`
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    IsReturnRate float64 `form:"is_return_rate" json:"is_return_rate" bson:"is_return_rate"`
    StaySecondsRate float64 `form:"stay_seconds_rate" json:"stay_seconds_rate" bson:"stay_seconds_rate"`
    JumpOutRate float64 `form:"jump_out_rate" json:"jump_out_rate" bson:"jump_out_rate"`
    DropOutRate float64 `form:"drop_out_rate" json:"drop_out_rate" bson:"drop_out_rate"`
    PvRate float64 `form:"pv_rate" json:"pv_rate" bson:"pv_rate"`
    
    SearchText string `form:"search_text" json:"search_text" bson:"search_text"`
    SearchSkuClick int64 `form:"search_sku_click" json:"search_sku_click" bson:"search_sku_click"`
    SearchLoginEmail int64 `form:"search_login_email" json:"search_login_email" bson:"search_login_email"`
    SearchSkuCart int64 `form:"search_sku_cart" json:"search_sku_cart" bson:"search_sku_cart"`
    SearchSkuOrder int64 `form:"search_sku_order" json:"search_sku_order" bson:"search_sku_order"`
    SearchSkuOrderSuccess int64 `form:"search_sku_order_success" json:"search_sku_order_success" bson:"search_sku_order_success"`
    SearchQty int64 `form:"search_qty" json:"search_qty" bson:"search_qty"`
    
    SearchSkuClickRate float64 `form:"search_sku_click_rate" json:"search_sku_click_rate" bson:"search_sku_click_rate"`
    SearchSaleRate float64 `form:"search_sale_rate" json:"search_sale_rate" bson:"search_sale_rate"`
     
}



type WholeUrl struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeUrlValue `form:"value" json:"value" bson:"value"`
}
// url
type WholeUrlValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Url string `form:"url" json:"url" bson:"url"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    Language map[string]int64 `form:"language" json:"language" bson:"language"`
    IpCount int64 `form:"ip_count" json:"ip_count" bson:"ip_count"`
    
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
    FecApp map[string]int64 `form:"fec_app" json:"fec_app" bson:"fec_app"`
    IsReturn int64 `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage int64 `form:"first_page" json:"first_page" bson:"first_page"`
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    IsReturnRate float64 `form:"is_return_rate" json:"is_return_rate" bson:"is_return_rate"`
    StaySecondsRate float64 `form:"stay_seconds_rate" json:"stay_seconds_rate" bson:"stay_seconds_rate"`
    JumpOutRate float64 `form:"jump_out_rate" json:"jump_out_rate" bson:"jump_out_rate"`
    DropOutRate float64 `form:"drop_out_rate" json:"drop_out_rate" bson:"drop_out_rate"`
    PvRate float64 `form:"pv_rate" json:"pv_rate" bson:"pv_rate"`
}



type WholeFirstUrl struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value WholeFirstUrlValue `form:"value" json:"value" bson:"value"`
}
// first url
type WholeFirstUrlValue struct{
    Id string `form:"id" json:"id" bson:"id"`
    Url string `form:"url" json:"url" bson:"url"`
    WebsiteId string `form:"website_id" json:"website_id" bson:"website_id"`
    BrowserName map[string]int64 `form:"browser_name" json:"browser_name" bson:"browser_name"`
    Pv int64 `form:"pv" json:"pv" bson:"pv"`
    Uv int64 `form:"uv" json:"uv" bson:"uv"`
    RatePv float64 `form:"rate_pv" json:"rate_pv" bson:"rate_pv"`
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    JumpOutCount int64 `form:"jump_out_count" json:"jump_out_count" bson:"jump_out_count"`
    DropOutCount int64 `form:"drop_out_count" json:"drop_out_count" bson:"drop_out_count"`
    CountryCode map[string]int64 `form:"country_code" json:"country_code" bson:"country_code"`
    
    Resolution map[string]int64 `form:"resolution" json:"resolution" bson:"resolution"`
    ColorDepth map[string]int64 `form:"color_depth" json:"color_depth" bson:"color_depth"`
    Language map[string]int64 `form:"language" json:"language" bson:"language"`
    IpCount int64 `form:"ip_count" json:"ip_count" bson:"ip_count"`
    
    Devide map[string]int64 `form:"devide" json:"devide" bson:"devide"`
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
    FecApp map[string]int64 `form:"fec_app" json:"fec_app" bson:"fec_app"`
    IsReturn int64 `form:"is_return" json:"is_return" bson:"is_return"`
    FirstPage int64 `form:"first_page" json:"first_page" bson:"first_page"`
    ServiceDateStr string `form:"service_date_str" json:"service_date_str" bson:"service_date_str"`
    IsReturnRate float64 `form:"is_return_rate" json:"is_return_rate" bson:"is_return_rate"`
    StaySecondsRate float64 `form:"stay_seconds_rate" json:"stay_seconds_rate" bson:"stay_seconds_rate"`
    JumpOutRate float64 `form:"jump_out_rate" json:"jump_out_rate" bson:"jump_out_rate"`
    DropOutRate float64 `form:"drop_out_rate" json:"drop_out_rate" bson:"drop_out_rate"`
    PvRate float64 `form:"pv_rate" json:"pv_rate" bson:"pv_rate"`
}






