package shell
/**
 * shell脚本处理文件
 * 1.从脚本传递的第一个参数获取处理N天的数据
 * 2.从脚本的第二个参数，获取是否要删除原来的es的所有index，是否重新跑es的数据
 *    譬如  go run fec-go-shell.go 1 removeEsAllIndex 将会删除掉 es中的所有index
 *    删除掉后，因此每个部分在执行脚本统计前都会进行initMapping操作，因此，不会存在其他的问题
 *    删除后，您可以将历史数据重新跑一次进行恢复，譬如跑最近一个月的数据： go run fec-go-shell.go 30
 * 3.根据第一步骤循环遍历，开始按照天数循环处理数据
 * 4.将当前日期下的所有网站，遍历，进行数据统计。
 * 5.依次遍历，处理各个部分的数据统计，将mongodb中的数据进行mapreduce处理，将结果数据复制到elasticSearch中。
 
 */
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


// 网站数量默认值
var WebsiteCount int = 1
// 网站pv数最大值
var PvCount int = 5000

// 废弃函数
// 计算数据，以及把数据同步到ES
// 该函数只能处理一天的数据处理，因此废弃
func MapReductAndSyncDataToEs(){
    dateStr := helper.DateUTCStr()
    err := mapReduceByDate(dateStr)
    if err != nil {
        log.Println(err.Error())
    }
}

// 计算数据，以及把数据同步到ES
// 根据脚本传递的第一个参数，决定处理当前时间前n天的数据统计。
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
    
	// 删除所有的elasticSearch Index, 
	// 对于新建mapping部分，在计算脚本的时候，都会init相应的mapping
	// 譬如：对于./shell/advertise/eid.go 444行代码出：  esAdvertiseEidTypeMapping := helper.GetEsAdvertiseEidTypeMapping()
	// 因此删除原来的所有的es index，只需要将下面的
	operateType := os.Args[2]
	if operateType == "removeEsAllIndex" {
		log.Println("begin remove es all index")
		err := RemoveSpecialEsIndex()
		if err != nil {
			log.Println(err.Error())
		}
	}
    
	
	
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
		// 不做数据限制
		/*
        collCount := 0
        err = mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
            collCount, err = coll.Count()
            return err
        })
        if collCount > PvCount {
            return errors.New("pv count is Limit Exceeded ")
        }
		*/
        
        
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
			log.Println("########: customer Uuid error")
            return err
        }
        
        
        // 处理：Advertise Eid
		log.Println("###########")
        log.Println("Advertise Eid")
        OutAdvertiseEidCollName := helper.GetOutAdvertiseEidCollName(websiteId)
        err = advertise.EidMapReduct(dbName, collName, OutAdvertiseEidCollName, websiteId)
        if err != nil {
			log.Println("########: Advertise Eid error")
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