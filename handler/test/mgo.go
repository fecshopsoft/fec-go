package test

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "net/http"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
)

type ProductImg struct {
    Gallery []map[string]interface{} `bson:"gallery"` 
    Main map[string]interface{} `bson:"main"` 
}

type ProductFlat struct { 
    Id_ bson.ObjectId `bson:"_id"` 
    Spu string `bson:"spu"` 
    Sku string `bson:"sku"` 
    Name map[string]interface{} `bson:"name"` 
    CreatedAt int64 `bson:"created_at"` 
    Category string `bson:"category"` 
    Image ProductImg `bson:"image"` 
}
func (productFlat ProductFlat) TableName() string {
    return "product_flat"
}

func MgoFind(c *gin.Context){
    // session, err := mgo.Dial("127.0.0.1:27017")
    // if err != nil {
    //    c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
    //    return
    // } 
    // db := session.DB("fecshop_demo")
    // coll := db.C("product_flat")
    var productFlat ProductFlat
    var productFlats []ProductFlat
    // collName := "product_flat"
    
    mongodb.MC(productFlat.TableName(), func(coll *mgo.Collection) error {
        // c.Find(M{"_id": id}).One(m)
        // Log("find one m", m, m["name"], m["img"])
        coll.Find(nil).All(&productFlats) 
        var err error
        return err
    })
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "productFlats": productFlats,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}


func MgoMapReduce(c *gin.Context){
    mapStr := `
        function() { 
            // emit(this.uuid, {uuid: 1});
            if (this.browser_name) {
                emit(this.browser_name, {browser_name: 1});
            }
        }
    `
    
    reduceStr := `
        function(key, emits) { 
            // this_uuid_count = 0; 
            var this_browser_name = 0
            for(var i in emits){
                if( emits[i].browser_name){
                    this_browser_name 		+=  emits[i].browser_name;
                }
            }  
            return {
                browser_name: this_browser_name
            };
        }
    
    `
    job := &mgo.MapReduce{
        Map:      mapStr,
        Reduce:   reduceStr,
    }
    type resultValue struct{
        BrowserNameCount int64 `browser_name`
    }
    var result []struct { 
        Id string `_id`
        Value resultValue 
    }
    err := mongodb.MC("trace_info", func(coll *mgo.Collection) error {
        _, err := coll.Find(nil).MapReduce(job, &result)
        return err
    })
    
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    // for _, item := range result {
    //    fmt.Println(item.Value)
    // }
    // 生成返回结果
    re := util.BuildSuccessResult(gin.H{
        "success": "success",
        "productFlats": result,
    })
    // 返回json
    c.JSON(http.StatusOK, re)
}






