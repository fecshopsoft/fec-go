package helper

import(
    "github.com/gin-gonic/gin"
    //"fmt"
)

// 定义当前用户 废弃
// var currentCustomer interface{}

// 从header中取出来相关的数据
func GetHeader(c *gin.Context, key string) string{
    if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}