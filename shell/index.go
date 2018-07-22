package shell
/**
 * shell包的入口文件
 * 1.初始化mongodb的表以及表索引
 * 2.开始脚本处理
 */
import(
    // "github.com/fecshopsoft/fec-go/config"
    "log"
    "github.com/fecshopsoft/fec-go/helper"
    fecHander "github.com/fecshopsoft/fec-go/handler/fec"
)

func GoShell() { 
    // 初始化数据库以及索引
    InitDbIndex()
    // 开始脚本处理，进行mapreduce计算，结果写入elasticSearch
    MapReductAndSyncDataToEsMutilDay()
}

// 初始化mongodb表，以及表索引。
func InitDbIndex() {
    log.Println(helper.DateTimeUTCStr() + " - Init Db Index Begin ...")
    // 初始化mongodb表，以及表索引。
    err := fecHander.InitTraceDataCollIndex()
    if err != nil {
        log.Println("################11")
        log.Println(err.Error())
    }
    
    log.Println(helper.DateTimeUTCStr() + " - Init Db Index Complete ...")
}




