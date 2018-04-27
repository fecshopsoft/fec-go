package shell

import(
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/whole"
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
    day, err := helper.Int(os.Args[1])
    if err != nil {
        log.Println(err.Error())
        return 
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
        
    }
    return err

}