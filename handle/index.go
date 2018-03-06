package handle

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/fecshopsoft/fec-go/util"
)


func NotFound(c *gin.Context) {
	//if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		
	//}
    c.AbortWithStatusJSON(http.StatusNotFound, util.BuildFailResult("未知资源"))
}