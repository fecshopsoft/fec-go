package common

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "errors"
    // "time"
    "unicode/utf8"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mysqldb"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/handler/customer"
    //"fmt"
)

type WebsiteInfo struct {
    Id int64 `form:"id" json:"id"`
    SiteName string `form:"site_name" json:"site_name" binding:"required"`
    Domain string `form:"domain" json:"domain" binding:"required"`
    TraceJsUrl string `form:"trace_js_url" json:"trace_js_url" `
    TraceApiUrl string `form:"trace_api_url" json:"trace_api_url" `
    SiteUid string `form:"site_uid" json:"site_uid"`
    AccessToken string `form:"access_token" json:"access_token"`
    OwnId int64 `form:"own_id" json:"own_id"`
    Status int64 `form:"status" json:"status" binding:"required"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
    
    PaymentEndTime int64 `xorm:"payment_end_time" form:"payment_end_time" json:"payment_end_time"`
    WebsiteDayMaxCount int64 `xorm:"website_day_max_count" form:"website_day_max_count" json:"website_day_max_count"`
    
}

func (websiteInfo WebsiteInfo) TableName() string {
    return "base_website_info"
}
var enableStatus int = 1
var FecTraceJsUrl string = "trace.fecshop.com/fec_trace.js"
var FecTraceApiUrl string = "120.24.37.249:3000/fec/trace/api"
/**
 * 增加一条记录
 */
func WebsiteAddOne(c *gin.Context){
    var websiteInfo WebsiteInfo
    err := c.ShouldBindJSON(&websiteInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    domain := websiteInfo.Domain
    if helper.IsValidDomain(domain) == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("domain [right format is: www.fecshop.com] is not a domain format, "))
        return
    }
    // 处理own_id
    own_id, err := customer.Get3SaveDataOwnId(c, websiteInfo.OwnId)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 查看创建site是否达到最大数
    sites, err := GetWebsiteByOwnId(own_id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    ownIdSiteCount := len(sites)
    ownCustomer, err := customer.GetCustomerOneById(own_id)
    // 如果允许的新建site的最大数 <= 当前的site数，
    if ownCustomer.WebsiteCount <= int64(ownIdSiteCount) {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("max site count limit"))
        return
    }
    
    websiteInfo.SiteUid = helper.RandomUUID()
    // access_token, err := helper.GenerateAccessToken()
    access_token, err := helper.GenerateAccessTokenBySiteId(websiteInfo.SiteUid)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    websiteInfo.AccessToken = access_token
    websiteInfo.OwnId = own_id
    websiteInfo.TraceJsUrl = FecTraceJsUrl
    websiteInfo.TraceApiUrl = FecTraceApiUrl
    customerId := helper.GetCurrentCustomerId(c)
    websiteInfo.CreatedCustomerId = customerId
    // 处理   PaymentEndTime WebsiteDayMaxCount, 如果不是超级用户，无权修改这个字段
    if helper.IsSuperAdmin(c) == false {
        websiteInfo.PaymentEndTime = 0
        websiteInfo.WebsiteDayMaxCount = 0
    }
    // 插入
    affected, err := engine.Insert(&websiteInfo)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "websiteInfo":websiteInfo,
    })
    c.JSON(http.StatusOK, result)
}



/**
 * 通过id为条件，更新一条记录
 */
func WebsiteUpdateById(c *gin.Context){
    var websiteInfo WebsiteInfo
    err := c.ShouldBindJSON(&websiteInfo);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    domain := websiteInfo.Domain
    if helper.IsValidDomain(domain) == false {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("domain [right format is: www.fecshop.com] is not a domain format, "))
        return
    }
    // 处理own_id
    own_id, err := customer.Get3SaveDataOwnId(c, websiteInfo.OwnId)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    websiteInfo.OwnId = own_id
    cols := "site_name,domain,trace_js_url,status,own_id,updated_at"
    // 处理   PaymentEndTime WebsiteDayMaxCount, 如果是超级用户，才可以修改这个字段
    if helper.IsSuperAdmin(c) == true {
        cols += ",payment_end_time,website_day_max_count"
    }
    // 更新
    affected, err := engine.Where("id = ?",websiteInfo.Id).Cols(cols).Update(&websiteInfo)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    } 
    result := util.BuildSuccessResult(gin.H{
        "affected":affected,
        "websiteInfo":websiteInfo,
    })
    c.JSON(http.StatusOK, result)
}
/**
 * 删除一条记录
 */
func WebsiteDeleteById(c *gin.Context){
    var websiteInfo WebsiteInfo
    var id helper.DeleteId
    err := c.ShouldBindJSON(&id);
    // customerId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    affected, err := engine.Where("id = ?",id.Id).Delete(&websiteInfo)
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
func WebsiteDeleteByIds(c *gin.Context){
    engine := mysqldb.GetEngine()
    var websiteInfo WebsiteInfo
    var ids helper.DeleteIds
    err := c.ShouldBindJSON(&ids);
    //c.JSON(http.StatusOK, ids)
    affected, err := engine.In("id", ids.Ids).Delete(&websiteInfo)
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
func WebsiteList(c *gin.Context){
    // params := c.Request.URL.Query()
    // 获取参数并处理
    var sortD string
    var sortColumns string
    var own_id int64
    defaultPageNum:= c.GetString("defaultPageNum")
    defaultPageCount := c.GetString("defaultPageCount")
    page, _  := strconv.Atoi(c.DefaultQuery("page", defaultPageNum))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", defaultPageCount))
    site_name     := c.DefaultQuery("site_name", "")
    own_id_i, _ := strconv.Atoi(c.DefaultQuery("own_id", ""))
    own_id = int64(own_id_i)
    status_i, _ := strconv.Atoi(c.DefaultQuery("status", ""))
    status := int64(status_i)
    sort     := c.DefaultQuery("sort", "")
    created_at_begin := c.DefaultQuery("created_begin_timestamps", "")
    created_at_end   := c.DefaultQuery("created_end_timestamps", "")
    if utf8.RuneCountInString(sort) >= 2 {
        sortD = string([]byte(sort)[:1])
        sortColumns = string([]byte(sort)[1:])
    } 
    whereParam := make(mysqldb.XOrmWhereParam)
    if site_name != "" {
        whereParam["site_name"] = []string{"like", site_name}
    }  
    own_id, err := customer.Get3OwnId(c, own_id)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    if status != 0 {
        whereParam["status"] = status
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
    var websiteInfo WebsiteInfo
    counts, err := engine.Where(whereStr, whereVal...).Count(&websiteInfo)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    // 得到结果数据
    var websiteInfos []WebsiteInfo
    err = query.Find(&websiteInfos) 
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    ownNameOps, err := customer.Get3OwnNameOps(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    createdCustomerOps, err := GetWebsiteCreatedCustomerOps(websiteInfos)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    customerType := helper.GetCurrentCustomerType(c)
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "items": websiteInfos,
        "total": counts,
        "createdCustomerOps": createdCustomerOps,
        "ownNameOps": ownNameOps,
        "customerType": customerType,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}




// 通过创建人ids，得到创建人的name
func GetWebsiteCreatedCustomerOps(websiteInfos []WebsiteInfo) ([]helper.VueSelectOps, error){
    var groupArr []helper.VueSelectOps
    var ids []int64
    for i:=0; i<len(websiteInfos); i++ {
        websiteInfo := websiteInfos[i]
        ids = append(ids, websiteInfo.CreatedCustomerId)
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
 * 根据 market_group_ids 查询得到 WebsiteInfo
 */
func GetWebsiteByIds(market_group_ids []int64) ([]WebsiteInfo, error){
    // 得到结果数据
    var websiteInfos []WebsiteInfo
    err := engine.In("id", market_group_ids).Find(&websiteInfos) 
    if err != nil{
        return websiteInfos, err
    }
    return websiteInfos, nil
}

/**
 * 根据 market_group_ids 查询得到 WebsiteInfo
 */
func GetWebsiteByOwnId(own_id int64) ([]WebsiteInfo, error){
    // 得到结果数据
    var websiteInfos []WebsiteInfo
    err := engine.Where("own_id = ? ", own_id).Find(&websiteInfos) 
    if err != nil{
        return websiteInfos, err
    }
    return websiteInfos, nil
}

/**
 * 根据 market_group_ids 查询得到 WebsiteInfo
 */
func GetActiveWebsiteByOwnId(own_id int64) ([]WebsiteInfo, error){
    // 得到结果数据
    var websiteInfos []WebsiteInfo
    err := engine.Where("own_id = ? and status = ? ", own_id, enableStatus).Find(&websiteInfos) 
    if err != nil{
        return websiteInfos, err
    }
    return websiteInfos, nil
}


/**
 * 根据 market_group_ids 查询得到 WebsiteInfo
 */
func GetAllActiveWebsiteId() ([]WebsiteInfo, error){
    var websiteInfos []WebsiteInfo
    activeCustomers, err := customer.GetAllEnableCommonCustomer()
    if err != nil{
        return websiteInfos, err
    }
    var customerIds []int64
    for i:=0; i<len(activeCustomers); i++ {
        activeCustomer := activeCustomers[i]
        customerId := activeCustomer.Id
        customerIds = append(customerIds,customerId)
    }
    
    
    err = engine.In("own_id", customerIds).Where("status = ? ", enableStatus).Find(&websiteInfos) 
    if err != nil{
        return websiteInfos, err
    }
    return websiteInfos, nil
    
}

/**
 * 通过id查询一条记录
 */
func GetCustomerOneById(id int64) (WebsiteInfo, error){
    var websiteInfo WebsiteInfo
    has, err := engine.Where("id = ?", id).Get(&websiteInfo)
    if err != nil {
        return websiteInfo, err
    } 
    if has == false {
        return websiteInfo, errors.New("get websiteInfo by id error, empty data.")
    }
    return websiteInfo, nil
}

/**
 * 根据 access_token 查询得到 WebsiteInfo
 */
func GetWebsiteByAccessToken(access_token string) (WebsiteInfo, error){
    // 得到结果数据
    var websiteInfo WebsiteInfo
    has, err := engine.Where("access_token = ?", access_token).Get(&websiteInfo)
    if err != nil {
        return websiteInfo, err
    } 
    if has == false {
        return websiteInfo, errors.New("get websiteInfo by id error, empty data.")
    }
    return websiteInfo, nil
}


/**
 * 列表查询
 */
func WebsiteJsCode(c *gin.Context){
    site_id_str := c.DefaultQuery("site_id", "")
    site_id_i, _ := strconv.Atoi(site_id_str)
    site_id := int64(site_id_i)
    if site_id == 0{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("you should pass param: site_id"))
        return  
    }
    websiteInfo, err := GetCustomerOneById(site_id)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    trace_js_url := websiteInfo.TraceJsUrl // trace_js_url := "static.xxx-cdn.com/xxx/js/trace.js"
    site_uid := websiteInfo.SiteUid
    if trace_js_url == ""{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("trace_js_url is null, you must update trace_js_url info "))
        return  
    }
    if site_uid == ""{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("site_uid is null, error "))
        return  
    }
    common_js := ` // 将该js代码添加的所有网站页面，您可以添加到网站的底部
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['website_id', '` + site_uid + `']);
    _maq.push(['fec_store', store_name]);
    _maq.push(['fec_lang', store_language]);
    _maq.push(['fec_app', store_app_name]);
    _maq.push(['fec_currency', store_currency]);
    
	(function() {
		var ma = document.createElement('script'); ma.type = 'text/javascript'; ma.async = true;
		ma.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + '` + trace_js_url + `';
		var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ma, s);
	})();
