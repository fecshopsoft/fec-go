package whole

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/model"
    "context"
    "strconv"
    "log"
    // "reflect"
    "encoding/json"
    "github.com/olivere/elastic"
    "github.com/fecshopsoft/fec-go/db/esdb"
)

/**
 * 1.对于common，需要价格websiteId的权限，给下面的用户，通过设置权限，来限制type==3的用户查看的websiteId
 * 2.在router部分加一个middleware，如果type==3，通过用户的parent_id，然后查看相应
 *   的common admin的所有权限的 active website_id，然后去权限表里面查看允许的website_id，然后取他们
 *   的交集，作为当前用户的可用websiteId
 * 2.1.如果type=2，则找到可用的websiteId
 * 2.2.如果type=1，如果request传递了common admin id，则使用该，否则随机找一个common admin作为 own_id
 * 3.对于type=1，服务端需要返回common own id给vue
 * 4.vue将 common own id保存到storage，并设置相应的关联的变量的值，
 * 5.vue的每次请求，都要从storage中取值，如果存在，则在放到请求里面，这个主要是针对type=1的超级admin用户
 * 6.对于vue的请求，如果是超级admin，那么就会使用传递的own_admin，如果没有，则随机找一个admin
 * 7.vue端需要有全部common own id 列表，可以通过select切换，切换后，需要重新 将 common own id保存到storage
 * 8.如果websiteId没有选择，则将可用的websiteId里面找一个可用的，然后将可用的
 *   websiteIds 和 选择的websiteId 发送给vue。vue进行刷新
 */

func SearchList(c *gin.Context){
    
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    sort     := c.DefaultQuery("sort", "")
    sort_dir := c.DefaultQuery("sort_dir", "")
    
    service_date_str_begin := c.DefaultQuery("service_date_str_begin", "")
    service_date_str_end := c.DefaultQuery("service_date_str_end", "")
    search_text := c.DefaultQuery("search_text", "")
    uv_begin := c.DefaultQuery("uv_begin", "")
    uv_end := c.DefaultQuery("uv_end", "")
    // 搜索条件
    q := elastic.NewBoolQuery()
    // service_date_str 范围搜索
    if service_date_str_begin != "" || service_date_str_end != "" {
        newRangeQuery := elastic.NewRangeQuery("service_date_str")
        if service_date_str_begin != "" {
            newRangeQuery.Gte(service_date_str_begin)
        }
        if service_date_str_end != "" {
            newRangeQuery.Lt(service_date_str_end)
        }
        q = q.Must(newRangeQuery)
    }
    // search_text 搜索
    if search_text != "" {
        q = q.Must(elastic.NewTermQuery("search_text", search_text))
    }
    // uv 范围搜索
    if uv_begin != "" || uv_end != "" {
        newRangeQuery := elastic.NewRangeQuery("uv")
        if uv_begin != "" {
            uvBeginInt, _  := strconv.Atoi(uv_begin)
            newRangeQuery.Gte(uvBeginInt)
        }
        if uv_end != "" {
            uvEndInt, _  := strconv.Atoi(uv_end)
            newRangeQuery.Lt(uvEndInt)
        }
        q = q.Must(newRangeQuery)
    }
    ////// chosen_own_id,  selectOwnIds, 
    chosen_website_id, selectWebsiteIds, err := ActiveWebsite(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    if chosen_website_id == "" {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("website id is empty"))
        return
    }
    // 添加website_id 搜索条件
    q = q.Must(elastic.NewTermQuery("website_id", chosen_website_id))
    
    // esIndexName := helper.GetEsIndexName(chosen_website_id)
    esWholeSearchTypeName :=  helper.GetEsWholeSearchTypeName()
    esIndexName := helper.GetEsIndexNameByType(esWholeSearchTypeName)
    client, err := esdb.Client()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // q = q.Must(elastic.NewTermQuery("sku", "Safari"))
    // q = q.Must(elastic.NewRangeQuery("pv").From(3).To(60))
    // q = q.Must(elastic.NewRangeQuery("service_date_str").Gt("2018-04-20").Lt("2018-04-21"))
    //termQuery := elastic.NewTermQuery("sku", "Safari")
    // termQuery := elastic.NewRangeQuery("sku", "Safari")
    //rangeQuery := NewRangeQuery("pv").Gt(3)
    log.Println(8888888888888)
    log.Println(esIndexName)
    log.Println(esWholeSearchTypeName)
    log.Println(page-1)
    log.Println(limit)
    log.Println(sort)
    search := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esWholeSearchTypeName).
        Query(q).
        From((page-1)*limit).Size(limit).
        Pretty(true)
    if sort != "" {
        if sort_dir == "ascending" {
            search.Sort(sort, true)
        } else {
            search.Sort(sort, false)
        }
    }
    searchResult, err := search.Do(context.Background())   
    /*
    searchResult, err := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esWholeSkuTypeName).
        Query(q).        // specify the query
        //Sort("user", true).      // sort by "user" field, ascending
        From(0).Size(10).        // take documents 0-9
        Pretty(true).            // pretty print request and response JSON
        Do(context.Background()) // execute
    */
    if err != nil{
        log.Println(err.Error())
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    var ts []model.WholeSearchValue
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var wholeSearch model.WholeSearchValue
            err := json.Unmarshal(*hit.Source, &wholeSearch)
            if err != nil{
                c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
                return
            }
            ts = append(ts, wholeSearch)
        }
    }
    
    siteNameOptions, siteImgUrls, err := getSiteNameAndImgUrls(c, selectWebsiteIds)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "total": searchResult.Hits.TotalHits,
        "items": ts,
        "page": page,
        "limit": limit,
        "chosen_website_id": chosen_website_id,
        "selectWebsiteIds": selectWebsiteIds,
        "siteIdOptions": siteNameOptions,
        "siteImgUrls": siteImgUrls,
        
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

