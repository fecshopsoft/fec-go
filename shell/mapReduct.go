package shell

import(
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/whole"
    "github.com/fecshopsoft/fec-go/shell/advertise"
    "github.com/fecshopsoft/fec-go/shell/customer"
    "github.com/fecshopsoft/fec-go/db/esdb"
    "log"
    "github.com/globalsign/mgo"
    "os"
    "errors"
    "sync"
     "github.com/fecshopsoft/fec-go/db/mongodb"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
)

var WebsiteCount int = 1
var PvCount int = 200


// 计算数据，以及把数据同步到ES
func MapReductAndSyncDataToEs(){
    dateStr := helper.DateUTCStr()
    err := mapReduceByDate(dateStr)
    if err != nil {
        log.Println(err.Error())
    }
}

// 计算数据，以及把数据同步到ES
func MapReductAndSyncDataToEsMutilDay(){
    var day int
    if len(os.Args) > 1 {
        var err error
        day, err = helper.Int(os.Args[1])
        if err != nil {
            log.Println(err.Error())
        }
    }
    if day == 0 {
        day = 1
    }
    log.Println("###########")
    log.Println(day)
    dateStr := helper.DateUTCStr()
    timestamps := helper.GetTimestampsByDate(dateStr)
    
    for i:=0; i < day; i++ {
        preDayTimeStamps := timestamps - 86400 * int64(i)
        preDateStr := helper.GetDateTimeUtcByTimestamps(preDayTimeStamps)
        log.Println(preDateStr)
        log.Println("mapReduceByDate ... ")
        err := mapReduceByDate(preDateStr)
        if err != nil {
            log.Println(err.Error())
        }
    }
    
    
    
}



