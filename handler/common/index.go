package common

import(
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    "github.com/go-xorm/xorm"
    // "github.com/gin-gonic/gin"
    // "github.com/fecshopsoft/fec-go/helper"
    // "github.com/fecshopsoft/fec-go/handler/customer"
    "log"
)

var engine *(xorm.Engine)

// init 函数在程序启动时执行，后面不会再执行。
func init(){
    engine = mysqldb.GetEngine()
    log.Println("Base Info GetEngine complete")
}

type VueSelectOps struct{
    Key string `form:"key" json:"key"`
    DisplayName string `form:"display_name" json:"display_name"`
}

