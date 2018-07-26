package advertise

import(
    "github.com/gin-gonic/gin"
    //"github.com/fecshopsoft/fec-go/helper"
    //"github.com/fecshopsoft/fec-go/initialization"
    commonH "github.com/fecshopsoft/fec-go/handler/common"
    customerH "github.com/fecshopsoft/fec-go/handler/customer"
    //"strconv"
    "errors"
)

// 得到当前账户的可用的信息
// 得到 chosen_website_id, selectWebsiteIds, err
// 得到这些信息是为了在vue端做搜索部分使用。各个不同用户，看到的搜索内容不同。
func ActiveWebsite(c *gin.Context) (string,  []string, error) {
    // 选中的 own_id
    //var chosen_own_id int64
    // 选中的 website_id
    var chosen_website_id string
    // 可以在下拉条切换的所有 WebsiteIds
   // var selectOwnIds []int64
    // 可以在下拉条切换的所有 WebsiteIds
    var selectWebsiteIds []string
    
    // 当前账户的type
    //customerType := helper.GetCurrentCustomerType(c)
	
	
    // 这个是系统启动时候初始化的变量，并且cron更新该变量
    //ownIdWithWebsiteId := initialization.OwnIdWithWebsiteId
    // 得到选中的 chosen_own_id
    //own_id_param, _  := strconv.Atoi(c.DefaultQuery("own_id", ""))
   // if own_id_param != 0 {
    //    chosen_own_id = int64(own_id_param)
   // }
    // 得到选中的 chosen_website_id
    website_id_param := c.DefaultQuery("website_id", "")
    if website_id_param != "" {
        chosen_website_id = website_id_param
    }
	
	// 得到websites  以及选中
	websites, _ := commonH.GetAllActiveWebsites()
	for j:=0; j<len(websites); j++ {
		website := websites[j]
		if chosen_website_id == "" {
			chosen_website_id = website.SiteUid
		}
		//WebsiteInfos[siteUid] = &siteInfo
		//webInfos[siteUid] = &SiteInfo{website.PaymentEndTime, website.WebsiteDayMaxCount}
		selectWebsiteIds = append(selectWebsiteIds,  website.SiteUid)
	}
    
    // 返回成功信息
    return chosen_website_id, selectWebsiteIds, nil 
}

// 通过own_id 得到 员工和设计人员数组
func getContentAndDesign(c *gin.Context) ([]customerH.VueSelectOps, []customerH.VueSelectOps, error){
    var contentNames []customerH.VueSelectOps
    var designNames  []customerH.VueSelectOps
    customers, err := customerH.GetActiveCustomers()
	// 得到所有的
    if err != nil {
        return contentNames, designNames, err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        jobType := customer.JobType
        if jobType == 1 {
            contentNames = append(contentNames, customerH.VueSelectOps{Key:customer.Id, DisplayName: customer.Username})
        } else if jobType == 2 {
            designNames = append(designNames, customerH.VueSelectOps{Key:customer.Id, DisplayName: customer.Username})
        } 
        
    }
    return contentNames, designNames, err
}


// 得到 marketGroups
func getMarketGroup() ([]customerH.VueSelectOps, error){
    var marketGroupNames []customerH.VueSelectOps
    groups, err := commonH.GetAllMarketGroup()
    if err != nil {
        return marketGroupNames, err
    }
    for i:=0; i<len(groups); i++ {
        marketGroup := groups[i]
        marketGroupNames = append(marketGroupNames, customerH.VueSelectOps{Key:marketGroup.Id, DisplayName: marketGroup.Name})
    }
    return marketGroupNames, err
}


func getSiteNames(c *gin.Context, selectWebsiteIds []string) ([]commonH.VueSelectOps, error) {
    var siteNames []commonH.VueSelectOps
    // customerType := helper.GetCurrentCustomerType(c)
    sites, err := commonH.GetWebsiteBySiteUids(selectWebsiteIds)
    if err != nil {
        return siteNames, err
    }
    for i:=0; i<len(sites); i++ {
        site := sites[i]
        siteNames = append(siteNames, commonH.VueSelectOps{Key:site.SiteUid, DisplayName: site.SiteName})
    }
    return siteNames, nil
}

func getSiteNameAndImgUrls(c *gin.Context, selectWebsiteIds []string) ([]commonH.VueSelectOps, map[string]string, error) {
    var siteNames []commonH.VueSelectOps
    siteImgUrls := make(map[string]string)
    // customerType := helper.GetCurrentCustomerType(c)
    sites, err := commonH.GetWebsiteBySiteUids(selectWebsiteIds)
    if err != nil {
        return siteNames, siteImgUrls, err
    }
    for i:=0; i<len(sites); i++ {
        site := sites[i]
        siteNames = append(siteNames, commonH.VueSelectOps{Key:site.SiteUid, DisplayName: site.SiteName})
        siteImgUrls[site.SiteUid] = site.SkuImageApiUrl
        // siteImgUrls = append(siteImgUrls, commonH.VueSelectOps{Key:site.SiteUid, DisplayName: site.SkuImageApiUrl})
    }
    return siteNames, siteImgUrls, nil
}

// 前台传递的websiteId是否是合法的。
func GetReqWebsiteId(c *gin.Context) (string, error) {
    
    // 选中的 website_id
    // var chosen_website_id string
    // 可以在下拉条切换的所有 WebsiteIds
    // var selectOwnIds []int64
    // 可以在下拉条切换的所有 WebsiteIds
    // var selectWebsiteIds []string
    
    // 当前账户的type
    //customerType := helper.GetCurrentCustomerType(c)

    // 前端传递的website_id 是否为空，为空则返回
    website_id_param := c.DefaultQuery("website_id", "")
    if website_id_param == "" {
        return "", errors.New(" request get website_id is empty")
    }
	return website_id_param, nil
}
