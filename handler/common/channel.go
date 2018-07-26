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

type ChannelInfo struct {
    Id int64 `form:"id" json:"id"`
    Channel string `form:"channel" json:"channel" binding:"required"`
    ChannelChild string `form:"channel_child" json:"channel_child" binding:"required"`
    //OwnId int64 `form:"own_id" json:"own_id"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
}

func (channelInfo ChannelInfo) TableName() string {
    return "base_channel_info"
}


/**
 * 增加一条记录
 */
func ChannelAddOne(c *gin.Context){
    var channelInfo ChannelInfo
    err := c.ShouldBindJSON(&channelInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    customerId := helper.GetCurrentCustomerId(c)
    channelInfo.CreatedCustomerId = customerId
    // 插入
    affected, err := engine.Insert(&channelInfo)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "channelInfo":channelInfo,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 通过id为条件，更新一条记录
 */
func ChannelUpdateById(c *gin.Context){
    var channelInfo ChannelInfo
    err := c.ShouldBindJSON(&channelInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    
    // 更新
    affected, err := engine.Where("id = ?",channelInfo.Id).Cols("channel, channel_child, updated_at").Update(&channelInfo)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "channelInfo":channelInfo,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 删除一条记录
 */
func ChannelDeleteById(c *gin.Context){
    var channelInfo ChannelInfo
    var id helper.DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&channelInfo)
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
func ChannelDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var channelInfo ChannelInfo
    var ids helper.DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&channelInfo)
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
func ChannelList(c *gin.Context){
    // params := c.Request.URL.Query()
    
    // 获取参数并处理
    var sortD string
    var sortColumns string
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    channel     := c.DefaultQuery("channel", "")
    
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    if channel != "" {
        whereParam["channel"] = []string{"like", channel}
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
    var channelInfo ChannelInfo
    counts, err := engine.Where(whereStr, whereVal...).Count(&channelInfo)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var channelInfos []ChannelInfo
    err = query.Find(&channelInfos) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    
    createdCustomerOps, err := GetChannelCreatedCustomerOps(channelInfos)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": channelInfos,
        "total": counts,
        "createdCustomerOps": createdCustomerOps,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}




// 通过创建人ids，得到创建人的name
func GetChannelCreatedCustomerOps(channelInfos []ChannelInfo) ([]helper.VueSelectOps, error){
    var groupArr []helper.VueSelectOps
    var ids []int64
    for i:=0; i<len(channelInfos); i++ {
        channelInfo := channelInfos[i]
        ids = append(ids, channelInfo.CreatedCustomerId)
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

// 得到channel child Ops
func GetChannelChildOpsByChannel(channel string) ([]helper.VueSelectStrOps, error){
    var groupArr []helper.VueSelectStrOps
    channelInfos, err := GetChildChannelByChannel(channel)
    if err != nil{
        return nil, err
    }
    for i:=0; i<len(channelInfos); i++ {
        channelInfo := channelInfos[i]
        groupArr = append(groupArr, helper.VueSelectStrOps{Key: channelInfo.ChannelChild, DisplayName: channelInfo.ChannelChild})
    }
    return groupArr, nil

} 
// 得到channel ops
func GetChannelOps()([]helper.VueSelectStrOps, error){
    var groupArr []helper.VueSelectStrOps
    channelInfos, err := GetChannels()
    if err != nil{
        return nil, err
    }
    s := make(map[string]string)
    for i:=0; i<len(channelInfos); i++ {
        channelInfo := channelInfos[i]
        if _, ok := s[channelInfo.Channel]; !ok {
            s[channelInfo.Channel] = channelInfo.Channel
            groupArr = append(groupArr, helper.VueSelectStrOps{Key: channelInfo.Channel, DisplayName: channelInfo.Channel})
        }
    }
    return groupArr, nil
}

/**
 * 查询得到 ChannelInfo
 */
func GetChannels() ([]ChannelInfo, error){
    // 得到结果数据
    var channelInfos []ChannelInfo
    err := engine.Find(&channelInfos) 
    if err != nil{
        return channelInfos, err
    }
    return channelInfos, nil
}

/**
 * channel 查询得到 ChannelInfo
 */
func GetChildChannelByChannel(channel string) ([]ChannelInfo, error){
    // 得到结果数据
    var channelInfos []ChannelInfo
    err := engine.Where("channel = ? ", channel).Find(&channelInfos) 
    if err != nil{
        return channelInfos, err
    }
    return channelInfos, nil
}
 

/**
 * 根据 market_group_ids 查询得到 MarketGroup

func GetChannelByIds(market_group_ids []int64) ([]MarketGroup, error){
    // 得到结果数据
    var marketGroups []MarketGroup
    err := engine.In("id",market_group_ids).Find(&marketGroups) 
    if err != nil{
        return marketGroups, err
    }
    return marketGroups, nil
}
 */