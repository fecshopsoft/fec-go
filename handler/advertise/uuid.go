package advertise

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/fecshopsoft/fec-go/helper"
    model "github.com/fecshopsoft/fec-go/shell/customerModel"
    "github.com/fecshopsoft/fec-go/initialization"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
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

func UuidList(c *gin.Context){
    
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    sort     := c.DefaultQuery("sort", "")
    sort_dir := c.DefaultQuery("sort_dir", "")
    
    service_date_str_begin := c.DefaultQuery("service_date_str_begin", "")
    service_date_str_end := c.DefaultQuery("service_date_str_end", "")
    customer_id := c.DefaultQuery("customer_id", "")
    customer_email := c.DefaultQuery("customer_email", "")
    fec_content := c.DefaultQuery("fec_content", "")
    fec_design := c.DefaultQuery("fec_design", "")
    
    fid := c.DefaultQuery("fid", "")
    fec_source := c.DefaultQuery("fec_source", "")
    fec_campaign := c.DefaultQuery("fec_campaign", "")
    
    isReturn, _  := strconv.Atoi(c.DefaultQuery("is_return", ""))
    
    
    pv_begin := c.DefaultQuery("pv_begin", "")
    pv_end := c.DefaultQuery("pv_end", "")
     
    visit_page_sku_begin := c.DefaultQuery("visit_page_sku_begin", "")
    visit_page_sku_end := c.DefaultQuery("visit_page_sku_end", "")
    visit_page_cart_begin := c.DefaultQuery("visit_page_cart_begin", "")
    visit_page_cart_end := c.DefaultQuery("visit_page_cart_end", "")
    visit_page_order_begin := c.DefaultQuery("visit_page_order_begin", "")
    visit_page_order_end := c.DefaultQuery("visit_page_order_end", "")
    visit_page_order_processing_begin := c.DefaultQuery("visit_page_order_processing_begin", "")
    visit_page_order_processing_end := c.DefaultQuery("visit_page_order_processing_end", "")
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
    if isReturn != 0 {
        if isReturn != 1 {
            isReturn = 0
        }
        q = q.Must(elastic.NewTermQuery("is_return", isReturn))
    }
    // customer_id 搜索
    if customer_id != "" {
        q = q.Must(elastic.NewTermQuery("customer_id", customer_id))
    }
    if fec_content != "" {
        q = q.Must(elastic.NewTermQuery("fec_content_main", fec_content))
    }
    if fec_design != "" {
        q = q.Must(elastic.NewTermQuery("fec_design_main", fec_design))
    }
    
    if fid != "" {
        q = q.Must(elastic.NewTermQuery("fid_main", fid))
    }
    if fec_source != "" {
        q = q.Must(elastic.NewTermQuery("fec_source_main", fec_source))
    }
    if fec_campaign != "" {
        q = q.Must(elastic.NewTermQuery("fec_campaign_main", fec_campaign))
    }
    
    // pv 范围搜索
    if pv_begin != "" || pv_end != "" {
        newRangeQuery := elastic.NewRangeQuery("pv")
        if pv_begin != "" {
            pvBeginInt, _  := strconv.Atoi(pv_begin)
            newRangeQuery.Gte(pvBeginInt)
        }
        if pv_end != "" {
            pvEndInt, _  := strconv.Atoi(pv_end)
            newRangeQuery.Lt(pvEndInt)
        }
        q = q.Must(newRangeQuery)
    }
    
    // visit_page_sku 范围搜索
    if visit_page_sku_begin != "" || visit_page_sku_end != "" {
        newRangeQuery := elastic.NewRangeQuery("visit_page_sku")
        if visit_page_sku_begin != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_sku_begin)
            newRangeQuery.Gte(pvEndInt)
        }
        if visit_page_sku_end != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_sku_end)
            newRangeQuery.Lt(pvEndInt)
        }
        q = q.Must(newRangeQuery)
    }
    
    // visit_page_cart 范围搜索
    if visit_page_cart_begin != "" || visit_page_cart_end != "" {
        newRangeQuery := elastic.NewRangeQuery("visit_page_cart")
        if visit_page_cart_begin != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_cart_begin)
            newRangeQuery.Gte(pvEndInt)
        }
        if visit_page_cart_end != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_cart_end)
            newRangeQuery.Lt(pvEndInt)
        }
        q = q.Must(newRangeQuery)
    }
    
    // visit_page_order 范围搜索
    if visit_page_order_begin != "" || visit_page_order_end != "" {
        newRangeQuery := elastic.NewRangeQuery("visit_page_order")
        if visit_page_order_begin != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_order_begin)
            newRangeQuery.Gte(pvEndInt)
        }
        if visit_page_order_end != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_order_end)
            newRangeQuery.Lt(pvEndInt)
        }
        q = q.Must(newRangeQuery)
    }    
    
    // visit_page_order_processing 范围搜索
    if visit_page_order_processing_begin != "" || visit_page_order_processing_end != "" {
        newRangeQuery := elastic.NewRangeQuery("visit_page_order_processing")
        if visit_page_order_processing_begin != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_order_processing_begin)
            newRangeQuery.Gte(pvEndInt)
        }
        if visit_page_order_processing_end != "" {
            pvEndInt, _  := strconv.Atoi(visit_page_order_processing_end)
            newRangeQuery.Lt(pvEndInt)
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
    
    // 查询出来当前的员工和设计者
    contentNames, designNames, err := getContentAndDesign(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 添加website_id 搜索条件
    q = q.Must(elastic.NewTermQuery("website_id", chosen_website_id))
    
    if customer_email != "" {
        // 做查询
        customerDbName := helper.GetCustomerDbName()
        customerCollName := helper.GetCustomerCollName(chosen_website_id)
        var uuidCustomer model.UuidCustomer
        _ = mongodb.MDC(customerDbName, customerCollName, func(coll *mgo.Collection) error {
             _ = coll.Find(bson.M{"emails": customer_email}).One(&uuidCustomer)
            return nil
        })
        if uuidCustomer.CustomerId != "" {
            q = q.Must(elastic.NewTermQuery("customer_id", uuidCustomer.CustomerId))   
        }
    }
    
    // esIndexName := helper.GetEsIndexName(chosen_website_id)
    esCustomerUuidTypeName :=  helper.GetEsCustomerUuidTypeName()
    esIndexName := helper.GetEsIndexNameByType(esCustomerUuidTypeName)
    client, err := esdb.Client()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // q = q.Must(elastic.NewTermQuery("country_code", "Safari"))
    // q = q.Must(elastic.NewRangeQuery("pv").From(3).To(60))
    // q = q.Must(elastic.NewRangeQuery("service_date_str").Gt("2018-04-20").Lt("2018-04-21"))
    //termQuery := elastic.NewTermQuery("country_code", "Safari")
    // termQuery := elastic.NewRangeQuery("country_code", "Safari")
    //rangeQuery := NewRangeQuery("pv").Gt(3)
    log.Println(8888888888888)
    log.Println(esIndexName)
    log.Println(esCustomerUuidTypeName)
    log.Println(page-1)
    log.Println(limit)
    log.Println(sort)
    search := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esCustomerUuidTypeName).
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
        Type(esCustomerUuidTypeName).
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
    
    var designGroupArr  []helper.VueSelectOps
    var contentGroupArr []helper.VueSelectOps
    var marketGroupArr  []helper.VueSelectOps
    
    var ts []model.CustomerUuidValue2
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var customerUuid model.CustomerUuidValue2
            err := json.Unmarshal(*hit.Source, &customerUuid)
            if err != nil{
                c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
                return
            }
            ts = append(ts, customerUuid)
            
            fecDesignInt64, _ := helper.Int64(customerUuid.FecDesignMain)
            designGroupVal := initialization.CustomerIdWithUsername[fecDesignInt64]
            designGroupArr = append(designGroupArr, helper.VueSelectOps{Key: fecDesignInt64, DisplayName: designGroupVal})
            
            fecContent64, _ := helper.Int64(customerUuid.FecContentMain)
            contentGroupVal := initialization.CustomerIdWithUsername[fecContent64]
            contentGroupArr = append(contentGroupArr, helper.VueSelectOps{Key: fecContent64, DisplayName: contentGroupVal})
            
            fecMarketGroup64, _ := helper.Int64(customerUuid.FecMarketGroupMain)
            marketGroupVal := initialization.MarketGroupIdWithName[fecMarketGroup64]
            marketGroupArr = append(marketGroupArr, helper.VueSelectOps{Key: fecMarketGroup64, DisplayName: marketGroupVal})
            
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
        "chosen_website_id": chosen_website_id,
        "selectWebsiteIds": selectWebsiteIds,
        "siteIdOptions": siteNameOptions,
        "siteImgUrls": siteImgUrls,
        "designGroupOps": designGroupArr,
        "contentGroupOps": contentGroupArr,
        "marketGroupOps": marketGroupArr,
        "contentSelectOps": contentNames,
        "designSelectOps": designNames,
    })
    
    // 返回json
    c.JSON(http.StatusOK, result)
}



