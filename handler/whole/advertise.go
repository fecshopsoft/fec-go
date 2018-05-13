package whole

import(
    "github.com/fecshopsoft/fec-go/handler/customer"
    "github.com/fecshopsoft/fec-go/handler/common"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/url"
    "log"
    "github.com/fecshopsoft/fec-go/initialization"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    "github.com/fecshopsoft/fec-go/helper"
)


type Advertise struct {
    Id int64 `form:"id" json:"id"`
    AdvertiseId string `form:"advertise_id" json:"advertise_id" `
    FecSource string `form:"fec_source" json:"fec_source"`  // binding:"required"
    FecMedium string `form:"fec_medium" json:"fec_medium"`
    FecCampaign string `form:"fec_campaign" json:"fec_campaign"`
    FecContent string `form:"fec_content" json:"fec_content"`
    FecDesign string `form:"fec_design" json:"fec_design"`
    Url string `form:"url" json:"url"`
    AdvertiseCost float64 `form:"advertise_cost" json:"advertise_cost"`
    Remark string `form:"remark" json:"remark"`
    AdvertiseUrl string `form:"advertise_url" json:"advertise_url"`
    
    MarketGroup int64 `form:"market_group" json:"market_group"`
    OwnId int64 `form:"own_id" json:"own_id"`
    
    AdvertiseBeginDate int64 `xorm:"created"  form:"advertise_begin_date" json:"advertise_begin_date"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
}

func (advertise Advertise) TableName() string {
    return "advertise"
}

// 得到 trend  info
func AdvertiseInit(c *gin.Context){
    var channelChildOps []helper.VueSelectStrOps
    channel := c.DefaultQuery("channel", "")
    channel_child := c.DefaultQuery("channel_child", "")
    is_create := c.DefaultQuery("is_create", "")
    log.Println("AdvertiseInit###########")
    // 得到用户的主id
    main_id := customer.GetCustomerMainId(c)
    log.Println(main_id)
    // 通过主id，在表：base_channel_info 通过own_id查询，得到channel数据
    channelOps, err := common.GetChannelOpsByOwnId(main_id)
    log.Println("######1")
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    if channel != "" {
        channelChildOps, err = common.GetChannelChildOpsByOwnIdAndChannel(main_id, channel)
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return  
        }
    }
    log.Println("######2")
    // 得到design person
    var designOptions []helper.VueSelectOps
    if is_create == "1" {
        customers, err := customer.GetDesigiPersonByOwnId(main_id)
        if err != nil{
            c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
            return  
        }
        for i:=0; i<len(customers); i++ {
            customer := customers[i]
            designOptions = append(designOptions, helper.VueSelectOps{Key: customer.Id, DisplayName: customer.Username})
        }
    }
     log.Println("######3")
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "channel": channel,
        "channelOptions": channelOps,
        "channel_child": channel_child,
        "channelChildOptions": channelChildOps,
        "designOptions": designOptions,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
    
}


func AdvertiseGenerateUrl(c *gin.Context){
    engine := mysqldb.GetEngine()
    errStr := ""
    warnStr := ""
    successInfo := gin.H{}
    channel         := c.DefaultQuery("channel", "")
    channel_child   := c.DefaultQuery("channel_child", "")
    campaign        := c.DefaultQuery("campaign", "")
    design_person   := c.DefaultQuery("design", "")
    advertise_url   := c.DefaultQuery("advertise_url", "")
    advertise_cost  := c.DefaultQuery("advertise_cost", "")
    remark          := c.DefaultQuery("remark", "")
    // 判断必填参数是否缺失
    if advertise_url == "" {
        errStr += "广告链接URL不能为空，"
    }
    _, err := url.Parse(advertise_url)
    if err != nil {
        // 判断格式
        // 生成返回结果
        errStr += "广告链接URL格式不正确，格式：http://www.fecshop.com/xxxx，"
    }
    
    if channel == "" {
        errStr += "渠道不能为空，"
    }
    if channel_child == "" {
        errStr += "子渠道不能为空，"
    }
    if campaign == "" {
        warnStr += "活动没有填写（可选），"
    }
    
    
    if design_person == "" {
        warnStr += "广告图片设计师没有填写（可选），"
    } 
    
    var advertiseCost float64
    if advertise_cost == "" {
        warnStr += "广告费用没有填写（可选），"
    } else {
        advertiseCost, err = helper.Float64(advertise_cost)
        if err != nil {
            errStr += "广告费用格式请填写数字（" + err.Error() + "）,"
        }
    }
    if advertiseCost == 0 {
    
    } 
    
    if remark == "" {
        warnStr += "广告备注没有填写（可选），"
    }
    
    
    // 其他信息
    currentCustomerId := helper.GetCurrentCustomerId(c)
    own_id := customer.GetCustomerMainId(c)
    
    customerOne, err := customer.GetCustomerOneById(currentCustomerId)
    if err != nil {
        errStr += err.Error()
    }
    marketGroupId := customerOne.MarketGroupId
    // 写入数据库
    var advertise Advertise
    if errStr == "" {
        advertise.FecSource = channel
        advertise.FecMedium = channel_child
        advertise.FecCampaign = campaign
        advertise.FecContent = helper.Str64(currentCustomerId)
        advertise.FecDesign = design_person
        advertise.Url = advertise_url
        advertise.AdvertiseCost = advertiseCost
        advertise.Remark = remark
        advertise.MarketGroup = int64(marketGroupId)  //
        advertise.OwnId = own_id
        advertise.CreatedCustomerId =  currentCustomerId //
        // 插入
        affected, err := engine.Insert(&advertise)
        if err != nil {
            errStr += err.Error()
        } else if affected <= 0 {
            errStr += "插入失败"
        } else {
            id := advertise.Id
            advertise_id := int64(100000000) + id
            advertise.AdvertiseId = helper.Str64(advertise_id)
            // url 添加参数
            u, _ := url.Parse(advertise.Url)
            q, _ := url.ParseQuery(u.RawQuery)
            
            q.Add("fec_source",advertise.FecSource)
            q.Add("fec_medium",advertise.FecMedium)
            if advertise.FecCampaign != "" {
                q.Add("fec_campaign",advertise.FecCampaign)
            }
            if advertise.FecContent != "" {
                q.Add("fec_content",advertise.FecContent)
            }
            if advertise.FecDesign != "" {
                q.Add("fec_design",advertise.FecDesign)
            }
            q.Add("fid",advertise.AdvertiseId)
            
            u.RawQuery = q.Encode()
            advertise.AdvertiseUrl = u.String()
            affected, err := engine.Where("id = ?", id).MustCols("advertise_id, advertise_url").Update(&advertise)
            if err != nil {
                errStr += err.Error()
            } else if affected <= 0 {
                errStr += "update 失败"
            } else {
                // successStr += "success"
                fecDesignInt64, _ := helper.Int64(advertise.FecDesign)
                fecContent64, _ := helper.Int64(advertise.FecContent)
                
                successInfo = gin.H{
                    "advertise_url": advertise.AdvertiseUrl,
                    "created_at": advertise.CreatedAt,
                    "origin_url": advertise_url,
                    "fid": advertise_id,
                    "channel": advertise.FecSource,
                    "channel_child": advertise.FecMedium,
                    "campaign": advertise.FecCampaign,
                    "design_person": initialization.CustomerIdWithName[fecDesignInt64],
                    "advertise_person": initialization.CustomerIdWithName[fecContent64],
                    "advertise_market_group": initialization.MarketGroupIdWithName[advertise.MarketGroup],
                    "advertise_cost": advertise.AdvertiseCost,
                    "advertise_remark": advertise.Remark,
                }
                
                
            }           
        
        }
        
    }
    
    // 字符串拼接，生成url  QueryEscape(s string) string
    // 每个参数值都用 url.QueryEscape(traceGetInfo.Search)
    
    // 生成一个广告id（fid) 作为广告的唯一值。
    // fecshop.appfront.fancyecommerce.com?fid=1111&fec_source=2222&fec_medium=3333&fec_campaign=4444&fec_content=5555&fec_design=6666
    // fec_source 渠道，fec_medium 子渠道， fec_campaign 活动， fec_content 员工，fec_design 美工
    // successStr += "success"
    result := util.BuildSuccessResult(gin.H{
        "error": errStr,
        "warning": warnStr,
        "success_info": successInfo,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}











