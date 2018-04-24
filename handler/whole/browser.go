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

func BrowserList(c *gin.Context){
    
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    sort     := c.DefaultQuery("sort", "")
    sort_dir := c.DefaultQuery("sort_dir", "")
    
    
    service_date_str_begin := c.DefaultQuery("service_date_str_begin", "")
    service_date_str_end := c.DefaultQuery("service_date_str_end", "")
    browser_name := c.DefaultQuery("browser_name", "")
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
    // browser_name 搜索
    if browser_name != "" {
        q = q.Must(elastic.NewTermQuery("browser_name", browser_name))
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
    //////
    chosen_own_id, chosen_website_id, selectOwnIds, selectWebsiteIds, err := ActiveOwnIdAndWebsite(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    esIndexName := helper.GetEsIndexName(chosen_website_id)
    esWholeBrowserTypeName :=  helper.GetEsWholeBrowserTypeName()
    client, err := esdb.Client()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // q = q.Must(elastic.NewTermQuery("browser_name", "Safari"))
    // q = q.Must(elastic.NewRangeQuery("pv").From(3).To(60))
    // q = q.Must(elastic.NewRangeQuery("service_date_str").Gt("2018-04-20").Lt("2018-04-21"))
    //termQuery := elastic.NewTermQuery("browser_name", "Safari")
    // termQuery := elastic.NewRangeQuery("browser_name", "Safari")
    //rangeQuery := NewRangeQuery("pv").Gt(3)
    log.Println(8888888888888)
    log.Println(esIndexName)
    log.Println(esWholeBrowserTypeName)
    log.Println(page-1)
    log.Println(limit)
    log.Println(sort)
    search := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esWholeBrowserTypeName).
        Query(q).
        From(page-1).Size(limit).
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
        Type(esWholeBrowserTypeName).
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
    
    
    var ts []model.WholeBrowserValue
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var wholeBrowser model.WholeBrowserValue
            err := json.Unmarshal(*hit.Source, &wholeBrowser)
            if err != nil{
                c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
                return
            }
            ts = append(ts, wholeBrowser)
        }
    }
    ownNameOptions, err := getOwnNames(c, selectOwnIds)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    siteNameOptions, err := getSiteNames(c, selectWebsiteIds)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "total": searchResult.Hits.TotalHits,
        "items": ts,
        "chosen_own_id": chosen_own_id,
        "chosen_website_id": chosen_website_id,
        "selectOwnIds": selectOwnIds,
        "selectWebsiteIds": selectWebsiteIds,
        "ownNameOptions": ownNameOptions,
        "siteIdOptions": siteNameOptions,
        
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}
















