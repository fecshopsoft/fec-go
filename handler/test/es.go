package test

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/model"
    "context"
    // "reflect"
    "encoding/json"
    "github.com/olivere/elastic"
    "github.com/fecshopsoft/fec-go/db/esdb"
)



func EsFind(c *gin.Context){
    esIndexName := helper.GetEsIndexName("9b17f5b4-b96f-46fd-abe6-a579837ccdd9")
    esWholeBrowserTypeName :=  helper.GetEsWholeBrowserTypeName()
    client, err := esdb.Client()
    if err != nil {
        // Handle error
        panic(err)
    }
    q := elastic.NewBoolQuery()
    // q = q.Must(elastic.NewTermQuery("browser_name", "Safari"))
    // q = q.Must(elastic.NewRangeQuery("pv").From(3).To(60))
    // q = q.Must(elastic.NewRangeQuery("service_date_str").Gt("2018-04-20").Lt("2018-04-21"))
    //termQuery := elastic.NewTermQuery("browser_name", "Safari")
    // termQuery := elastic.NewRangeQuery("browser_name", "Safari")
    //rangeQuery := NewRangeQuery("pv").Gt(3)
    searchResult, err := client.Search().
        Index(esIndexName).        // search in index "twitter"
        Type(esWholeBrowserTypeName).
        Query(q).        // specify the query
        //Sort("user", true).      // sort by "user" field, ascending
        From(0).Size(10).        // take documents 0-9
        Pretty(true).            // pretty print request and response JSON
        Do(context.Background()) // execute
    if err != nil {
        // Handle error
        panic(err)
    }
    
    
    var ts []model.WholeBrowserValue
    if searchResult.Hits.TotalHits > 0 {
        // Iterate through results
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index

            // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
            var wholeBrowser model.WholeBrowserValue
            err := json.Unmarshal(*hit.Source, &wholeBrowser)
            if err != nil {
                // Deserialization failed
            }
            ts = append(ts, wholeBrowser)
        }
    }
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "searchResult.Hits.TotalHits": searchResult.Hits.TotalHits,
        "searchResult": searchResult,
        "ts": ts,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}


