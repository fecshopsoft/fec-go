package esdb
import (
    "context"
	// "encoding/json"
	"log"
	// "reflect"
	// "time"
    "sync"
    "errors"
    "github.com/fecshopsoft/fec-go/helper"
	"github.com/olivere/elastic"
)

var once sync.Once
var esClient *(elastic.Client)
var esUrl string = "http://127.0.0.1:9200"


func Client() (*(elastic.Client), error){
    var err error
    once.Do(func() {
        // Starting with elastic.v5, you must pass a context to execute each service
        ctx := context.Background()

        // Obtain a client and connect to the default Elasticsearch installation
        // on 127.0.0.1:9200. Of course you can configure your client to connect
        // to other hosts and configure it in various other ways.
        /*
        client, err := elastic.NewClient(
            elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
            elastic.SetBasicAuth("user", "secret"))

        */
        esClient, err = elastic.NewClient(elastic.SetURL(esUrl))
        if err != nil {
            return 
        }

        // Ping the Elasticsearch server to get e.g. the version number
        info, code, err := esClient.Ping(esUrl).Do(ctx)
        if err != nil {
            return 
        }
        log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

        // Getting the ES version number is quite common, so there's a shortcut
        esversion, err := esClient.ElasticsearchVersion(esUrl)
        if err != nil {
            return 
        }
        log.Printf("Elasticsearch version %s\n", esversion)
    })
    
    return esClient, err
}

// 创建 elasticSearch 的 Mapping
func InitMapping(esIndexName string, esTypeName string, typeMapping string) error{
    var err error
    indexMapping := helper.GetEsIndexMapping()
    ctx := context.Background()
    client, err := Client()
    if err != nil {
		return err
	}
    // Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(esIndexName).Do(ctx)
	if err != nil {
        log.Println("IndexExists" + err.Error())
		return err
	}
    //log.Println("es index: " + esIndexName)
    //log.Println("es type: " + esTypeName)
    //log.Println("es index mapping: " + indexMapping)
    //log.Println("es type mapping: " + typeMapping)
	if !exists {
        log.Println("es index not exists: " + esIndexName)
		// Create a new index.
		createIndex, err := client.CreateIndex(esIndexName).Body(indexMapping).Do(ctx)
		if err != nil {
            log.Println("CreateIndex" + err.Error())
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
            return errors.New("create index:" + esIndexName + ", not Ack nowledged")
		}
	}
    /**
     * 判断 type 是否存在
        exists, err = client.TypeExists().Index(esIndexName).Type(esTypeName).Do(ctx)
        if err != nil {
            return err
        }
        if !exists {
        
        }
    */
    // PutMapping() *IndicesPutMappingService
     
    putresp, err := client.PutMapping().Index(esIndexName).Type(esTypeName).BodyString(typeMapping).Do(context.TODO())
    // 新建 mapping
    //indicesCreateResult, err := elastic.NewIndicesCreateService(client).Index(esIndexName).BodyString(mapping).Do(ctx)
    if err != nil {
        log.Println("NewIndicesCreateService" + err.Error())
        return err
    }
    if !putresp.Acknowledged {
        // Not acknowledged
        return errors.New("create mapping fail, esIndexName:" + esIndexName + ", esTypeName:" + esTypeName + ", not Ack nowledged")
    }
    
    // 插入数据
    /*
    type WholeBrowserData struct {
        BrowserId     string                `json:"browser_id"`
        BrowserName  string                `json:"browser_name"`
    }

    // Index a tweet (using JSON serialization)
	wholeBrowserData := WholeBrowserData{BrowserId: "BrowserId", BrowserName: "BrowserName" }
	put1, err := client.Index().
		Index(esIndexName).
		Type(esTypeName).
		Id("1").
		BodyJson(wholeBrowserData).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
    */
    
    
    return err
}

// https://github.com/olivere/elastic/issues/105
// https://github.com/olivere/elastic/wiki/QueryDSL
// https://www.elastic.co/guide/cn/elasticsearch/guide/current/combining-filters.html
func FilterList(){


}

/* 单个更新数据
for j:=0; j<len(wholeBrowsers); j++ {
    wholeBrowser := wholeBrowsers[j]
    wholeBrowserValue := wholeBrowser.Value
    // wholeBrowserValue.Devide = nil
    // wholeBrowserValue.CountryCode = nil
    ///wholeBrowserValue.Operate = nil
    log.Println("ID_:" + wholeBrowser.Id_)
    wholeBrowserValue.Id = wholeBrowser.Id_
    err := esdb.UpsertType(esIndexName, esWholeBrowserTypeName, wholeBrowser.Id_, wholeBrowserValue)
    
    if err != nil {
        log.Println("11111" + err.Error())
        return err
    }
}
*/
func UpsertType(esIndexName string, esTypeName string, idStr string, bodyJson interface{}) error{
    var err error
    ctx := context.Background()
    client, err := Client()
    upsertResult, err := client.Index().
		Index(esIndexName).
		Type(esTypeName).
		Id(idStr).
		BodyJson(bodyJson).
		Do(ctx)
	if err != nil {
		return err
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", upsertResult.Id, upsertResult.Index, upsertResult.Type)
    return err
}

/**
 * 批量操作
 *
    bulkRequest, err := esdb.Bulk()
    if err != nil {
        return err
    }
    for j:=0; j<len(wholeBrowsers); j++ {
        wholeBrowser := wholeBrowsers[j]
        wholeBrowserValue := wholeBrowser.Value
        wholeBrowserValue.Id = wholeBrowser.Id_
        req := esdb.BulkUpsertTypeDoc(esIndexName, esWholeBrowserTypeName, wholeBrowser.Id_, wholeBrowserValue)
        bulkRequest = bulkRequest.Add(req)
    }
    bulkResponse, err := esdb.BulkRequestDo(bulkRequest)
    // bulkResponse, err := bulkRequest.Do()
    if err != nil {
        return err
    }
    if bulkResponse != nil {
        log.Println(bulkResponse)
    }
 *
 */
// 获取批量执行的bulk
func Bulk() (*(elastic.BulkService), error){
    client, err := Client()
    return client.Bulk(), err
}
// 获取批量更新的doc
func BulkUpsertTypeDoc(esIndexName string, esTypeName string, idStr string, docStruct interface{}) *(elastic.BulkUpdateRequest){

    req := elastic.NewBulkUpdateRequest().
            Index(esIndexName).
            Type(esTypeName).
            Id(idStr).
            Doc(docStruct).
            DocAsUpsert(true)
    return req
}
// 批量执行
func BulkRequestDo(bulkRequest *(elastic.BulkService)) (*(elastic.BulkResponse), error){
    ctx := context.Background()
    bulkResponse, err := bulkRequest.Do(ctx)
    return bulkResponse, err
}


// 通过indexName 删除 Index
func DeleteIndex(indexName string) error {
    var err error
    ctx := context.Background()
    client, err := Client()
    if err != nil {
		return err
	}
    exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
        log.Println("IndexExists" + err.Error())
		return err
	}
    if exists {
        deleteIndex, err := client.DeleteIndex(indexName).Do(ctx)
        if err != nil {
            // Handle error
            return err
        }
        if !deleteIndex.Acknowledged {
            // Not acknowledged
            return errors.New("delete index is not Acknowledged")
        }
    }
    return err
}


