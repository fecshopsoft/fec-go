package initialization

import (
    customerHandler "github.com/fecshopsoft/fec-go/handler/customer"
    commonHandler "github.com/fecshopsoft/fec-go/handler/common"
    // fecHandler "github.com/fecshopsoft/fec-go/handler/fec"
)



type DaySiteCountT map[string]map[string]int64  // string中存储的是websiteUid
var DaySiteCount DaySiteCountT

type SiteInfo struct{
    PaymentEndTime  int64  `form:"payment_end_time" json:"payment_end_time"` //payment_end_time
    WebsiteDayMaxCount int64  `form:"website_day_max_count" json:"website_day_max_count"` //website_day_max_count
}

type SiteInfos map[string]*SiteInfo  // string中存储的是websiteUid
var WebsiteInfos SiteInfos = make(SiteInfos)

// 初始化 WebsiteInfos
func InitWebsiteInfo() error{
    DaySiteCount  = make(DaySiteCountT)
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
        websites, err := commonHandler.GetWebsiteByOwnId(own_id)
        if err != nil {
            return err
        }
        for j:=0; j<len(websites); j++ {
            website := websites[j]
            siteUid := website.SiteUid
            //WebsiteInfos[siteUid] = &siteInfo
            WebsiteInfos[siteUid] = &SiteInfo{customer.PaymentEndTime, customer.WebsiteDayMaxCount}
        }
    }
    return nil
}