// 处理某一天的数据
func mapReduceByDate(dateStr string) error{
    var err error
    // 通过时间，得到数据库name
    dbName := helper.GetTraceDbNameByDate(dateStr)
    websiteInfos, err := commonHandler.GetAllActiveWebsiteId()
    if err != nil {
        return err
    }
    // 删除所有的elasticSearch Index
    //err = RemoveSpecialEsIndex()
    //if err != nil {
    //    return err
    //}
    // 判断网站数量
    if WebsiteCount < len(websiteInfos) {
        return errors.New("website count is Limit Exceeded ")
    }
    
    for i:=0; i<len(websiteInfos); i++ {
        websiteInfo := websiteInfos[i]
        websiteId := websiteInfo.SiteUid
        //esIndexName := helper.GetEsIndexName(websiteId)
        collName := helper.GetTraceDataCollName(websiteId)
        log.Println("###########")
        log.Println("OutWholeBrowserCollName")
        
        // 判断数据的pv数，是否超出限制？
        collCount := 0
        err = mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
            collCount, err = coll.Count()
            return err
        })
        if collCount > PvCount {
            return errors.New("pv count is Limit Exceeded ")
        }
        
        
        // 处理用户部分，合并email相同的用户。
        // 1.首先做email计数统计
        customerDbName := helper.GetCustomerDbName()
        customerCollName := helper.GetCustomerCollName(websiteId)
        OutCustomerEmailCollName := helper.GetOutCustomerEmailCollName(websiteId)
        err = customer.EmailMapReduct(customerDbName, customerCollName, OutCustomerEmailCollName, websiteId)
        if err != nil {
            return err
        }
        // 2.将email count >=2 的用户合并，然后将trace 基础数据进行更新
        err = customer.CustomerMergeByEmail(dbName, collName, customerDbName, customerCollName, OutCustomerEmailCollName, websiteId)
        if err != nil {
            return err
        }
        
        
        // 处理：Whole Browser
        OutWholeBrowserCollName := helper.GetOutWholeBrowserCollName(websiteId)
        err = whole.BrowserMapReduct(dbName, collName, OutWholeBrowserCollName, websiteId)
        if err != nil {
            return err
        }
        log.Println("###########")
        log.Println("OutWholeAllCollName")
        // Whole All
        OutWholeAllCollName := helper.GetOutWholeAllCollName(websiteId)
        err = whole.AllMapReduct(dbName, collName, OutWholeAllCollName, websiteId)
        if err != nil {
            return err
        }
        log.Println("###########")
        log.Println("OutWholeReferCollName")
        // 处理：Whole Refer
        OutWholeReferCollName := helper.GetOutWholeReferCollName(websiteId)
        err = whole.ReferMapReduct(dbName, collName, OutWholeReferCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Country
        OutWholeCountryCollName := helper.GetOutWholeCountryCollName(websiteId)
        err = whole.CountryMapReduct(dbName, collName, OutWholeCountryCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Devide
        OutWholeDevideCollName := helper.GetOutWholeDevideCollName(websiteId)
        err = whole.DevideMapReduct(dbName, collName, OutWholeDevideCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Sku
        OutWholeSkuCollName := helper.GetOutWholeSkuCollName(websiteId)
        err = whole.SkuMapReduct(dbName, collName, OutWholeSkuCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Sku Refer
        OutWholeSkuReferCollName := helper.GetOutWholeSkuReferCollName(websiteId)
        err = whole.SkuReferMapReduct(dbName, collName, OutWholeSkuReferCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Search
        OutWholeSearchCollName := helper.GetOutWholeSearchCollName(websiteId)
        err = whole.SearchMapReduct(dbName, collName, OutWholeSearchCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Search Lang
        OutWholeSearchLangCollName := helper.GetOutWholeSearchLangCollName(websiteId)
        err = whole.SearchLangMapReduct(dbName, collName, OutWholeSearchLangCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole Url
        OutWholeUrlCollName := helper.GetOutWholeUrlCollName(websiteId)
        err = whole.UrlMapReduct(dbName, collName, OutWholeUrlCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Whole First Url
        OutWholeFirstUrlCollName := helper.GetOutWholeFirstUrlCollName(websiteId)
        err = whole.FirstUrlMapReduct(dbName, collName, OutWholeFirstUrlCollName, websiteId)
        if err != nil {
            return err
        }
        // 处理：Whole Category
        OutWholeCategoryCollName := helper.GetOutWholeCategoryCollName(websiteId)
        err = whole.CategoryMapReduct(dbName, collName, OutWholeCategoryCollName, websiteId)
        if err != nil {
            return err
        }
        // 处理：Whole App
        OutWholeAppCollName := helper.GetOutWholeAppCollName(websiteId)
        err = whole.AppMapReduct(dbName, collName, OutWholeAppCollName, websiteId)
        if err != nil {
            return err
        }
        
        
        // 处理：Advertise Fid
        OutAdvertiseFidCollName := helper.GetOutAdvertiseFidCollName(websiteId)
        err = advertise.FidMapReduct(dbName, collName, OutAdvertiseFidCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Advertise Content
        OutAdvertiseContentCollName := helper.GetOutAdvertiseContentCollName(websiteId)
        err = advertise.ContentMapReduct(dbName, collName, OutAdvertiseContentCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Advertise MarketGroup
        OutAdvertiseMarketGroupCollName := helper.GetOutAdvertiseMarketGroupCollName(websiteId)
        err = advertise.MarketGroupMapReduct(dbName, collName, OutAdvertiseMarketGroupCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Advertise Design
        OutAdvertiseDesignCollName := helper.GetOutAdvertiseDesignCollName(websiteId)
        err = advertise.DesignMapReduct(dbName, collName, OutAdvertiseDesignCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Advertise  Campaign
        OutAdvertiseCampaignCollName := helper.GetOutAdvertiseCampaignCollName(websiteId)
        err = advertise.CampaignMapReduct(dbName, collName, OutAdvertiseCampaignCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Advertise  Medium
        OutAdvertiseMediumCollName := helper.GetOutAdvertiseMediumCollName(websiteId)
        err = advertise.MediumMapReduct(dbName, collName, OutAdvertiseMediumCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理：Advertise  Source
        OutAdvertiseSourceCollName := helper.GetOutAdvertiseSourceCollName(websiteId)
        err = advertise.SourceMapReduct(dbName, collName, OutAdvertiseSourceCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理 Advertise Edm （source + campaign）
        log.Println("###########")
        log.Println("Advertise Edm")
        OutAdvertiseEdmCollName := helper.GetOutAdvertiseEdmCollName(websiteId)
        err = advertise.EdmMapReduct(dbName, collName, OutAdvertiseEdmCollName, websiteId)
        if err != nil {
            return err
        }
        
        // 处理 customer Uuid
        log.Println("###########")
        log.Println("customer Uuid")
        OutCustomerUuidCollName := helper.GetOutCustomerUuidCollName(websiteId)
        err = customer.UuidMapReduct(dbName, collName, OutCustomerUuidCollName, websiteId)
        if err != nil {
            return err
        }
        
        
        // 处理：Advertise Eid
        OutAdvertiseEidCollName := helper.GetOutAdvertiseEidCollName(websiteId)
        err = advertise.EidMapReduct(dbName, collName, OutAdvertiseEidCollName, websiteId)
        if err != nil {
            return err
        }
        
    }
    return err

}

var removeEsIndexErr error
var once sync.Once

func RemoveSpecialEsIndex() error{
    // var err error
    once.Do(func() {
        log.Println("RemoveSpecialEsIndex, begin ...")
        var s []string;
        var t string
        
        t =  helper.GetEsCustomerUuidTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseEdmTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseMediumTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseSourceTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseCampaignTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseDesignTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseMarketGroupTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseContentTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseFidTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeAppTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeCategoryTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeFirstUrlTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeUrlTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeSearchTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeSkuReferTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeSkuTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeDevideTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeCountryTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeReferTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeAllTypeName()
        s = append(s, t )
        
        t =  helper.GetEsWholeBrowserTypeName()
        s = append(s, t )
        
        t =  helper.GetEsAdvertiseEidTypeName()
        s = append(s, t )
        
        for i:=0; i<len(s); i++ {
            typeName := s[i]
            removeEsIndexErr = RemoveEsIndex(typeName)
            if removeEsIndexErr != nil {
                return
            }
        }
        log.Println("RemoveSpecialEsIndex, complete ...")
    })
    return removeEsIndexErr      
}
// t =  helper.GetEsAdvertiseCampaignTypeName()
func RemoveEsIndex(typeName string) error{
    esIndexName := helper.GetEsIndexNameByType(typeName)
    log.Println("RemoveSpecialEsIndex, esIndexName:" + esIndexName)
    err := esdb.DeleteIndex(esIndexName)
    return err
}