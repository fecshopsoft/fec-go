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