func UuidOne(c *gin.Context){
    id     := c.DefaultQuery("id", "")
    website_id     := c.DefaultQuery("website_id", "")
    // 添加website_id 搜索条件
    
    if id == "" || website_id == "" {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("is and website_id can not empty"))
        return
    }
    q := elastic.NewBoolQuery()
    q = q.Must(elastic.NewTermQuery("website_id", website_id))
    q = q.Must(elastic.NewTermQuery("id", id))
    
    // esIndexName := helper.GetEsIndexName(chosen_website_id)
    esCustomerUuidTypeName :=  helper.GetEsCustomerUuidTypeName()
    esIndexName := helper.GetEsIndexNameByType(esCustomerUuidTypeName)
    client, err := esdb.Client()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    search := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esCustomerUuidTypeName).
        Query(q).
        From(0).Size(1).
        Pretty(true)
    
    searchResult, err := search.Do(context.Background())   
    /*
    searchResult, err := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esCustomerUuidTypeName).
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
    
    var ts model.CustomerUuidValue
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var customerUuid model.CustomerUuidValue
            err := json.Unmarshal(*hit.Source, &customerUuid)
            if err != nil{
                c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
                return
            }
            ts = customerUuid
            
        }
    }
    ft := make(map[string]interface{})
    ft["devide"] = getMpFormat(ts.Devide)
    ft["language"] = getMpFormat(ts.Language)
    ft["fec_app"] = getMpFormat(ts.FecApp)
    ft["search"] = getMpFormat(ts.Search)
    ft["category"] = getMpFormat(ts.Category)
    ft["sku"] = getSkuMpFormat(ts.Sku, ts.SkuCart, ts.SkuOrder, ts.SkuOrderSuccess)
    
    ft["operate"] = getMpFormat(ts.Operate)
    ft["country_code"] = getMpFormat(ts.CountryCode)
    ft["browser_name"] = getMpFormat(ts.BrowserName)
    ft["resolution"] = getMpFormat(ts.Resolution)
    ft["color_depth"] = getMpFormat(ts.ColorDepth)
    tsc := ts.CustomerEmail
    customerEmails := ""
    for i:=0; i<len(tsc); i++ {
        if (i != 0) {
            customerEmails += ", " + tsc[i]
        } else {
            customerEmails += tsc[i]
        }
    }
    ft["customer_emails"] = customerEmails
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "item": ts,
        "format": ft,
    })
    
    // 返回json
    c.JSON(http.StatusOK, result)
}

func getSkuMpFormat(sku map[string]int64, sku_cart map[string]int64, sku_order map[string]int64, sku_order_success map[string]int64) []MpSkuInt {
    var mp []MpSkuInt
    exist := make(map[string]string)
    for sk, _ := range sku {
        if _, ok := exist[sk]; !ok {
            exist[sk] = sk
        }
    } 
    for sk,_ := range sku_cart {
        if _, ok := exist[sk]; !ok {
            exist[sk] = sk
        }
    } 
    for sk,_ := range sku_order {
        if _, ok := exist[sk]; !ok {
            exist[sk] = sk
        }
    }
    for sk,_ := range sku_order_success {
        if _, ok := exist[sk]; !ok {
            exist[sk] = sk
        }
    }
    // 遍历 
    for sk,_ := range exist {
        var sku_count int64 = 0
        var sku_cart_count int64 = 0
        var sku_order_count int64 = 0
        var sku_order_success_count int64 = 0
        if _, ok := sku[sk]; ok {
            sku_count = sku[sk]
        }
        if _, ok := sku_cart[sk]; ok {
            sku_cart_count = sku_cart[sk]
        }
        if _, ok := sku_order[sk]; ok {
            sku_order_count = sku_order[sk]
        }
        if _, ok := sku_order_success[sk]; ok {
            sku_order_success_count = sku_order_success[sk]
        }
        // 组装
        var p MpSkuInt
        p.Namet = sk
        p.SkuCount = sku_count
        p.CartCount = sku_cart_count
        p.OrderCount = sku_order_count
        p.OrderSuccessCount = sku_order_success_count
        mp = append(mp, p)
    }
    return mp
}

func getMpFormat(m map[string]int64)[]MpInt{
    var mp []MpInt
    for k,v := range m {
        var mpi MpInt
        mpi.Namet = k
        mpi.Count = v
        mp = append(mp, mpi)
    }
    return mp
}

type MpInt struct{
    Namet string `form:"namet" json:"namet" bson:"namnamete"`
    Count int64 `form:"count" json:"count" bson:"count"`
}

type MpSkuInt struct{
    Namet string `form:"namet" json:"namet" bson:"namnamete"`
    SkuCount int64 `form:"sku_count" json:"sku_count" bson:"sku_count"`
    CartCount int64 `form:"cart_count" json:"cart_count" bson:"cart_count"`
    OrderCount int64 `form:"order_count" json:"order_count" bson:"order_count"`
    OrderSuccessCount int64 `form:"order_success_count" json:"order_success_count" bson:"order_success_count"`
}

// 得到 trend  info
func UuidTrendInfo(c *gin.Context){
    customer_id := c.DefaultQuery("customer_id", "")
    if customer_id == ""{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("request get param fid can not empty"))
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
    // 加入浏览器
    q = q.Must(elastic.NewTermQuery("customer_id", customer_id))
    website_id, err := GetReqWebsiteId(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    q = q.Must(elastic.NewTermQuery("website_id", website_id))
    // esIndexName := helper.GetEsIndexName(website_id)
    esCustomerUuidTypeName :=  helper.GetEsCustomerUuidTypeName()
    esIndexName := helper.GetEsIndexNameByType(esCustomerUuidTypeName)
    client, err := esdb.Client()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    search := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esCustomerUuidTypeName).
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
    StaySecondsTrend := make(map[string]float64)
    VisitPageSkuTrend := make(map[string]int64)
    VisitPageCategoryTrend := make(map[string]int64)
    VisitPageSearchTrend := make(map[string]int64)
    VisitPageCartTrend := make(map[string]int64)
    VisitPageOrderPendingTrend := make(map[string]int64)
    VisitPageOrderPendingAmountTrend := make(map[string]float64)
    VisitPageOrderProcessingTrend := make(map[string]int64)
    VisitPageOrderProcessingAmountTrend := make(map[string]float64)
    
    
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var customerUuid model.CustomerUuidValue2
            err := json.Unmarshal(*hit.Source, &customerUuid)
            if err != nil{
                c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
                return
            }
            serviceDateStr := customerUuid.ServiceDateStr
            
            pvTrend[serviceDateStr] = customerUuid.Pv
            StaySecondsTrend[serviceDateStr] = customerUuid.StaySeconds
            VisitPageSkuTrend[serviceDateStr] = customerUuid.VisitPageSku
            VisitPageCategoryTrend[serviceDateStr] = customerUuid.VisitPageCategory
            VisitPageSearchTrend[serviceDateStr] = customerUuid.VisitPageSearch
            VisitPageCartTrend[serviceDateStr] = customerUuid.VisitPageCart
            VisitPageOrderPendingTrend[serviceDateStr] = customerUuid.VisitPageOrderPending
            VisitPageOrderPendingAmountTrend[serviceDateStr] = customerUuid.VisitPageOrderPendingAmount
            VisitPageOrderProcessingTrend[serviceDateStr] = customerUuid.VisitPageOrderProcessing
            VisitPageOrderProcessingAmountTrend[serviceDateStr] = customerUuid.VisitPageOrderProcessingAmount
             
        }
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "trend": gin.H{
            "pv": pvTrend,
            "stay_seconds": StaySecondsTrend,
            "visit_page_sku": VisitPageSkuTrend,
            "visit_page_category": VisitPageCategoryTrend,
            "visit_page_search": VisitPageSearchTrend,
            "visit_page_cart": VisitPageCartTrend,
            "visit_page_order_pending": VisitPageOrderPendingTrend,
            "visit_page_order_pending_amount": VisitPageOrderPendingAmountTrend,
            "visit_page_order_processing": VisitPageOrderProcessingTrend,
            "visit_page_order_processing_amount": VisitPageOrderProcessingAmountTrend,
            
        },
    })
    // 返回json
    c.JSON(http.StatusOK, result)
    
}










