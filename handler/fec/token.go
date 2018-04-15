package fec

import(
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/helper"
    // "github.com/fecshopsoft/fec-go/handler/common"
    "net/http"
    "time"
    "log"
    "github.com/fecshopsoft/fec-go/initialization"
)



func PermissionAccessToken(c *gin.Context){
    
    access_token := helper.GetHeader(c, "Access-Token");
    if  access_token == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("Access-Token can not empty"))
        return
	}
    
    website_id, err := helper.GetSiteUIdByAccessToken(access_token)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    if  website_id == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("can not get website_id by Access-Token"))
        return
	}
    // 数据库查询，该token是否有效
    /* 数据库不需要查询了，上面的就是查询的。
    _, err = common.GetWebsiteByAccessToken(access_token)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    */
    isG := verifyWebsiteId(c, website_id)
    if isG == false {
        return
    }
    // 设置 currentWebsiteId
    c.Set("currentWebsiteId", website_id)
    
}

/*
type SiteCount struct{
    WebsiteUid  string   //payment_end_time
    Count  int64  //website_day_max_count
}
*/




func PermisstionWebsiteId(c *gin.Context){
    // defaultPageNum:= c.GetString("defaultPageNum")
    websiteId     := c.DefaultQuery("website_id", "")
    _ = verifyWebsiteId(c, websiteId)
    return
}

func verifyWebsiteId(c *gin.Context, websiteId string) (bool){
    if websiteId == "" {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("website_id is empty"))
        return false
    }
    
    // 查看这个websiteId是否存在
    siteInfo, ok := initialization.WebsiteInfos[websiteId]
    // 如果不存在，则退出
    if ok != true {  
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("website_id is not right"))
        return false
    } 
    
    log.Println(siteInfo)
    
    paymentEndTime := initialization.WebsiteInfos[websiteId].PaymentEndTime
    websiteDayMaxCount := initialization.WebsiteInfos[websiteId].WebsiteDayMaxCount
    // 如果当前时间 > paymentEndTime , 则说明过期
    nowTime := time.Now().Unix()
    if nowTime > paymentEndTime {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("Your payment time has expired, please contact the administrator to recharge"))
        return false
    }
    nowTimeStr := time.Unix(nowTime, 0).Format("2006-01-02 03:04:05")
    log.Println(nowTimeStr)
    nowTimeStr = nowTimeStr[0:10]
    log.Println(nowTimeStr)
    if _, ok := initialization.DaySiteCount[nowTimeStr]; ok == false {
        siteCount := make(map[string]int64)
        siteCount[websiteId] = 0
        initialization.DaySiteCount[nowTimeStr] = siteCount
    }
    
    if _, ok := initialization.DaySiteCount[nowTimeStr][websiteId]; ok == false {
        initialization.DaySiteCount[nowTimeStr][websiteId] = 0
    }
    
    
    if websiteDayMaxCount < initialization.DaySiteCount[nowTimeStr][websiteId] {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("max limit "))
        return false
    }
    nowCount := initialization.DaySiteCount[nowTimeStr][websiteId];
    log.Println("nowCount")
    log.Println(nowCount)
    initialization.DaySiteCount[nowTimeStr][websiteId] = initialization.DaySiteCount[nowTimeStr][websiteId] + 1
    log.Println(initialization.DaySiteCount)
    return true
}













