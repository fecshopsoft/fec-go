package cron

import(
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
   
    "github.com/fecshopsoft/fec-go/initialization"
)


func UpdateSite(c *gin.Context){
    err := initialization.InitWebsiteInfo()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "success": "success",
        "WebsiteInfos": initialization.WebsiteInfos,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}



/*

*/