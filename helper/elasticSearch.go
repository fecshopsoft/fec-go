package helper

import (
   
   
)

var numberOfShards string   = "10"
var numberOfReplicas string = "1"

// 得到当前时间对应的 Es   的 Index Name
func GetEsIndexName(websiteId string) (string){
    return "trace_" + websiteId
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
                "browser_id":       {"type":"keyword"},
                "browser_name":     {"type":"keyword"},
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
                "success_order_no_count":   {"type":"integer"}
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
    return `
        {
            "settings":{
                "number_of_shards": ` + numberOfShards + `,
                "number_of_replicas": ` + numberOfReplicas + `
            },
            "mappings":{
                "whole_all_data":{
                    "properties":{
                        "all_id":           {"type":"keyword"},
                        "pv":               {"type":"integer"},
                        "uv":               {"type":"integer"},
                        "ip_count":         {"type":"integer"},
                        
                        "jump_out_count":   {"type":"integer"},
                        "drop_out_count":   {"type":"integer"},
                        "stay_seconds":     {"type":"integer"},
                        
                        "devide":           {"type":"keyword"},
                        "country_code":     {"type":"keyword"},
                        "browser_name":     {"type":"keyword"},
                        "operate":          {"type":"keyword"},
                        
                        "is_return":        {"type":"integer"},
                        "first_page":       {"type":"integer"},
                        
                        "service_date_str": {"type":"date"},
                        "stay_seconds_rate":{"type":"float"},
                        "jump_out_rate":    {"type":"float"},
                        "drop_out_rate":    {"type":"float"},
                        "is_return_rate":   {"type":"float"},
                        "pv_rate":          {"type":"float"},
                        
                        "uv_is_return":         {"type":"integer"},
                        "uv_is_return_rate":    {"type":"float"},
                        "resolution":           {"type":"keyword"},
                        "color_depth":          {"type":"keyword"},
                        "language":             {"type":"keyword"},
                        "login_email_count":    {"type":"integer"},
                        
                        "register_email_count": {"type":"integer"},
                        "cart_count":           {"type":"integer"},
                        "order_count":          {"type":"integer"},
                        "success_order_count":  {"type":"integer"},
                        
                        
                        "order_amount":         {"type":"float"},
                        "success_order_amount": {"type":"float"},
                        
                        "category_count":       {"type":"integer"},
                        "product_count":        {"type":"integer"},
                        "search_count":         {"type":"integer"},
                        "sale_rate":            {"type":"float"},
                        
                    }
                }
            }
        }
    `
}










