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

// 同步数据
func UpsertType(esIndexName string, esTypeName string, idStr string, bodyJson interface{}) error{
    log.Println("5555" + esIndexName)
    log.Println("555" + esTypeName)
    log.Println("55" + idStr)
    log.Println(bodyJson)
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


