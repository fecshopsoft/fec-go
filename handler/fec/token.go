package fec

import(
    "github.com/gin-gonic/gin"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/handler/common"
    "net/http"
)


func PermissionAccessToken(c *gin.Context){
    
    access_token := helper.GetHeader(c, "Access-Token");
    if  access_token == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("Access-Token can not empty"))
        return
	}
    
    website_id, err := helper.GetSiteUIdByAccessToken(access_token)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    if  website_id == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("can not get website_id by Access-Token"))
        return
	}
    // 数据库查询，该token是否有效
    _, err = common.GetWebsiteByAccessToken(access_token)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 设置 currentWebsiteId
    c.Set("currentWebsiteId", website_id)
    
}



