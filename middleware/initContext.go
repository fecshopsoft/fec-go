package middleware

import(
    "github.com/gin-gonic/gin"
    // "net/http"
    // "github.com/fecshopsoft/fec-go/util"
)

func InitContext(c *gin.Context) {
    // 设置默认第几页
    c.Set("defaultPageNum", "1")
    // 设置每页显示的默认个数
    c.Set("defaultPageCount", "20")
}