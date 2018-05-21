package common

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
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/handler/customer"
    //"fmt"
)

type MarketGroup struct {
    Id int64 `form:"id" json:"id"`
    Name string `form:"name" json:"name" binding:"required"`
    OwnId int64 `form:"own_id" json:"own_id"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
}

func (marketGroup MarketGroup) TableName() string {
    return "base_market_group"
}


/**
 * 增加一条记录
 */
func MarketGroupAddOne(c *gin.Context){
    var marketGroup MarketGroup
    err := c.ShouldBindJSON(&marketGroup);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 处理own_id
    own_id, err := customer.Get3SaveDataOwnId(c, marketGroup.OwnId)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    marketGroup.OwnId = own_id
    customerId := helper.GetCurrentCustomerId(c)
    marketGroup.CreatedCustomerId = customerId
    // 插入
    affected, err := engine.Insert(&marketGroup)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "marketGroup":marketGroup,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过id为条件，更新一条记录
 */
func MarketGroupUpdateById(c *gin.Context){
    var marketGroup MarketGroup
    err := c.ShouldBindJSON(&marketGroup);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 处理own_id
    own_id, err := customer.Get3SaveDataOwnId(c, marketGroup.OwnId)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    marketGroup.OwnId = own_id
    // 更新
    affected, err := engine.Where("id = ?",marketGroup.Id).Cols("name,own_id,updated_at").Update(&marketGroup)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "marketGroup":marketGroup,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 删除一条记录
 */
func MarketGroupDeleteById(c *gin.Context){
    var marketGroup MarketGroup
    var id helper.DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&marketGroup)
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
func MarketGroupDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var marketGroup MarketGroup
    var ids helper.DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&marketGroup)
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
func MarketGroupList(c *gin.Context){
    // params := c.Request.URL.Query()
    
    // 获取参数并处理
    var sortD string
    var sortColumns string
    var own_id int64
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    name     := c.DefaultQuery("name", "")
    own_id_i, _ := strconv.Atoi(c.DefaultQuery("own_id", ""))
    own_id = int64(own_id_i)
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    if name != "" {
        whereParam["name"] = []string{"like", name}
    }  
    own_id, err := customer.Get3OwnId(c, own_id)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    if own_id != 0 {
        whereParam["own_id"] = own_id
    }
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
    var marketGroup MarketGroup
    counts, err := engine.Where(whereStr, whereVal...).Count(&marketGroup)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var marketGroups []MarketGroup
    err = query.Find(&marketGroups) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    ownNameOps, err := customer.Get3OwnNameOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    createdCustomerOps, err := GetMGCreatedCustomerOps(marketGroups)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": marketGroups,
        "total": counts,
        "createdCustomerOps": createdCustomerOps,
        "ownNameOps": ownNameOps,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

func GetAllMarketGroup() ([]MarketGroup, error){
    // 得到结果数据
    var marketGroups []MarketGroup
    err := engine.Find(&marketGroups) 
    return marketGroups, err
}


// 通过创建人ids，得到创建人的name
func GetMGCreatedCustomerOps(marketGroups []MarketGroup) ([]helper.VueSelectOps, error){
    var groupArr []helper.VueSelectOps
    var ids []int64
    for i:=0; i<len(marketGroups); i++ {
        marketGroup := marketGroups[i]
        ids = append(ids, marketGroup.CreatedCustomerId)
    }
    customers, err := customer.GetCustomerUsernameByIds(ids)
    if err != nil{
        return nil, err
    }
    for i:=0; i<len(customers); i++ {
        customer := customers[i]
        groupArr = append(groupArr, helper.VueSelectOps{Key: customer.Id, DisplayName: customer.Username})
    }
    return groupArr, nil
}


/**
 * 根据 market_group_ids 查询得到 MarketGroup
 */
func GetMarketGroupByIds(market_group_ids []int64) ([]MarketGroup, error){
    // 得到结果数据
    var marketGroups []MarketGroup
    err := engine.In("id",market_group_ids).Find(&marketGroups) 
    if err != nil{
        return marketGroups, err
    }
    return marketGroups, nil
}



/**
 * 根据 OwnId 查询得到 MarketGroups
 */
func GetMarketGroupByOwnId(own_id int64) ([]MarketGroup, error){
    // 得到结果数据
    var marketGroups []MarketGroup
    err := engine.Where(" own_id = ? ",own_id).Find(&marketGroups) 
    if err != nil{
        return marketGroups, err
    }
    return marketGroups, nil
}





