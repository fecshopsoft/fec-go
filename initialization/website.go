package initialization

import (
    customerHandler "github.com/fecshopsoft/fec-go/handler/customer"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
    // fecHandler "github.com/fecshopsoft/fec-go/handler/fec"
    "log"
)



type DaySiteCountT map[string]map[string]int64  // string中存储的是websiteUid
var DaySiteCount DaySiteCountT

type SiteInfo struct{
    PaymentEndTime  int64  `form:"payment_end_time" json:"payment_end_time"` //payment_end_time
    WebsiteDayMaxCount int64  `form:"website_day_max_count" json:"website_day_max_count"` //website_day_max_count
}

type SiteInfos map[string]*SiteInfo  // string中存储的是websiteUid
var WebsiteInfos SiteInfos = make(SiteInfos)

// 统计后的记过，需要ownId和websiteId的对应关系
// 将初始化关系保存带 OwnIdWithWebsiteId 中
var OwnIdWithWebsiteId map[int64][]string = make(map[int64][]string)
var CustomerIdWithMarketGroup map[int64]int64 = make(map[int64]int64)
var CustomerIdWithUsername map[int64]string = make(map[int64]string)
var CustomerIdWithName map[int64]string = make(map[int64]string)
var MarketGroupIdWithName map[int64]string = make(map[int64]string)

func init(){
    DaySiteCount  = make(DaySiteCountT)
}
// 初始化 WebsiteInfos
func InitWebsiteInfo() error{
    var webInfos SiteInfos = make(SiteInfos)
	customers, err:= customerHandler.GetPaymentActiveCustomers()
    if err != nil {
        return err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        // var siteInfo SiteInfo
        // siteInfo.PaymentEndTime = customer.PaymentEndTime
        // siteInfo.WebsiteDayMaxCount = customer.WebsiteDayMaxCount
        own_id := customer.Id
        // 查询出来该customer_id 对应的所有的website
        websites, err := commonHandler.GetActiveWebsiteByOwnId(own_id)
        if err != nil {
            return err
        }
        var websiteIds []string
        for j:=0; j<len(websites); j++ {
            website := websites[j]
            siteUid := website.SiteUid
            //WebsiteInfos[siteUid] = &siteInfo
            webInfos[siteUid] = &SiteInfo{website.PaymentEndTime, website.WebsiteDayMaxCount}
            websiteIds = append(websiteIds, siteUid)
        }
        OwnIdWithWebsiteId[own_id] = websiteIds
        CustomerIdWithMarketGroup[customer.Id] = customer.MarketGroupId
        CustomerIdWithName[customer.Id] = customer.Name
        CustomerIdWithUsername[customer.Id] = customer.Username
        log.Println("############")
        log.Println(customer.Id)
        log.Println(customer.MarketGroupId)
    }
    WebsiteInfos = webInfos
    log.Println(OwnIdWithWebsiteId)
    // 计算： MarketGroupIdWithName
    marketGroups, err := commonHandler.GetAllMarketGroup()
    if err != nil {
        return err
    }
    for k:=0; k<len(marketGroups); k++ {
        marketGroup := marketGroups[k]
        MarketGroupIdWithName[marketGroup.Id] = marketGroup.Name
    }
    
    return nil
}