</script>`
    
    category_js := ` // 将该js代码添加到分类页面
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['category', category]);  // category：填写分类的name，如果是多语言网站，那么这里填写默认语言的分类name
</script>`

    product_js := ` // 将该js添加到产品页面
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['sku', sku]); // sku：当前产品页面的sku编码
</script>`
    
    cart_js := ` // 将该js添加到购物车页面，
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['cart', cart]); // cart：购物车页面的购物车数据，该值的示例数据请看下面
</script>`

    example_cart_js := ` // 上面 cart 变量的示例格式数据。
[
    {
        "sku":"grxjy56002622",
        "qty":1,
        "price":35.52
    },
    {
        "sku":"grxjy5606622",
        "qty":4,
        "price":75.11
    }
]`

    search_js := ` // 将该js添加到搜索页面
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['search', search_text]);
</script>`
    example_search_js := ` search_text 变量的示例数据
{
    "text": "fashion handbag", // 搜索词
    "result_qty":5  // 搜索的产品个数
}`
    
    
    
    
    
    
    

    login_js := ` // 登录成功页面添加的js，如果是ajax登录，那么需要通过API，服务端发送用户登录email
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['login_email', $login_email]);
</script>`
    
    register_js := ` // 注册成功页面添加的js，如果是ajax登录，那么需要通过API，服务端发送用户注册email
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['register_email', $register_email]);
</script>`

    order_js := ` // 如果您的订单生成后，需要跳转到支付平台，如果有跳转桥页，可以在该桥页添加下面的js，如果没有，直接跳转，那么需要在服务端调用api发送数据。
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['order', order]);
</script>`
    example_order_js := ` // 上面的 order 变量的示例数据
{
    "invoice": "500023149", // 订单号
    "order_type": "standard or express", // standard（标准支付流程类型）express（基于api的支付类型，譬如paypal快捷支付。）
    "payment_status":"pending", // pending（未支付成功）
    "payment_type":"paypal", // 支付渠道，譬如是paypal还是西联等支付渠道
    "currency":"RMB", // 当前货币
    "currency_rate":6.2, // 公式：当前金额 * 汇率 = 美元金额
    "amount":35.52, // 订单总金额
    "shipping":0.00, // 运费金额
    "discount_amount":0.00, // 折扣金额
    "coupon":"xxxxx", // 优惠券，没有则为空
    "city":"fdasfds", // 城市
    "email":"2358269014@qq.com", // 下单填写的email
    "first_name":"terry", //
    "last_name":"water", //
    "zip":"266326", // 邮编
    "country_code":"US", // 国家简码
    "state_code":"CT", // 省或州简码
    "country_name":"Unite States", // 国家全称
    "state_name":"Ctsrse", // 省或州全称
    "address1":"address street 1", // 详细地址1
    "address2":"address street 2", // 详细地址2
    "products":[ // 产品详情
        {
            "sku":"xxxxyr", // sku
            "name":"Fashion Solid Color Warm Coat", // 产品名称
            "qty":1, // 个数
            "price":25.92 // 产品单价
        },
        {
            "sku":"yyyy", // sku
            "name":"Fashion Waist Warm Coat", // 产品名称
            "qty":1, // 个数
            "price":34.16 // 产品单价
        }
    ]
}`
    success_order_js := ` // 将该js，添加到订单支付成功页面
