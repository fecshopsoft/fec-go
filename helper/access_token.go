package helper

import(
    "github.com/gin-gonic/gin"
)

 
// 得到当前的 customerUsername
func GetCurrentWebsiteId(c *gin.Context) (string){
    return c.GetString("currentWebsiteId")
}  


