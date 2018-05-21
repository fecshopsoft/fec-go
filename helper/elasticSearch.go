package helper

import (
   
   
)

var numberOfShards string   = "10"
var numberOfReplicas string = "1"
// 将mongodb数据同步到es，批量插入的数据条数
var BulkSyncCount = 200
// 在vue中趋势数据展示多少天的数据
var TrendDays int64 = 30

// 废弃：得到当前时间对应的 Es   的 Index Name
func GetEsIndexName(websiteId string) (string){
    return "trace_" + websiteId
}
// 通过typename，得到index，现在是一个type，一个index。
func GetEsIndexNameByType(typeName string) (string){
    return "trace_" + typeName     //+ "_" + websiteId
}

func GetEsIndexMapping() string {
    return `
        {
            "settings":{
                "number_of_shards": ` + numberOfShards + `,
                "number_of_replicas": ` + numberOfReplicas + `
            }
        }
    `
}

// Whole Browser Type Name
func GetEsWholeBrowserTypeName() (string){
    return "whole_browser_data"
}
// Whole Browser Type Mapping
// https://github.com/olivere/elastic/issues/755
func GetEsWholeBrowserTypeMapping() (string){
    return `{
		"whole_browser_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "browser_id":       {"type":"keyword"},
                "browser_name":     {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
    /*
    return `
        {
            "settings":{
                "number_of_shards": ` + numberOfShards + `,
                "number_of_replicas": ` + numberOfReplicas + `
            },
            "mappings":{
                "whole_browser_data":{
                    "properties":{
                        "browser_id":       {"type":"keyword"},
                        "browser_name":     {"type":"keyword"},
                        "pv":               {"type":"integer"},
                        "uv":               {"type":"integer"},
                        "jump_out_count":   {"type":"integer"},
                        "drop_out_count":   {"type":"integer"},
                        "stay_seconds":     {"type":"integer"},
                        "devide":           {"type":"keyword"},
                        "country_code":     {"type":"keyword"},
                        "operate":          {"type":"keyword"},
                        "is_return":        {"type":"integer"},
                        "first_page":       {"type":"integer"},
                        "service_date_str": {"type":"date"},
                        "stay_seconds_rate":{"type":"float"},
                        "jump_out_rate":    {"type":"float"},
                        "drop_out_rate":    {"type":"float"},
                        "is_return_rate":   {"type":"float"},
                        "pv_rate":          {"type":"float"},
                        "sku_sale_rate":    {"type":"float"},
                        "cart_count":               {"type":"integer"},
                        "order_count":              {"type":"integer"},
                        "success_order_count":      {"type":"integer"},
                        "success_order_no_count":   {"type":"integer"},
                    }
                }
            }
        }
    `
    */
}

// Whole All Type Name
func GetEsWholeAllTypeName() (string){
    return "whole_all_data"
}

// Whole All Type Mapping
func GetEsWholeAllTypeMapping() (string){
    return `{
		"whole_all_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":         {"type":"float"},
                "success_order_amount": {"type":"float"},
                
                "ip_count":             {"type":"integer"},
                "login_email_count":    {"type":"integer"},
                "register_email_count": {"type":"integer"},
                "category_count":       {"type":"integer"},
                "product_count":        {"type":"integer"},
                "search_count":         {"type":"integer"}
            }
        }
	}`
}


// Whole Refer Type Name
func GetEsWholeReferTypeName() (string){
    return "whole_refer_data"
}
// Whole Refer Type Mapping
func GetEsWholeReferTypeMapping() (string){
    return `{
		"whole_refer_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "first_referrer_domain":     {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}

// Whole Country Type Name
func GetEsWholeCountryTypeName() (string){
    return "whole_country_data"
}
// Whole Country Type Mapping
func GetEsWholeCountryTypeMapping() (string){
    return `{
		"whole_country_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "country_code":     {"type":"keyword"},
                "country_name":     {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}


// Whole Devide Type Name
func GetEsWholeDevideTypeName() (string){
    return "whole_devide_data"
}
// Whole Devide Type Mapping
func GetEsWholeDevideTypeMapping() (string){
    return `{
		"whole_devide_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "devide":           {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}


// Whole Sku Type Name
func GetEsWholeSkuTypeName() (string){
    return "whole_sku_data"
}
// Whole Sku Type Mapping
func GetEsWholeSkuTypeMapping() (string){
    return `{
		"whole_sku_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "sku":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}


// Whole Sku Refer Type Name
func GetEsWholeSkuReferTypeName() (string){
    return "whole_sku_refer_data"
}
// Whole Sku Refer Type Mapping
func GetEsWholeSkuReferTypeMapping() (string){
    return `{
		"whole_sku_refer_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "sku":              {"type":"keyword"},
                "first_referrer_domain": {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}



// Whole Search Type Name
func GetEsWholeSearchTypeName() (string){
    return "whole_search_data"
}
// Whole Search Type Mapping
func GetEsWholeSearchTypeMapping() (string){
    return `{
		"whole_search_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "search_text":       {"type":"keyword"},
                "search_sku_click":             {"type":"integer"},
                "search_login_email":           {"type":"integer"},
                "search_sku_cart":              {"type":"integer"},
                "search_sku_order":             {"type":"integer"},
                "search_sku_order_success":     {"type":"integer"},
                "search_qty":                   {"type":"integer"},
                "search_sku_click_rate":        {"type":"float"},
                "search_sale_rate":             {"type":"float"}
            }
        }
	}`
}



// Whole Search Lang Type Name
func GetEsWholeSearchLangTypeName() (string){
    return "whole_search_lang_data"
}
// Whole Search Lang Type Mapping
func GetEsWholeSearchLangTypeMapping() (string){
    return `{
		"whole_search_lang_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "language":         {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "search_text":       {"type":"keyword"},
                "search_sku_click":             {"type":"integer"},
                "search_login_email":           {"type":"integer"},
                "search_sku_cart":              {"type":"integer"},
                "search_sku_order":             {"type":"integer"},
                "search_sku_order_success":     {"type":"integer"},
                "search_qty":                   {"type":"integer"},
                "search_sku_click_rate":        {"type":"float"},
                "search_sale_rate":             {"type":"float"}
            }
        }
	}`
}



// Whole Url Type Name
func GetEsWholeUrlTypeName() (string){
    return "whole_url_data"
}
// Whole Url Type Mapping
func GetEsWholeUrlTypeMapping() (string){
    return `{
		"whole_url_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "url":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"}
            }
        }
	}`
}




// Whole First Url Type Name
func GetEsWholeFirstUrlTypeName() (string){
    return "whole_first_url_data"
}
// Whole First Url Type Mapping
func GetEsWholeFirstUrlTypeMapping() (string){
    return `{
		"whole_first_url_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "url":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"}
            }
        }
	}`
}



// Whole Category Type Name
func GetEsWholeCategoryTypeName() (string){
    return "whole_category_data"
}
// Whole Url Type Mapping
func GetEsWholeCategoryTypeMapping() (string){
    return `{
		"whole_category_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "category":         {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"}
            }
        }
	}`
}




// Whole App Type Name
func GetEsWholeAppTypeName() (string){
    return "whole_app_data"
}
// Whole App Type Mapping
func GetEsWholeAppTypeMapping() (string){
    return `{
		"whole_app_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "website_id":       {"type":"keyword"},
                "app":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}





// Advertise Fid Type Name
func GetEsAdvertiseFidTypeName() (string){
    return "advertise_fid_data"
}
// Advertise Fid Type Mapping
func GetEsAdvertiseFidTypeMapping() (string){
    return `{
		"advertise_fid_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "fid":              {"type":"keyword"},
                "fec_market_group": {"type":"keyword"},
                "fec_content":      {"type":"keyword"},
                "fec_source":       {"type":"keyword"},
                "fec_design":       {"type":"keyword"},
                "fec_medium_main":  {"type":"keyword"},
                
                "success_order_c_all_uv_rate": {"type":"float"},
                "success_order_c_success_uv_rate": {"type":"float"},
                
                "register_count":   {"type":"integer"},
                "login_count":      {"type":"integer"},
                "category_count":   {"type":"integer"},
                "sku_count":        {"type":"integer"},
                "search_count":     {"type":"integer"},
                
                "website_id":       {"type":"keyword"},
                "app":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}



// Advertise Content Type Name
func GetEsAdvertiseContentTypeName() (string){
    return "advertise_content_data"
}
// Advertise Content Type Mapping , 下面是去除的mapping
// "fid":              {"type":"keyword"},
// "fec_medium_main":  {"type":"keyword"},
// "fec_design":       {"type":"keyword"},
func GetEsAdvertiseContentTypeMapping() (string){
    return `{
		"advertise_content_data":{
            "properties":{
                "id":               {"type":"keyword"},
                
                "fec_market_group": {"type":"keyword"},
                "fec_content":      {"type":"keyword"},
                "fec_source":       {"type":"keyword"},
                
                "success_order_c_all_uv_rate": {"type":"float"},
                "success_order_c_success_uv_rate": {"type":"float"},
                
                "register_count":   {"type":"integer"},
                "login_count":      {"type":"integer"},
                "category_count":   {"type":"integer"},
                "sku_count":        {"type":"integer"},
                "search_count":     {"type":"integer"},
                
                "website_id":       {"type":"keyword"},
                "app":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}



// Advertise market_group Type Name
func GetEsAdvertiseMarketGroupTypeName() (string){
    return "advertise_market_group_data"
}
// Advertise market_group Type Mapping , 下面是去除的mapping
// "fid":              {"type":"keyword"},
// "fec_medium_main":  {"type":"keyword"},
// "fec_design":       {"type":"keyword"},
// "fec_content":      {"type":"keyword"},
// "fec_source":       {"type":"keyword"},
func GetEsAdvertiseMarketGroupTypeMapping() (string){
    return `{
		"advertise_market_group_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "fec_market_group": {"type":"keyword"},
                "success_order_c_all_uv_rate": {"type":"float"},
                "success_order_c_success_uv_rate": {"type":"float"},
                
                "register_count":   {"type":"integer"},
                "login_count":      {"type":"integer"},
                "category_count":   {"type":"integer"},
                "sku_count":        {"type":"integer"},
                "search_count":     {"type":"integer"},
                
                "website_id":       {"type":"keyword"},
                "app":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}


// Advertise design Type Name
func GetEsAdvertiseDesignTypeName() (string){
    return "advertise_design_data"
}
// Advertise market_group Type Mapping , 下面是去除的mapping
// "fid":              {"type":"keyword"},
// "fec_medium_main":  {"type":"keyword"},
// "fec_design":       {"type":"keyword"},
// "fec_content":      {"type":"keyword"},
// "fec_source":       {"type":"keyword"},
func GetEsAdvertiseDesignTypeMapping() (string){
    return `{
		"advertise_design_data":{
            "properties":{
                "id":               {"type":"keyword"},
                "fec_design":       {"type":"keyword"},
                "success_order_c_all_uv_rate":      {"type":"float"},
                "success_order_c_success_uv_rate":  {"type":"float"},
                
                "register_count":   {"type":"integer"},
                "login_count":      {"type":"integer"},
                "category_count":   {"type":"integer"},
                "sku_count":        {"type":"integer"},
                "search_count":     {"type":"integer"},
                
                "website_id":       {"type":"keyword"},
                "app":              {"type":"keyword"},
                "pv":               {"type":"integer"},
                "uv":               {"type":"integer"},
                "ip_count":         {"type":"integer"},
                "jump_out_count":   {"type":"integer"},
                "drop_out_count":   {"type":"integer"},
                "stay_seconds":     {"type":"integer"},
                "is_return":        {"type":"integer"},
                "first_page":       {"type":"integer"},
                "service_date_str": {"type":"date"},
                "stay_seconds_rate":{"type":"float"},
                "jump_out_rate":    {"type":"float"},
                "drop_out_rate":    {"type":"float"},
                "is_return_rate":   {"type":"float"},
                "pv_rate":          {"type":"float"},
                "sku_sale_rate":    {"type":"float"},
                "cart_count":               {"type":"integer"},
                "order_count":              {"type":"integer"},
                "success_order_count":      {"type":"integer"},
                "success_order_no_count":   {"type":"integer"},
                "order_amount":             {"type":"float"},
                "success_order_amount":     {"type":"float"}
            }
        }
	}`
}
