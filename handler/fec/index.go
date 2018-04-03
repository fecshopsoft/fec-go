package fec

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SaveJsData(c *gin.Context){
    
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}