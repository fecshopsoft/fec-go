package advertise

import(
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/initialization"
    commonH "github.com/fecshopsoft/fec-go/handler/common"
    customerH "github.com/fecshopsoft/fec-go/handler/customer"
    "strconv"
    "errors"
)

// 得到当前账户的可用的信息
// 得到 chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, err
// 得到这些信息是为了在vue端做搜索部分使用。各个不同用户，看到的搜索内容不同。
func ActiveOwnIdAndWebsite(c *gin.Context) (int64, string, []int64, []string, error) {
    // 选中的 own_id
    var chosen_own_id int64
    // 选中的 website_id
    var chosen_website_id string
    // 可以在下拉条切换的所有 WebsiteIds
    var selectOwnIds []int64
    // 可以在下拉条切换的所有 WebsiteIds
    var selectWebsiteIds []string
    
    // 当前账户的type
    customerType := helper.GetCurrentCustomerType(c)
    // 这个是系统启动时候初始化的变量，并且cron更新该变量
    ownIdWithWebsiteId := initialization.OwnIdWithWebsiteId
    // 得到选中的 chosen_own_id
    own_id_param, _  := strconv.Atoi(c.DefaultQuery("own_id", ""))
    if own_id_param != 0 {
        chosen_own_id = int64(own_id_param)
    }
    // 得到选中的 chosen_website_id
    website_id_param := c.DefaultQuery("website_id", "")
    if website_id_param != "" {
        chosen_website_id = website_id_param
    }
    
    // 如果是超级用户
    if customerType == helper.AdminSuperType {
        for ownId, websiteIds := range ownIdWithWebsiteId {
            if len(websiteIds) == 0 {
                continue
            }
            if chosen_own_id == 0 {
                chosen_own_id = ownId
            }
            selectOwnIds = append(selectOwnIds, ownId)
            if chosen_own_id == ownId {
                selectWebsiteIds = websiteIds
                // 如果选择的chosen_website_id 为0，则将websiteIds中的第一个元素赋值给chosen website id
                if chosen_website_id == "" {
                    for _, websiteId := range websiteIds {
                        chosen_website_id = websiteId
                        break
                    }
                }
                break
            }
        }
    } else if customerType == helper.AdminCommonType {
        chosen_own_id = helper.GetCurrentCustomerId(c)
        selectOwnIds = append(selectOwnIds, chosen_own_id)
        for ownId, websiteIds := range ownIdWithWebsiteId {
            if chosen_own_id == ownId {
                // 如果存在，则设置 chosen_own_id ownIds
                selectWebsiteIds = websiteIds
                // 如果选择的chosen_website_id 为0，则将websiteIds中的第一个元素赋值给chosen website id
                isCorrectWebsiteId := 0
                if chosen_website_id != "" {
                    for _, websiteId := range websiteIds {
                        if chosen_website_id == websiteId {
                            isCorrectWebsiteId = 1
                            break
                        }
                    }
                } 
                // 说明上面传递的 chosen_website_id ，在数组中不存在，因此不合法，下面将websiteIds的第一个元素赋值给chosen_website_id
                if isCorrectWebsiteId == 0 {
                    for _, websiteId := range websiteIds {
                        chosen_website_id = websiteId
                        break
                    }
                }   
                break
            }
        }
    } else if customerType == helper.AdminChildType { // type == 3 的用户
        customer_id := helper.GetCurrentCustomerId(c)
        customer, err := customerH.GetCustomerOneById(customer_id)
        if err != nil{
            return chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, err
        }
        chosen_own_id = customer.ParentId
        selectOwnIds = append(selectOwnIds, chosen_own_id)
        for ownId, websiteIds := range ownIdWithWebsiteId {
            if chosen_own_id == ownId {
                // 如果存在，则设置 chosen_own_id ownIds
                selectWebsiteIds = websiteIds
                // 如果选择的chosen_website_id 为0，则将websiteIds中的第一个元素赋值给chosen website id
                isCorrectWebsiteId := 0
                if chosen_website_id != "" {
                    for _, websiteId := range websiteIds {
                        if chosen_website_id == websiteId {
                            isCorrectWebsiteId = 1
                            break
                        }
                    }
                } 
                // 说明上面传递的 chosen_website_id ，在数组中不存在，因此不合法，下面将websiteIds的第一个元素赋值给chosen_website_id
                if isCorrectWebsiteId == 0 {
                    for _, websiteId := range websiteIds {
                        chosen_website_id = websiteId
                        break
                    }
                }   
                break
            }
        }
    }
    // 进行判断，如果发现为空，则返回报错
    if (chosen_own_id == 0) {
        return chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, errors.New("chosen_own_id is empty")
    } else if (chosen_website_id == "") {
        return chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, errors.New("chosen_website_id is empty")
    } else if (len(selectOwnIds) == 0) {
        return chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, errors.New("selectOwnIds is empty")
    } else if (len(selectWebsiteIds) == 0) {
        return chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, errors.New("selectWebsiteIds is empty")
    }
    // 返回成功信息
    return chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, nil 
}

