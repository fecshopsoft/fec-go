package shell

import(
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/whole"
    "log"
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
        esIndexName := helper.GetEsIndexName(websiteId)
        collName := helper.GetTraceDataCollName(websiteId)
        
        // Whole Browser
        OutWholeBrowserCollName := helper.GetOutWholeBrowserCollName(websiteId)
        err = whole.BrowserMapReduct(dbName, collName, OutWholeBrowserCollName, esIndexName)
        if err != nil {
            return err
        }
        
        /*
        // Whole All
        OutWholeAllCollName := helper.GetOutWholeAllCollName(websiteId)
        err = whole.AllMapReduct(dbName, collName, OutWholeAllCollName, esIndexName)
        if err != nil {
            return err
        }
        */
        
    }
    return err

}