package customer

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    // "errors"
    // "time"
    "unicode/utf8"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    //"fmt"
)

type Resource struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    UrlKey string `form:"url_key" json:"url_key" binding:"required"`
    RequestMethod int `form:"request_method" json:"request_method" binding:"required"`
    GroupId int64 `form:"group_id" json:"group_id" binding:"required"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
}

func (resource Resource) TableName() string {
    return "resource"
}

/**
 * 增加一条记录
 */
func ResourceAddOne(c *gin.Context){
    var resource Resource
    err := c.ShouldBindJSON(&resource);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    customerId := GetCurrentCustomerId(c)
    
    resource.CreatedCustomerId = customerId
    // 插入
    affected, err := engine.Insert(&resource)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "resource":resource,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过id为条件，更新一条记录
 */
func ResourceUpdateById(c *gin.Context){
    var resource Resource
    err := c.ShouldBindJSON(&resource);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Update(&resource, &Resource{Id:resource.Id})
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "resource":resource,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 删除一条记录
 */
func ResourceDeleteById(c *gin.Context){
    var resource Resource
    var id DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&resource)
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
func ResourceDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var resource Resource
    var ids DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&resource)
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
func ResourceList(c *gin.Context){
    // params := c.Request.URL.Query()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    name     := c.DefaultQuery("name", "")
    url_key  := c.DefaultQuery("url_key", "")
    request_method, _ := strconv.Atoi(c.DefaultQuery("request_method", ""))
    group_id, _ := strconv.Atoi(c.DefaultQuery("group_id", ""))
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    whereParam["name"] = []string{"like", name}
    whereParam["url_key"] = []string{"like", url_key}
    whereParam["request_method"] = request_method
    whereParam["group_id"] = group_id
    whereParam["created_at"] = []string{"scope", created_at_begin, created_at_end}
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
    var resource Resource
    counts, err := engine.Where(whereStr, whereVal...).Count(&resource)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var resources []Resource
    err = query.Find(&resources) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    resourceGroupOps, err := ResourceGroupOps()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    createdCustomerOps, err := GetResCreatedCustomerOps(resources)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": resources,
        "total": counts,
        "reqMethodOps": ReqMethodOps(),
        "createdCustomerOps": createdCustomerOps,
        "resourceGrpOps": resourceGroupOps,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}


func GetResCreatedCustomerOps(resources []Resource) ([]VueSelectOps, error){
    var groupArr []VueSelectOps
    var ids []int64
    for i:=0; i<len(resources); i++ {
        resource := resources[i]
        ids = append(ids, resource.CreatedCustomerId)
    }
    customers, err := GetCustomerUsernameByIds(ids)
    if err != nil{
        return nil, err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        groupArr = append(groupArr, VueSelectOps{Key: customer.Id, DisplayName: customer.Username})
    }
    // if groupArr == nil {
    //     return nil, errors.New("customer ids is empty")
    // }
    return groupArr, nil
}