// 通过own_id 得到 员工和设计人员数组
func getContentAndDesign(c *gin.Context, own_id int64) ([]customerH.VueSelectOps, []customerH.VueSelectOps, error){
    var contentNames []customerH.VueSelectOps
    var designNames  []customerH.VueSelectOps
    customers, err := customerH.GetEnableCustomerChild(own_id)
    if err != nil {
        return contentNames, designNames, err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        jobType := customer.JobType
        if jobType == 1 {
            contentNames = append(contentNames, customerH.VueSelectOps{Key:customer.Id, DisplayName: customer.Name})
        } else if jobType == 2 {
            designNames = append(designNames, customerH.VueSelectOps{Key:customer.Id, DisplayName: customer.Name})
        } 
        
    }
    return contentNames, designNames, err
}


// 通过own_id 得到 员工和设计人员数组
func getMarketGroup(own_id int64) ([]customerH.VueSelectOps, error){
    var marketGroupNames []customerH.VueSelectOps
    groups, err := commonH.GetMarketGroupByOwnId(own_id)
    if err != nil {
        return marketGroupNames, err
    }
    for i:=0; i<len(groups); i++ {
        marketGroup := groups[i]
        marketGroupNames = append(marketGroupNames, customerH.VueSelectOps{Key:marketGroup.Id, DisplayName: marketGroup.Name})
    }
    return marketGroupNames, err
}

func getOwnNames(c *gin.Context, selectOwnIds []int64) ([]customerH.VueSelectOps, error) {
    var ownNames []customerH.VueSelectOps
    //customerType := helper.GetCurrentCustomerType(c)
    customers, err := customerH.GetCustomerUsernameByIds(selectOwnIds)
    if err != nil {
        return ownNames, err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        ownNames = append(ownNames, customerH.VueSelectOps{Key:customer.Id, DisplayName: customer.Username})
    }
    return ownNames, nil
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
    customerType := helper.GetCurrentCustomerType(c)
    // 这个是系统启动时候初始化的变量，并且cron更新该变量
    ownIdWithWebsiteId := initialization.OwnIdWithWebsiteId
    
    // 前端传递的website_id 是否为空，为空则返回
    website_id_param := c.DefaultQuery("website_id", "")
    if website_id_param == "" {
        return "", errors.New(" request get website_id is empty")
    }
    // 如果是超级用户, 则直接返回。
    if customerType == helper.AdminSuperType {
        return website_id_param, nil
    } else if customerType == helper.AdminCommonType {
        chosen_own_id := helper.GetCurrentCustomerId(c)
        for ownId, websiteIds := range ownIdWithWebsiteId {
            // 找到当前的ownId对应的数据
            if chosen_own_id == ownId {
                for _, websiteId := range websiteIds {
                    if website_id_param == websiteId {
                        // 从合法的websiteIds中匹配，如果找到匹配项，则说明正确。
                        return website_id_param, nil
                    }
                }
            }
        }
    } else if customerType == helper.AdminChildType { // type == 3 的用户
        customer_id := helper.GetCurrentCustomerId(c)
        customer, err := customerH.GetCustomerOneById(customer_id)
        if err != nil{
            return "", err
        }
        chosen_own_id := customer.ParentId
        for ownId, websiteIds := range ownIdWithWebsiteId {
            // 找到当前的ownId对应的数据
            if chosen_own_id == ownId {
                for _, websiteId := range websiteIds {
                    if website_id_param == websiteId {
                        // 从合法的websiteIds中匹配，如果找到匹配项，则说明正确。
                        return website_id_param, nil
                    }
                }
            }
        }
    }
    return "",errors.New("request get param website_id is not right") 
}
