package fec

import(
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
)



/*
    // 附加字段 service_timestamp
    // 服务器接收数据的时间戳
    ServiceTimestamp int64 `form:"service_timestamp" json:"service_timestamp" bson:"service_timestamp"`
    // 服务器接收数据, 格式：Y-m-d H:i:s
    ServiceDatetime string `form:"service_datetime" json:"service_datetime" bson:"service_datetime"`
    // 服务器接收数据, 格式：Y-m-d
    ServiceDate string `form:"service_date" json:"service_date" bson:"service_date"`
    // 页面停留时间
    StaySeconds float64 `form:"stay_seconds" json:"stay_seconds" bson:"stay_seconds"`
    // 由于按照时间分库，站点分表，查询当前表，是否存在uuid，如果不存在，则 uuid_first_page = 1，否则 uuid_first_page = 0
    UuidFirstPage int `form:"uuid_first_page" json:"uuid_first_page" bson:"uuid_first_page"`
    // Ip First Page ，类似上面的 uuid_first_page
    IpFirstPage int `form:"ip_first_page" json:"ip_first_page" bson:"ip_first_page"`
    // uuid 
    UuidFirstCategory int `form:"uuid_first_category" json:"uuid_first_category" bson:"uuid_first_category"`
    //
    IpFirstCategory int `form:"ip_first_category" json:"ip_first_category" bson:"ip_first_category"`
    // 去掉某些参数后的url
    UrlNew string `form:"url_new" json:"url_new" bson:"url_new"`
    // 登录后访问搜索页面的用户
    SearchLoginEmail int `form:"search_login_email" json:"search_login_email" bson:"search_login_email"`
*/

// 得到停留时间
func updatePreStaySeconds(dbName string, collName string, uuid string, serviceTimestamp int64) (error){
    var staySeconds float64 = 0
    
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"uuid": uuid}).Sort("-service_timestamp").One(&traceInfo)
        
        // 如果查询不到，则
        if traceInfo.ServiceTimestamp == 0 {
            return nil
        }
        // 得到停留时间    
        staySeconds = float64(serviceTimestamp - traceInfo.ServiceTimestamp)
        if staySeconds <= 0 {
            staySeconds = 0.1
        // 当时间大于600秒，则取600秒，以免造成停留时间过长。
        } else if staySeconds > 600 {
            staySeconds = 600
        }
        // 更新 上一次访问的停留时间。
        selector := bson.M{"_id": traceInfo.Id_}
        updateData := bson.M{"$set": bson.M{"stay_seconds": staySeconds}}
        _ = coll.Update(selector, updateData)
        //return err
        return nil
    })
    return err
} 



// 得到停留时间
func updatePreStaySecondsAndReturn(dbName string, collName string, uuid string, serviceTimestamp int64) (TraceInfo, error){
    var staySeconds float64 = 0
    var traceInfo TraceInfo
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        _ = coll.Find(bson.M{"uuid": uuid}).Sort("-service_timestamp").One(&traceInfo)
        
        // 如果查询不到，则
        if traceInfo.ServiceTimestamp == 0 {
            return nil
        }
        // 得到停留时间    
        staySeconds = float64(serviceTimestamp - traceInfo.ServiceTimestamp)
        if staySeconds <= 0 {
            staySeconds = 0.1
        // 当时间大于600秒，则取600秒，以免造成停留时间过长。
        } else if staySeconds > 600 {
            staySeconds = 600
        }
        // 更新 上一次访问的停留时间。
        selector := bson.M{"_id": traceInfo.Id_}
        updateData := bson.M{"$set": bson.M{"stay_seconds": staySeconds}}
        _ = coll.Update(selector, updateData)
        //return err
        return nil
    })
    return traceInfo, err
} 


func getUuidFirstPage(dbName string, collName string, uuid string) (int, error) {
    var uuidFirstPage int = 0
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"uuid": uuid}).One(&traceInfo)
        
        // 如果查询不到，则说明该ip为首次访问
        if traceInfo.Uuid == "" {
            uuidFirstPage = 1
        } else {
            return nil
        }
        return nil
    })
    return uuidFirstPage, err
} 

func getIpFirstPage(dbName string, collName string, ipStr string) (int, error) {
    var ipFirstPage int = 0
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"ip": ipStr}).One(&traceInfo)
        
        // 如果查询不到，则说明该ip为首次访问
        if traceInfo.Ip == "" {
            ipFirstPage = 1
        } else {
            return nil
        }
        return nil
    })
    return ipFirstPage, err
}   

func getUuidFirstCategory(dbName string, collName string, uuid string, categoryName string) (int, error) {
    var uuidFirstCategory int = 0
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"uuid": uuid, "category": categoryName}).One(&traceInfo)
        
        // 如果查询不到，则说明该ip为首次访问
        if traceInfo.Uuid == "" {
            uuidFirstCategory = 1
        } else {
            return nil
        }
        return nil
    })
    return uuidFirstCategory, err
} 
   
func getIpFirstCategory(dbName string, collName string, ipStr string, categoryName string) (int, error) {
    var ipFirstCategory int = 0
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"ip": ipStr, "category": categoryName}).One(&traceInfo)
        
        // 如果查询不到，则说明该ip为首次访问
        if traceInfo.Ip == "" {
            ipFirstCategory = 1
        } else {
            return nil
        }
        return nil
    })
    return ipFirstCategory, err
}  

func getUrlNew(originUrl string) string{
    return originUrl
}
    

func getSearchLoginEmail(dbName string, collName string, uuid string) (int, error) {
    var searchLoginEmail int = 0
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"uuid": uuid, "login_email": bson.M{"$exists":true}}).One(&traceInfo)
        
        // 如果查询,则说明是登录用户进行的搜索
        if traceInfo.Uuid != "" {
            searchLoginEmail = 1
        } else {
            return nil
        }
        return nil
    })
    return searchLoginEmail, err
}      
  
func getFirstVisitThisUrl(dbName string, collName string, uuid string, urlNew string) (int, error) {
    var firstVisitThisUrl int = 0
    err := mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        var traceInfo TraceMiddInfo
        
        _ = coll.Find(bson.M{"uuid": uuid, "url_new": urlNew}).One(&traceInfo)
        
        // 如果查询不到，则说明该URL为首次访问
        if traceInfo.Uuid == "" {
            firstVisitThisUrl = 1
        } else {
            return nil
        }
        return nil
    })
    return firstVisitThisUrl, err
} 
  