<script type="text/javascript">
	var _maq = _maq || [];
	_maq.push(['order', $order]);
</script>`
    
    example_success_order_js := ` // 上面 $order 变量的示例数据
{
    "invoice": "500023149", // 订单号
    "order_type": "standard or express", // standard（标准支付流程类型）express（基于api的支付类型，譬如paypal快捷支付。）
    "payment_status":"pending", // pending（未支付成功）
    "payment_type":"paypal", // 支付渠道，譬如是paypal还是西联等支付渠道
    "currency":"RMB", // 当前货币
    "currency_rate":6.2, // 公式：当前金额 * 汇率 = 美元金额
    "amount":35.52, // 订单总金额
    "shipping":0.00, // 运费金额
    "discount_amount":0.00, // 折扣金额
    "coupon":"xxxxx", // 优惠券，没有则为空
    "city":"fdasfds", // 城市
    "email":"2358269014@qq.com", // 下单填写的email
    "first_name":"terry", //
    "last_name":"water", //
    "zip":"266326", // 邮编
    "country_code":"US", // 国家简码
    "state_code":"CT", // 省或州简码
    "country_name":"Unite States", // 国家全称
    "state_name":"Ctsrse", // 省或州全称
    "address1":"address street 1", // 详细地址1
    "address2":"address street 2", // 详细地址2
    "products":[ // 产品详情
        {
            "sku":"xxxxyr", // sku
            "name":"Fashion Solid Color Warm Coat", // 产品名称
            "qty":1, // 个数
            "price":25.92 // 产品单价
        },
        {
            "sku":"yyyy", // sku
            "name":"Fashion Waist Warm Coat", // 产品名称
            "qty":1, // 个数
            "price":34.16 // 产品单价
        }
    ]
}`
    
    // 生成返回结果
    result := util.BuildSuccessResult(gin.H{
        "common_js": common_js,
        "category_js": category_js,
        "product_js": product_js,
        "cart_js": cart_js,
        "example_cart_js": example_cart_js,
        "order_js": order_js,
        "example_order_js": example_order_js,
        "success_order_js": success_order_js,
        "example_success_order_js": example_success_order_js,
        "login_js": login_js,
        "register_js": register_js,
        "search_js": search_js,
        "example_search_js": example_search_js,
    })
    // 返回json
    c.JSON(http.StatusOK, result)
}