// 得到 trend  info
func SearchTrendInfo(c *gin.Context){
    search_text := c.DefaultQuery("search_text", "")
    if search_text == ""{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("request get param search_text can not empty"))
        return
    }
    service_date_str := c.DefaultQuery("service_date_str", "")
    if service_date_str == ""{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("request get param service_date_str can not empty"))
        return
    }
    endtimestamps := helper.GetTimestampsByDate(service_date_str)
    //timestamps := helper.DateTimestamps()
    begintimestamps := endtimestamps - 86400 * helper.TrendDays
    preMonthDateStr := helper.GetDateTimeUtcByTimestamps(begintimestamps)
    
    // 搜索条件
    q := elastic.NewBoolQuery()
    // 加入时间
    newRangeQuery := elastic.NewRangeQuery("service_date_str")
    newRangeQuery.Gte(preMonthDateStr)
    newRangeQuery.Lte(service_date_str)
    log.Println(preMonthDateStr)
    log.Println(service_date_str)
    q = q.Must(newRangeQuery)
    // 加入 search_text
    q = q.Must(elastic.NewTermQuery("search_text", search_text))
    website_id, err := GetReqWebsiteId(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    q = q.Must(elastic.NewTermQuery("website_id", website_id))
    // esIndexName := helper.GetEsIndexName(website_id)
    esWholeSearchTypeName :=  helper.GetEsWholeSearchTypeName()
    esIndexName := helper.GetEsIndexNameByType(esWholeSearchTypeName)
    client, err := esdb.Client()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    search := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esWholeSearchTypeName).
        Query(q).
        From(0).Size(9999).
        Pretty(true)
    searchResult, err := search.Do(context.Background())   
    
    if err != nil{
        log.Println(err.Error())
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    pvTrend := make(map[string]int64)
    uvTrend := make(map[string]int64)
    ipCountTrend := make(map[string]int64)
    staySecondsTrend := make(map[string]float64)
    staySecondsRateTrend := make(map[string]float64)
    PvRateTrend := make(map[string]float64)
    JumpOutCountTrend := make(map[string]int64)
    DropOutCountTrend := make(map[string]int64)
    JumpOutRateTrend := make(map[string]float64)
    DropOutRateTrend := make(map[string]float64)
    IsReturnTrend := make(map[string]int64)
    IsReturnRateTrend := make(map[string]float64)
    
    SearchSkuClickTrend := make(map[string]int64)
    SearchLoginEmailTrend := make(map[string]int64)
    SearchSkuCartTrend := make(map[string]int64)
    SearchSkuOrderTrend := make(map[string]int64)
    SearchSkuOrderSuccessTrend := make(map[string]int64)
    SearchQtyTrend := make(map[string]int64)
    // search_sku_click
    // search_login_email
    // search_sku_cart
    // search_sku_order
    // search_sku_order_success
    // search_qty
    
    SearchSkuClickRateTrend := make(map[string]float64)
    SearchSaleRateTrend := make(map[string]float64)
    // search_sku_click_rate
    // search_sale_rate
    
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var wholeSearch model.WholeSearchValue
            err := json.Unmarshal(*hit.Source, &wholeSearch)
            if err != nil{
                c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
                return
            }
            serviceDateStr := wholeSearch.ServiceDateStr
            // pvTrend
            pvTrend[serviceDateStr] = wholeSearch.Pv
            // uvTrend
            uvTrend[serviceDateStr] = wholeSearch.Uv
            ipCountTrend[serviceDateStr] = wholeSearch.IpCount
            // staySecondsTrend
            staySecondsTrend[serviceDateStr] = wholeSearch.StaySeconds
            // staySecondsRateTrend
            staySecondsRateTrend[serviceDateStr] = wholeSearch.StaySecondsRate
            PvRateTrend[serviceDateStr] = wholeSearch.PvRate
            JumpOutCountTrend[serviceDateStr] = wholeSearch.JumpOutCount
            DropOutCountTrend[serviceDateStr] = wholeSearch.DropOutCount
            JumpOutRateTrend[serviceDateStr] = wholeSearch.JumpOutRate
            DropOutRateTrend[serviceDateStr] = wholeSearch.DropOutRate
            
            IsReturnTrend[serviceDateStr] = wholeSearch.IsReturn
            IsReturnRateTrend[serviceDateStr] = wholeSearch.IsReturnRate
            
            SearchSkuClickTrend[serviceDateStr] = wholeSearch.SearchSkuClick
            SearchLoginEmailTrend[serviceDateStr] = wholeSearch.SearchLoginEmail
            SearchSkuCartTrend[serviceDateStr] = wholeSearch.SearchSkuCart
            SearchSkuOrderTrend[serviceDateStr] = wholeSearch.SearchSkuOrder
            SearchSkuOrderSuccessTrend[serviceDateStr] = wholeSearch.SearchSkuOrderSuccess
            SearchQtyTrend[serviceDateStr] = wholeSearch.SearchQty
            
            SearchSkuClickRateTrend[serviceDateStr] = wholeSearch.SearchSkuClickRate
            SearchSaleRateTrend[serviceDateStr] = wholeSearch.SearchSaleRate
        }
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "trend": gin.H{
            "pv": pvTrend,
            "uv": uvTrend,
            "ip_count": ipCountTrend,
            "stay_seconds": staySecondsTrend,
            "stay_seconds_rate": staySecondsRateTrend,
            "pv_rate": PvRateTrend,
            "jump_out_count": JumpOutCountTrend,
            "drop_out_count": DropOutCountTrend,
            "jump_out_rate": JumpOutRateTrend,
            "drop_out_rate": DropOutRateTrend,
            
            "is_return": IsReturnTrend,
            "is_return_rate": IsReturnRateTrend,
            
            "search_sku_click": SearchSkuClickTrend,
            "search_login_email": SearchLoginEmailTrend,
            "search_sku_cart": SearchSkuCartTrend,
            "search_sku_order": SearchSkuOrderTrend,
            "search_sku_order_success": SearchSkuOrderSuccessTrend,
            "search_qty": SearchQtyTrend,
            
            "search_sku_click_rate": SearchSkuClickRateTrend,
            "search_sale_rate": SearchSaleRateTrend,
        },
    })
    // 返回json
    c.JSON(http.StatusOK, result)
    
}













