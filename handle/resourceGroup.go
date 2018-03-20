package handle

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    // "time"
    "unicode/utf8"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    //"fmt"
)

type ResourceGroup struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int `form:"created_customer_id" json:"created_customer_id"`
}

func (resourceGroup ResourceGroup) TableName() string {
    return "resource_group"
}

/**
 * 增加一条记录
 */
func ResourceGroupAddOne(c *gin.Context){
    var resourceGroup ResourceGroup
    err := c.ShouldBindJSON(&resourceGroup);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    customerId, err := GetCurrentCustomerId()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    resourceGroup.CreatedCustomerId = customerId
    // 插入
    affected, err := engine.Insert(&resourceGroup)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "resourceGroup":resourceGroup,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过id为条件，更新一条记录
 */
func ResourceGroupUpdateById(c *gin.Context){
    var resourceGroup ResourceGroup
    err := c.ShouldBindJSON(&resourceGroup);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Update(&resourceGroup, &ResourceGroup{Id:resourceGroup.Id})
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "resourceGroup":resourceGroup,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 删除一条记录
 */
func ResourceGroupDeleteById(c *gin.Context){
    var resourceGroup ResourceGroup
    var id DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&resourceGroup)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "id":id.Id,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过ids批量删除数据
 */
func ResourceGroupDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var resourceGroup ResourceGroup
    var ids DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&resourceGroup)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected": affected,
        "ids": ids.Ids,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 列表查询
 */
func ResourceGroupList(c *gin.Context){
    // params := c.Request.URL.Query()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    page, _  := strconv.Atoi(c.DefaultQuery("page", listDefaultPage))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", listPageCount))
    name     := c.DefaultQuery("name", "")
    sort     := c.DefaultQuery("sort", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    whereParam["name"] = []string{"like", name}
    whereStr, whereVal := mysqldb.GetXOrmWhere(whereParam)
    // 进行查询
    query := engine.Limit(limit, (page-1)*limit)
    if whereStr != "" {
        query = query.Where(whereStr, whereVal...)
    }
    // 排序
    if sortD == "+" && sortColumns != "" {
        query = query.Asc(sortColumns)
    } else if sortD == "-" && sortColumns != "" {
        query = query.Desc(sortColumns)
    }
    // 得到查询count数
    var resourceGroup ResourceGroup
    counts, err := engine.Where(whereStr, whereVal...).Count(&resourceGroup)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var resourceGroups []ResourceGroup
    err = query.Find(&resourceGroups) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": resourceGroups,
        "total":counts,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}
