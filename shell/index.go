package shell

import(
    // "github.com/fecshopsoft/fec-go/config"
    "log"
    "github.com/fecshopsoft/fec-go/helper"
    fecHander "github.com/fecshopsoft/fec-go/handler/fec"
)

func GoShell() { 
    // 初始化数据库以及索引
    InitDbIndex()
    // 初始化ElasticSearch Mapping
    // InitElasticSearchMapping()
    // 开始计算
    MapReductAndSyncDataToEsMutilDay()
}


func InitDbIndex() {
    log.Println(helper.DateTimeUTCStr() + " - Init Db Index Begin ...")
    // 创建 索引
    err := fecHander.InitTraceDataCollIndex()
    if err != nil {
        log.Println("################11")
        log.Println(err.Error())
    }
    
    log.Println(helper.DateTimeUTCStr() + " - Init Db Index Complete ...")
}

/*
func InitElasticSearchMapping() {
    log.Println(helper.DateTimeUTCStr() + " - Init ElasticSearch Mapping Begin ...")
    
    log.Println(helper.DateTimeUTCStr() + " - Init ElasticSearch Mapping Complete ...")

}
*/



