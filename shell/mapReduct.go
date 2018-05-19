package shell

import(
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/whole"
    "github.com/fecshopsoft/fec-go/shell/advertise"
    "log"
    "os"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
)

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
    for i:=0; i<len(websiteInfos); i++ {
        websiteInfo := websiteInfos[i]
        websiteId := websiteInfo.SiteUid
        //esIndexName := helper.GetEsIndexName(websiteId)
        collName := helper.GetTraceDataCollName(websiteId)
        
        // 处理：Whole Browser
        OutWholeBrowserCollName := helper.GetOutWholeBrowserCollName(websiteId)
        err = whole.BrowserMapReduct(dbName, collName, OutWholeBrowserCollName, websiteId)
        if err != nil {
            return err
        }
        
        // Whole All
        OutWholeAllCollName := helper.GetOutWholeAllCollName(websiteId)
        err = whole.AllMapReduct(dbName, collName, OutWholeAllCollName, websiteId)
        if err != nil {
            return err
        }
        
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
        
    }
    return err

}