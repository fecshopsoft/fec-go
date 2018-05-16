package whole

import(
    "github.com/fecshopsoft/fec-go/handler/customer"
    "github.com/fecshopsoft/fec-go/handler/common"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/config"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "net/url"
    "log"
    "unicode/utf8"
    "github.com/tealeg/xlsx"
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
            designOptions = append(designOptions, helper.VueSelectOps{Key: customer.Id, DisplayName: customer.Name})
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
    errStr := ""
    channel         := c.DefaultQuery("channel", "")
    channel_child   := c.DefaultQuery("channel_child", "")
    campaign        := c.DefaultQuery("campaign", "")
    design_person   := c.DefaultQuery("design", "")
    advertise_url   := c.DefaultQuery("advertise_url", "")
    advertise_cost  := c.DefaultQuery("advertise_cost", "")
    remark          := c.DefaultQuery("remark", "")
    
    // 其他信息
    currentCustomerId := helper.GetCurrentCustomerId(c)
    own_id := customer.GetCustomerMainId(c)
    
    customerOne, err := customer.GetCustomerOneById(currentCustomerId)
    if err != nil {
        errStr += err.Error()
    }
    marketGroupId := customerOne.MarketGroupId
    
    errStr, warnStr, successInfo, _ := generateAdvertiseInfo(channel, channel_child, campaign, design_person, advertise_url,  advertise_cost, remark, currentCustomerId, own_id, marketGroupId, errStr, "")
    result := util.BuildSuccessResult(gin.H{
        "error": errStr,
        "warning": warnStr,
        "success_info": successInfo,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

func generateAdvertiseInfo(channel string, channel_child string, campaign string, design_person string, advertise_url string,  advertise_cost string, remark string, currentCustomerId int64, own_id int64, marketGroupId int64, errStr string, fid string) (string, string, gin.H, string) {
    engine := mysqldb.GetEngine()
    warnStr := ""
    successInfo := gin.H{}
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
            //var advertise_id int64
            if fid == "" {
                //advertise_id = int64(1000000000) + id
                //advertise_id_str := helper.Str64(advertise_id)
                advertise.AdvertiseId = helper.RandomUUID()
                fid = advertise.AdvertiseId
            } else {
                advertise.AdvertiseId = fid
            }
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
                    "fid": fid,
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
    /*
    return gin.H{
        "error": errStr,
        "warning": warnStr,
        "success_info": successInfo,
    }
    */
    return errStr, warnStr, successInfo, fid

}


/**
 * 列表查询
 */
func AdvertiseList(c *gin.Context){
    engine := mysqldb.GetEngine()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    var own_id int64
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    advertise_id     := c.DefaultQuery("advertise_id", "")
    fec_source     := c.DefaultQuery("fec_source", "")
    fec_medium     := c.DefaultQuery("fec_medium", "")
    own_id_i, _ := strconv.Atoi(c.DefaultQuery("own_id", ""))
    own_id = int64(own_id_i)
    // own_id = customer.GetCustomerMainId(c)
    
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    if advertise_id != "" {
        whereParam["advertise_id"] = advertise_id
    }  
    if fec_source != "" {
        whereParam["fec_source"] = fec_source
    }  
    if fec_medium != "" {
        whereParam["fec_medium"] = fec_medium
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
    var advertise Advertise
    counts, err := engine.Where(whereStr, whereVal...).Count(&advertise)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var advertises []Advertise
    err = query.Find(&advertises) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    ownNameOps, err := customer.Get3OwnNameOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    var designGroupArr  []helper.VueSelectOps
    var contentGroupArr []helper.VueSelectOps
    var marketGroupArr  []helper.VueSelectOps
    
    for i:=0; i<len(advertises); i++ {
        advertise := advertises[i]
        fecDesignInt64, _ := helper.Int64(advertise.FecDesign)
        designGroupVal := initialization.CustomerIdWithName[fecDesignInt64]
        designGroupArr = append(designGroupArr, helper.VueSelectOps{Key: fecDesignInt64, DisplayName: designGroupVal})
        
        fecContent64, _ := helper.Int64(advertise.FecContent)
        contentGroupVal := initialization.CustomerIdWithName[fecContent64]
        contentGroupArr = append(contentGroupArr, helper.VueSelectOps{Key: fecContent64, DisplayName: contentGroupVal})
        
        marketGroup := advertise.MarketGroup
        marketGroupVal := initialization.MarketGroupIdWithName[advertise.MarketGroup]
        marketGroupArr = append(marketGroupArr, helper.VueSelectOps{Key: marketGroup, DisplayName: marketGroupVal})
        
        // helper.VueSelectOps{Key: channelInfo.Channel, DisplayName: channelInfo.Channel}
    }
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": advertises,
        "total": counts,
        "ownNameOps": ownNameOps,
        "designGroupOps": designGroupArr,
        "contentGroupOps": contentGroupArr,
        "marketGroupOps": marketGroupArr,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

func AdvertiseDownloadMutilXlsx(c *gin.Context){
    c.File("/root/go/src/github.com/fecshopsoft/fec-go/marketurl.xlsx")

}

func GenerateMutilAdvertise(c *gin.Context){
    //得到上传的文件
    file, err := c.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
    fileName := "/" + helper.RandomUUID() + ".xlsx"
    outfileName := "/out_" + helper.RandomUUID() + ".xlsx"
    saveUploadFileDir := config.Get("saveUploadFileDir")
    saveUploadFileName := saveUploadFileDir + fileName
    outUploadFileName := saveUploadFileDir + outfileName
    log.Println(saveUploadFileName)
    err = c.SaveUploadedFile(file, saveUploadFileName)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 用excel打开 文件 saveUploadFileName
    
    xlFile, err := xlsx.OpenFile(saveUploadFileName)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    
    var file2 *xlsx.File
    var sheet2 *xlsx.Sheet
    var row2 *xlsx.Row
    var cell2 *xlsx.Cell

    file2 = xlsx.NewFile()
    sheet2, err = file2.AddSheet("Sheet1")
    if err != nil {
        
    }
    // 增加标题
    row2 = sheet2.AddRow()
    cell2 = row2.AddCell()
    cell2.Value = "origin url（初始URL）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_source（渠道）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_medium（子渠道）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_campaign（活动）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_design（美工）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_cost（广告费）"
    cell2 = row2.AddCell()
    cell2.Value = "remark（广告备注）"
    
    cell2 = row2.AddCell()
    cell2.Value = "error_info（报错信息）"
    cell2 = row2.AddCell()
    cell2.Value = "warn_info（警告信息，可忽略）"
    cell2 = row2.AddCell()
    cell2.Value = "fid（广告唯一标示）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_person（广告员工）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_market_group（广告小组）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_url（广告url，您的广告请使用下面的url）"
    //
    errStr := ""
    currentCustomerId := helper.GetCurrentCustomerId(c)
    own_id := customer.GetCustomerMainId(c)
    customerOne, err := customer.GetCustomerOneById(currentCustomerId)
    if err != nil {
        errStr += err.Error()
    }
    marketGroupId := customerOne.MarketGroupId
    
    for _, sheet := range xlFile.Sheets {
        ii := 0
        for _, row := range sheet.Rows {
            ii = ii + 1
            if ii > 1 {
                rowCells    := row.Cells
                rowLen := len(rowCells)
                url := ""
                if rowLen > 0 {
                    url = rowCells[0].String()
                }
                source := ""
                if rowLen > 1 {
                    source = rowCells[1].String()
                }
                medium := ""
                if rowLen > 2 {
                    medium = rowCells[2].String()
                }
                campaign := ""
                if rowLen > 3 {
                    campaign = rowCells[3].String()
                }
                design := ""
                if rowLen > 4 {
                    design = rowCells[4].String()
                }
                //content     := rowCells[5].String()
                cost := ""
                if rowLen > 5 {
                    cost = rowCells[5].String()
                }
                remark := ""
                if rowLen > 6 {
                    remark = rowCells[6].String()
                }
                
                log.Println("url:" + url)
                log.Println("source:" + source)
                log.Println("medium:" + medium)
                log.Println("campaign:" + campaign)
                log.Println("design:" + design)
                // log.Println("content:" + content)
                log.Println("cost:" + cost)
                log.Println("remark:" + remark)
                // 加入原始信息
                row2 = sheet2.AddRow()
                cell2 = row2.AddCell()
                cell2.Value = url
                cell2 = row2.AddCell()
                cell2.Value = source
                cell2 = row2.AddCell()
                cell2.Value = medium
                cell2 = row2.AddCell()
                cell2.Value = campaign
                cell2 = row2.AddCell()
                cell2.Value = design
                cell2 = row2.AddCell()
                cell2.Value = cost
                cell2 = row2.AddCell()
                cell2.Value = remark
                // 计算信息
                errStr, warnStr, successInfo, _ := generateAdvertiseInfo(source, medium, campaign, design, url,  cost, remark, currentCustomerId, own_id, marketGroupId, errStr, "")
                cell2 = row2.AddCell()
                cell2.Value = errStr
                cell2 = row2.AddCell()
                cell2.Value = warnStr
                if errStr == "" && successInfo != nil {
                    
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["fid"].(string)
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["advertise_person"].(string)
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["advertise_market_group"].(string)
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["advertise_url"].(string)
                } else {
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                }
            }  
        }
    }
    
    err = file2.Save(outUploadFileName)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    
    c.File(outUploadFileName)

}




func GenerateMutilLinkAdvertise(c *gin.Context){
    //得到上传的文件
    file, err := c.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
    fileName := "/" + helper.RandomUUID() + ".xlsx"
    outfileName := "/out_" + helper.RandomUUID() + ".xlsx"
    saveUploadFileDir := config.Get("saveUploadFileDir")
    saveUploadFileName := saveUploadFileDir + fileName
    outUploadFileName := saveUploadFileDir + outfileName
    log.Println(saveUploadFileName)
    err = c.SaveUploadedFile(file, saveUploadFileName)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 用excel打开 文件 saveUploadFileName
    
    xlFile, err := xlsx.OpenFile(saveUploadFileName)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    
    var file2 *xlsx.File
    var sheet2 *xlsx.Sheet
    var row2 *xlsx.Row
    var cell2 *xlsx.Cell

    file2 = xlsx.NewFile()
    sheet2, err = file2.AddSheet("Sheet1")
    if err != nil {
        
    }
    // 增加标题
    row2 = sheet2.AddRow()
    cell2 = row2.AddCell()
    cell2.Value = "origin url（初始URL）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_source（渠道）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_medium（子渠道）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_campaign（活动）"
    cell2 = row2.AddCell()
    cell2.Value = "fec_design（美工）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_cost（广告费）"
    cell2 = row2.AddCell()
    cell2.Value = "remark（广告备注）"
    
    cell2 = row2.AddCell()
    cell2.Value = "error_info（报错信息）"
    cell2 = row2.AddCell()
    cell2.Value = "warn_info（警告信息，可忽略）"
    cell2 = row2.AddCell()
    cell2.Value = "fid（广告唯一标示）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_person（广告员工）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_market_group（广告小组）"
    cell2 = row2.AddCell()
    cell2.Value = "advertise_url（广告url，您的广告请使用下面的url）"
    //
    errStr := ""
    currentCustomerId := helper.GetCurrentCustomerId(c)
    own_id := customer.GetCustomerMainId(c)
    customerOne, err := customer.GetCustomerOneById(currentCustomerId)
    if err != nil {
        errStr += err.Error()
    }
    marketGroupId := customerOne.MarketGroupId
    fid := ""
    for _, sheet := range xlFile.Sheets {
        ii := 0
        for _, row := range sheet.Rows {
            ii = ii + 1
            if ii > 1 {
                rowCells    := row.Cells
                rowLen := len(rowCells)
                url := ""
                if rowLen > 0 {
                    url = rowCells[0].String()
                }
                source := ""
                if rowLen > 1 {
                    source = rowCells[1].String()
                }
                medium := ""
                if rowLen > 2 {
                    medium = rowCells[2].String()
                }
                campaign := ""
                if rowLen > 3 {
                    campaign = rowCells[3].String()
                }
                design := ""
                if rowLen > 4 {
                    design = rowCells[4].String()
                }
                //content     := rowCells[5].String()
                cost := ""
                if rowLen > 5 {
                    cost = rowCells[5].String()
                }
                remark := ""
                if rowLen > 6 {
                    remark = rowCells[6].String()
                }
                
                log.Println("url:" + url)
                log.Println("source:" + source)
                log.Println("medium:" + medium)
                log.Println("campaign:" + campaign)
                log.Println("design:" + design)
                // log.Println("content:" + content)
                log.Println("cost:" + cost)
                log.Println("remark:" + remark)
                // 加入原始信息
                row2 = sheet2.AddRow()
                cell2 = row2.AddCell()
                cell2.Value = url
                cell2 = row2.AddCell()
                cell2.Value = source
                cell2 = row2.AddCell()
                cell2.Value = medium
                cell2 = row2.AddCell()
                cell2.Value = campaign
                cell2 = row2.AddCell()
                cell2.Value = design
                cell2 = row2.AddCell()
                cell2.Value = cost
                cell2 = row2.AddCell()
                cell2.Value = remark
                // 计算信息
                var warnStr string
                var successInfo gin.H
                
                errStr, warnStr, successInfo, fid = generateAdvertiseInfo(source, medium, campaign, design, url,  cost, remark, currentCustomerId, own_id, marketGroupId, errStr, fid)
                cell2 = row2.AddCell()
                cell2.Value = errStr
                cell2 = row2.AddCell()
                cell2.Value = warnStr
                if errStr == "" && successInfo != nil {
                    
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["fid"].(string)
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["advertise_person"].(string)
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["advertise_market_group"].(string)
                    cell2 = row2.AddCell()
                    cell2.Value, _ = successInfo["advertise_url"].(string)
                } else {
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                    cell2 = row2.AddCell()
                    cell2.Value = ""
                }
            }  
        }
    }
    
    err = file2.Save(outUploadFileName)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    
    c.File(outUploadFileName)

}

/*
 return gin.H{
        "error": errStr,
        "warning": warnStr,
        "success_info": successInfo,
    }

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

*/
