package whole

import(
    "github.com/fecshopsoft/fec-go/handler/customer"
    "github.com/fecshopsoft/fec-go/handler/common"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/url"
    "github.com/fecshopsoft/fec-go/helper"
)



// 得到 trend  info
func AdvertiseInit(c *gin.Context){
    var channelChildOps []helper.VueSelectStrOps
    channel := c.DefaultQuery("channel", "")
    channel_child := c.DefaultQuery("channel_child", "")
    is_create := c.DefaultQuery("is_create", "")
    
    // 得到用户的主id
    main_id := customer.GetCustomerMainId(c)
    // 通过主id，在表：base_channel_info 通过own_id查询，得到channel数据
    channelOps, err := common.GetChannelOpsByOwnId(main_id)
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
    channel         := c.DefaultQuery("channel", "")
    channel_child   := c.DefaultQuery("channel_child", "")
    campaign        := c.DefaultQuery("campaign", "")
    design_person   := c.DefaultQuery("design", "")
    advertise_url   := c.DefaultQuery("advertise_url", "")
    advertise_cost  := c.DefaultQuery("advertise_cost", "")
    remark          := c.DefaultQuery("remark", "")
    u, err := url.Parse(advertise_url)
    if err != nil {
        // 判断格式
    }
    // 判断必填参数是否缺失
    
    // 字符串拼接，生成url  QueryEscape(s string) string
    // 每个参数值都用 url.QueryEscape(traceGetInfo.Search)
    
    // 生成一个广告id（fid) 作为广告的唯一值。
    // fecshop.appfront.fancyecommerce.com?fid=1111&fec_source=2222&fec_medium=3333&fec_campaign=4444&fec_content=5555&fec_design=6666
    // fec_source 渠道，fec_medium 子渠道， fec_campaign 活动， fec_content 员工，fec_design 美工
}











