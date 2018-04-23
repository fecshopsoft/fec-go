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
    SuccessOrderCount int64 `form:"success_order_count" json:"success_order_count" bson:"success_order_count"`
    SuccessOrderNoCount int64 `form:"success_order_no_count" json:"success_order_no_count" bson:"success_order_no_count"`
    Operate map[string]int64 `form:"operate" json:"operate" bson:"operate"`
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

