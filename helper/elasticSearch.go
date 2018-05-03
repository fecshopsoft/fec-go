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



// Whole Refer Type Name
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




