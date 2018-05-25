package customer

import(
    // "github.com/fecshopsoft/fec-go/shell/whole"
    // "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/fecshopsoft/fec-go/db/esdb"
    "github.com/fecshopsoft/fec-go/helper"
    "github.com/fecshopsoft/fec-go/shell/model"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "math"
    "log"
    
)

// uuid 统计计算部分
func UuidMapReduct(dbName string, collName string, outCollName string, website_id string) error {
    var err error
    mapStr := `
        function() {  
            website_id = "` + website_id + `";
            customer_id = this.customer_id ? this.customer_id : null;
            uuid = this.uuid ? this.uuid : null;
            register_email = this.register_email ? this.register_email : null;
            login_email = this.login_email ? this.login_email : null;
            
            service_date_str = this.service_date_str ? this.service_date_str : null;
            stay_seconds= this.stay_seconds ? this.stay_seconds : 0;
            sku 		= this.sku ? [this.sku] : null;
            category 	= this.category ? [this.category] : null;
            
            search 		= this.search ? [this.search] : null;
            cart		= this.cart ? this.cart : null;  // last cart page
            order		= this.order ? [this.order] : null;
            customer_email		= [];
            if(login_email){
                customer_email.push(login_email);
            }
            if(register_email){
                customer_email.push(register_email);
            }
            
            // user_agent= this.user_agent ? this.user_agent : null;
            // service_timestamp= this.service_timestamp ? this.service_timestamp : null;
           
            country_code= this.country_code ? this.country_code : null;
            ip= this.ip ? this.ip : null;
             devide= this.devide ? this.devide : null;
            
            browser_name= this.browser_name ? this.browser_name : null;
            browser_lang= this.browser_lang ? this.browser_lang : null;
            browser_version= this.browser_version ? this.browser_version : null;
            operate= this.operate ? this.operate : null;
            
            
            
            first_referrer_domain= this.first_referrer_domain ? this.first_referrer_domain : null;
            
            is_return= this.is_return ? this.is_return : 0;
            
            first_page_url= (this.first_page == 1 ) ? this.url : null;
            
            refer_url = null;
            if(this.first_page == 1){
                refer_url= this.refer_url ? this.refer_url : null;
            }
            if(!stay_seconds){
                out_page = this.url;
            }else{
                out_page = null;
            }
            visit_page_order_processing = 2;
            visit_page_order_pending = 2;
            visit_page_order_processing_amount = 0;
            visit_page_order_pending_amount = 0;
            visit_page_order_amount = 0;
            if(v_order = this.order){
                if(v_order['amount']){
                    visit_page_order_amount = v_order['amount'];
                }
                if(v_order['payment_status'] == 'payment_confirmed'){
                    visit_page_order_processing = 1;
                    if(v_order['amount']){
                        visit_page_order_processing_amount = v_order['amount'];
                    }
                }else{
                    visit_page_order_pending = 1;
                    if(v_order['amount']){
                        visit_page_order_pending_amount = v_order['amount'];
                    }
                }
                
                if(v_order['email']){
                    customer_email.push(v_order['email']);
                }
                
            }
            
            sku_cart = [];
            sku_order = [];
            sku_order_success = [];
            thiscart = this.cart;
            if(thiscart){
                for(x in thiscart){
                    c_one = thiscart[x];
                    $c_sku = c_one['sku'];
                    if($c_sku){
                        sku_cart.push($c_sku);
                    }
                }
            }
            thisorder = this.order;
            if(thisorder){ 
                c_products = thisorder.products;
                o_payment_status = thisorder.payment_status;
                for(x in c_products){
                    o_product = c_products[x];
                    o_sku = o_product['sku'];
                    if(o_sku){
                        if(o_payment_status == 'payment_confirmed'){
                            sku_order_success.push(o_sku);
                        }
                        sku_order.push(o_sku);
                    }
                }
            }
            
            
            
            if(customer_id){
                emit(this.service_date_str+"_"+this.customer_id,{
                    uuid:uuid,
                    customer_id:customer_id,
                    pv:1,
                    stay_seconds:stay_seconds,
                    register_email:register_email,
                    login_email:login_email,
                    service_date_str:service_date_str,
                    customer_email:customer_email,
                    fid:this.fid ? this.fid : null,
                        
                    fec_content:this.fec_content ? this.fec_content : null,
                    fec_market_group:this.fec_market_group ? this.fec_market_group : null,
                    fec_campaign:this.fec_campaign ? this.fec_campaign : null,
                    fec_source:this.fec_source ? this.fec_source : null,
                    fec_medium:this.fec_medium ? this.fec_medium : null,
                    fec_design:this.fec_design ? this.fec_design : null,
                    
                        
                    sku:sku,
                    sku_cart:sku_cart,
                    sku_order:sku_order,
                    sku_order_success:sku_order_success,
            
                    category:category,
                    search:search,
                    cart:cart,
                    order:order,
                    
                    visit_page_sku:this.sku ? 1 : 2,
                    visit_page_category:this.category ? 1 : 2,
                    visit_page_search:this.search ? 1 : 2,
                    visit_page_cart:this.cart ? 1 : 2,
                    visit_page_order:this.order ? 1 : 2,
                    visit_page_order_amount:visit_page_order_amount,
                    visit_page_order_processing:visit_page_order_processing,
                    visit_page_order_processing_amount:visit_page_order_processing_amount,
                    visit_page_order_pending:visit_page_order_pending,
                    visit_page_order_pending_amount:visit_page_order_pending_amount,
                    
                    domain:this.domain ? this.domain : null,
                    country_code:country_code,
                    country_name:this.country_name ? this.country_name : null,
                    ip:ip,
                    
                    devide:devide,
                    //user_agent:user_agent,
                    browser_name:browser_name,
                    browser_version:browser_version,
                    browser_lang:browser_lang,
                    operate:operate,
                    refer_url:refer_url,
                    first_referrer_domain:first_referrer_domain,
                    
                    is_return:is_return,
                    
                    color_depth:this.color_depth ? this.color_depth : null,
                    resolution:this.resolution ? this.resolution : null,
                    first_page_url:first_page_url,
                    out_page:out_page,
                    device_pixel_ratio:this.device_pixel_ratio ? this.device_pixel_ratio : null,
                    service_date_str:service_date_str,
                    
                    
                    data:[{
                        _id:this._id ? this._id : null,
                        ip:this.ip ? this.ip : null,
                        country_code:this.country_code ? this.country_code : null,
                        country_name:this.country_name ? this.country_name : null,
                        service_datetime:this.service_datetime ? this.service_datetime : null,
                        //service_timestamp:this.service_timestamp ? this.service_timestamp : null,
                        devide:this.devide ? this.devide : null,
                        uuid:this.uuid ? this.uuid : null,
                        fid:this.fid ? this.fid : null,
                        fec_content:this.fec_content ? this.fec_content : null,
                        fec_market_group:this.fec_market_group ? this.fec_market_group : null,
                        fec_campaign:this.fec_campaign ? this.fec_campaign : null,
                        fec_source:this.fec_source ? this.fec_source : null,
                        fec_medium:this.fec_medium ? this.fec_medium : null,
                        fec_design:this.fec_design ? this.fec_design : null,
                        
                        device_pixel_ratio:this.device_pixel_ratio ? this.device_pixel_ratio : null,
                        is_return:is_return,
                        user_agent:this.user_agent ? this.user_agent : null,
                        browser_name:this.browser_name ? this.browser_name : null,
                        browser_version:this.browser_version ? this.browser_version : null,
                        browser_date:this.browser_date ? this.browser_date : null,
                        browser_lang:this.browser_lang ? this.browser_lang : null,
                        operate:this.operate ?  this.operate : null,
                        operate_relase:this.operate_relase ? this.operate_relase : null,
                        domain:this.domain ? this.domain : null,
                        url:this.url ? this.url : null,
                        title:this.title ? this.title : null,
                        refer_url:this.refer_url ? this.refer_url  : null,
                        first_referrer_domain:this.first_referrer_domain ? this.first_referrer_domain : null,
                        resolution:this.resolution ? this.resolution : null,
                        color_depth:this.color_depth ? this.color_depth : null,
                        
                        first_page:this.first_page ? this.first_page : null,
                        url_new:this.url_new ? this.url_new : null,
                        login_email:this.login_email ? this.login_email : null ,
                        register_email:this.register_email ? this.register_email : null ,
                        sku:this.sku ? this.sku.replace(" ","") : null,
                        category:this.category ? this.category: null,
                        search:this.search ? this.search : null ,
                        cart:this.cart ? this.cart : null ,
                        stay_seconds : this.stay_seconds ? this.stay_seconds : null,
                        order:this.order ? this.order : null 
                        
                    }]
                    
                });
            }
            
        }
    `
    
    reduceStr := `
        function(key,emits){
            this_uuid 				= null;
            this_mid 				= null;
            
            this_customer_email		= [];
            this_fec_content 				= null;
            this_fec_market_group 			= null;
            this_fec_campaign 				= null;
            this_fec_source 				= null;
            this_fec_medium 				= null;
            this_fec_design 				= null;
                        
            this_customer_id		= null;
            this_identify 			= [];
            this_pv 				= 0;
            this_stay_seconds		= 0;
            this_service_date_str 	= null;
            this_register_email		= null;
            
            this_login_email		= null;
            
            this_sku				= [];
            this_sku_cart			= [];
            this_sku_order			= [];
            this_sku_order_success	= [];
            
            this_category		= [];
            this_search		= [];
            this_cart		= [];
            this_order		= [];
            
            
            this_visit_page_sku			= 2;
            this_visit_page_category	= 2;
            this_visit_page_search		= 2;
            this_visit_page_cart		= 2;
            this_visit_page_order		= 2;
            
            this_visit_page_order_processing		= 2;
            this_visit_page_order_pending			= 2;
            this_visit_page_order_amount			= 0;
            this_visit_page_order_processing_amount	= 0;
            this_visit_page_order_pending_amount	= 0;
                
            
            this_ip		= null;
            
            
            this_devide	= null;
            this_browser_name	= null;
            this_browser_lang	= null;
            this_browser_version= null;
            this_operate	= null;
            this_refer_url	= null;
            this_first_referrer_domain	= null;
            this_is_return	= null;
            this_first_page_url	= null;
            this_out_page   = null;
            this_country_code = null;
            this_country_name = null;
            this_data = [];
            this_color_depth = null;
            this_resolution  = null;
            this_device_pixel_ratio = null;
            for(var i in emits){
                
                if(emits[i].uuid){
                    this_uuid 			=  emits[i].uuid;
                }
                
                if(emits[i].mid){
                    this_mid 			=  emits[i].mid;
                }
                
                    
                if(emits[i].fec_content){
                    this_fec_content 			=  emits[i].fec_content;
                }
                if(emits[i].fec_market_group){
                    this_fec_market_group 			=  emits[i].fec_market_group;
                }
                if(emits[i].fec_campaign){
                    this_fec_campaign 			=  emits[i].fec_campaign;
                }
                if(emits[i].fec_source){
                    this_fec_source 			=  emits[i].fec_source;
                }
                if(emits[i].fec_medium){
                    this_fec_medium 			=  emits[i].fec_medium;
                }
                if(emits[i].fec_design){
                    this_fec_design 			=  emits[i].fec_design;
                }
                
            
                if(emits[i].customer_id){	
                    if(!this_customer_id){
                        this_customer_id	=  emits[i].customer_id;
                    }else if(this_customer_id > emits[i].customer_id){
                        this_customer_id	=  emits[i].customer_id;
                    }
                }
                if(emits[i].pv){	
                    this_pv 			+= emits[i].pv;
                }
                if(emits[i].stay_seconds){	
                    this_stay_seconds   += emits[i].stay_seconds;
                }
                if(emits[i].service_date_str){	
                    this_service_date_str = emits[i].service_date_str;
                }
                if(emits[i].register_email){
                    this_register_email	=  emits[i].register_email;
                }
                if(emits[i].login_email){
                    this_login_email	=  emits[i].login_email;
                }
                
                // this_sku
                emits_sku = emits[i].sku;
                if(emits_sku &&  (emits_sku.length > 0) ){
                    for(x in emits_sku){
                        e_sku = emits_sku[x];
                        if(this_sku.indexOf(e_sku) == -1){
                            this_sku.push(e_sku);
                        }
                    }
                }
                
                emits_customer_email = emits[i].customer_email;
                if(emits_customer_email &&  (emits_customer_email.length > 0) ){
                    for(x in emits_customer_email){
                        e_customer_email = emits_customer_email[x];
                        if(this_customer_email.indexOf(e_customer_email) == -1){
                            this_customer_email.push(e_customer_email);
                        }
                    }
                }
                
                
                
                // this_sku_cart			= [];
                
                emits_sku_cart = emits[i].sku_cart;
                if(emits_sku_cart &&  (emits_sku_cart.length > 0) ){
                    //this_sku			=  this_sku.concat(emits[i].sku);
                    for(x in emits_sku_cart){
                        e_sku = emits_sku_cart[x];
                        if(this_sku_cart.indexOf(e_sku) == -1){
                            this_sku_cart.push(e_sku);
                        }
                    }
                }
                
                // this_sku_order			= [];
                
                emits_sku_order = emits[i].sku_order;
                if(emits_sku_order &&  (emits_sku_order.length > 0) ){
                    //this_sku			=  this_sku.concat(emits[i].sku);
                    for(x in emits_sku_order){
                        e_sku = emits_sku_order[x];
                        if(this_sku_order.indexOf(e_sku) == -1){
                            this_sku_order.push(e_sku);
                        }
                    }
                }
                
                // this_sku_order_success	= [];
                emits_sku_order_success = emits[i].sku_order_success;
                if(emits_sku_order_success &&  (emits_sku_order_success.length > 0) ){
                    //this_sku			=  this_sku.concat(emits[i].sku);
                    for(x in emits_sku_order_success){
                        e_sku = emits_sku_order_success[x];
                        if(this_sku_order_success.indexOf(e_sku) == -1){
                            this_sku_order_success.push(e_sku);
                        }
                    }
                }
                
                
                
                
                if(emits[i].category){
                    this_category			=  this_category.concat(emits[i].category);
                }
                if(emits[i].search && emits[i].search.length > 0){
                    if(Object.prototype.toString.call( emits[i].search ) === '[object Array]'){
                        this_search			=  this_search.concat(emits[i].search);
                    }
                }
                if(emits[i].cart && emits[i].cart.length > 0){
                    if(Object.prototype.toString.call( emits[i].cart ) === '[object Array]'){
                        this_cart			=  this_cart.concat(emits[i].cart);
                    }
                }
                if(emits[i].order && emits[i].order.length > 0){
                    if(Object.prototype.toString.call( emits[i].order ) === '[object Array]'){
                        this_order			=  this_order.concat(emits[i].order);
                    }
                }
                
                if(emits[i].visit_page_sku  == 1){
                    this_visit_page_sku	=  emits[i].visit_page_sku;
                }
                if(emits[i].visit_page_category  == 1){
                    this_visit_page_category	=  emits[i].visit_page_category;
                }
                if(emits[i].visit_page_search  == 1){
                    this_visit_page_search	=  emits[i].visit_page_search;
                }
                if(emits[i].visit_page_cart  == 1){
                    this_visit_page_cart	=  emits[i].visit_page_cart;
                }
                
                if(emits[i].visit_page_order == 1){
                    this_visit_page_order	=  emits[i].visit_page_order;
                }
                
                if(emits[i].visit_page_order_processing == 1){
                    this_visit_page_order_processing	=  emits[i].visit_page_order_processing;
                }
                if(emits[i].visit_page_order_pending == 1){
                    this_visit_page_order_pending	=  emits[i].visit_page_order_pending;
                }
                
                if(emits[i].visit_page_order_amount){
                    this_visit_page_order_amount	+=  emits[i].visit_page_order_amount;
                }
                if(emits[i].visit_page_order_processing_amount){
                    this_visit_page_order_processing_amount	+=  emits[i].visit_page_order_processing_amount;
                }
                if(emits[i].visit_page_order_pending_amount){
                    this_visit_page_order_pending_amount	+=  emits[i].visit_page_order_pending_amount;
                }
                
                
                
                if(emits[i].ip){
                    this_ip				=  emits[i].ip;
                }
                if(emits[i].color_depth){
                    this_color_depth	=  emits[i].color_depth;
                }
                if(emits[i].resolution){
                    this_resolution		=  emits[i].resolution;
                }
                if(emits[i].device_pixel_ratio){
                    this_device_pixel_ratio =  emits[i].device_pixel_ratio;
                }
                if(emits[i].country_code){
                    this_country_code   =  emits[i].country_code;
                }
                if(emits[i].country_name){
                    this_country_name   =  emits[i].country_name;
                }
                
                if(emits[i].devide){	
                    this_devide			=  emits[i].devide;
                }
                if(emits[i].browser_name){	
                    this_browser_name	=  emits[i].browser_name;
                }
                if(emits[i].browser_lang){	
                    this_browser_lang	=  emits[i].browser_lang;
                }
                if(emits[i].browser_version){	
                    this_browser_version	=  emits[i].browser_version;
                }
                if(emits[i].operate){	
                    this_operate		=  emits[i].operate;
                }
                if(emits[i].refer_url){
                    this_refer_url		=  emits[i].refer_url;
                }
                
                if(emits[i].first_referrer_domain){
                    this_first_referrer_domain		=  emits[i].first_referrer_domain;
                }
                if(emits[i].is_return){
                    this_is_return		=  emits[i].is_return;
                }
                if(emits[i].first_page_url){
                    this_first_page_url		=  emits[i].first_page_url;
                }
                if(emits[i].out_page){
                    this_out_page		=  emits[i].out_page;
                }
                if(emits[i].data){
                    this_data = this_data.concat(emits[i].data);
                }
                
                
            }
            

            
            return {	
                uuid:this_uuid,
                mid:this_mid,
                
                fec_content:this_fec_content,
                fec_market_group:this_fec_market_group,
                fec_campaign:this_fec_campaign,
                fec_source:this_fec_source,
                fec_medium:this_fec_medium,
                fec_design:this_fec_design,
                customer_email:this_customer_email,
                
                customer_id:this_customer_id,
                pv:this_pv,
                stay_seconds:this_stay_seconds,
                service_date_str:this_service_date_str,	
                register_email:this_register_email,
                login_email:this_login_email,
                
                sku:this_sku,
                sku_cart:this_sku_cart,
                sku_order:this_sku_order,
                sku_order_success:this_sku_order_success,
                
            
                category:this_category,
                search:this_search,
                cart:this_cart,
                order:this_order,
                
                visit_page_sku:this_visit_page_sku,
                visit_page_category:this_visit_page_category,
                visit_page_search:this_visit_page_search,
                visit_page_cart:this_visit_page_cart,
                
                visit_page_order:this_visit_page_order,
                visit_page_order_processing:this_visit_page_order_processing,
                visit_page_order_pending:this_visit_page_order_pending,
                visit_page_order_amount:this_visit_page_order_amount,
                visit_page_order_processing_amount:this_visit_page_order_processing_amount,
                visit_page_order_pending_amount:this_visit_page_order_pending_amount,
                
                
                color_depth:this_color_depth,
                ip:this_ip,
                country_code:this_country_code,
                country_name:this_country_name,
                devide:this_devide,
                browser_name:this_browser_name,
                browser_lang:this_browser_lang,
                browser_version:this_browser_version,
                operate:this_operate,
                refer_url:this_refer_url,
                first_referrer_domain:this_first_referrer_domain,
                is_return:this_is_return,
                first_page_url:this_first_page_url,
                out_page:this_out_page,
                resolution:this_resolution,
                device_pixel_ratio:this_device_pixel_ratio,
                data:this_data
            };
        }
    `
    
    finalizeStr := `
        function (key, reducedVal) {
            reducedVal.website_id        = "` + website_id + `"
            return reducedVal;
        }
    `
    // 结果输出的 mongodb collection
    outDoc := bson.M{"replace": outCollName}
    // 执行mapreduce的job struct
    job := &mgo.MapReduce{
        Map:      mapStr,
        Reduce:   reduceStr,
        Finalize: finalizeStr,
        Out:      outDoc,
    }
    // 开始执行map reduce
    err = mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        _, err := coll.Find(nil).MapReduce(job, nil)
        return err
    })
    if err != nil {
        return err
    }
    // 上面mongodb maoreduce处理完的数据，需要存储到es中
    // 得到 type 以及 index name
    esCustomerUuidTypeName :=  helper.GetEsCustomerUuidTypeName()
    esIndexName := helper.GetEsIndexNameByType(esCustomerUuidTypeName)
    // es index 的type mapping
    esCustomerUuidTypeMapping := helper.GetEsCustomerUuidTypeMapping()
    // 删除index，如果mapping建立的不正确，可以执行下面的语句删除重建mapping
    //err = esdb.DeleteIndex(esIndexName)
    //if err != nil {
    //    return err
    //}
    // 初始化mapping
    err = esdb.InitMapping(esIndexName, esCustomerUuidTypeName, esCustomerUuidTypeMapping)
    if err != nil {
        return err
    }
    // 同步mongo数据到ES
    // mongodb中的数据总数
    mCount := 0
    // 得到总数
    err = mongodb.MDC(dbName, outCollName, func(coll *mgo.Collection) error {
        var err error
        mCount, err = coll.Count() 
        return err
    })
    if err != nil {
        return err
    }
    numPerPage := helper.BulkSyncCount
    pageNum := int(math.Ceil(float64(mCount) / float64(numPerPage)))
    for i:=0; i<pageNum; i++ {
        err = mongodb.MDC(dbName, outCollName, func(coll *mgo.Collection) error {
            var err error
            var customerUuids []model.CustomerUuid
            coll.Find(nil).Skip(i*pageNum).Limit(numPerPage).All(&customerUuids)
            log.Println("customerUuids length:")
            log.Println(len(customerUuids))
            
            /* 这个代码是upsert单行数据
            for j:=0; j<len(customerUuids); j++ {
                CustomerUuid := customerUuids[j]
                customerUuidValue := customerUuid.Value
                // customerUuidValue.Devide = nil
                // customerUuidValue.CountryCode = nil
                ///customerUuidValue.Operate = nil
                log.Println("ID_:" + customerUuid.Id_)
                customerUuidValue.Id = customerUuid.Id_
                err := esdb.UpsertType(esIndexName, esCustomerUuidTypeName, customerUuid.Id_, customerUuidValue)
                
                if err != nil {
                    log.Println("11111" + err.Error())
                    return err
                }
            }
            */
            if len(customerUuids) > 0 {
                // 使用bulk的方式，将数据批量插入到elasticSearch
                bulkRequest, err := esdb.Bulk()
                if err != nil {
                    log.Println("444" + err.Error())
                    return err
                }
                for j:=0; j<len(customerUuids); j++ {
                    customerUuid := customerUuids[j]
                    customerUuidValue := customerUuid.Value
                    customerUuidValue.Id = customerUuid.Id_
                    log.Println("888")
                    log.Println(esIndexName)
                    log.Println(esCustomerUuidTypeName)
                    log.Println(customerUuid.Id_)
                    log.Println(customerUuidValue)
                    req := esdb.BulkUpsertTypeDoc(esIndexName, esCustomerUuidTypeName, customerUuid.Id_, customerUuidValue)
                    bulkRequest = bulkRequest.Add(req)
                }
                bulkResponse, err := esdb.BulkRequestDo(bulkRequest)
                // bulkResponse, err := bulkRequest.Do()
                if err != nil {
                    log.Println("#############3")
                    log.Println("333" + err.Error())
                    return err
                }
                if bulkResponse != nil {
                    log.Println("#############4")
                    log.Println(bulkResponse)
                }
            }
            return err
        })
        
        if err != nil {
            log.Println("##############5" + err.Error())
            return err
        }
    }
    if err != nil {
        log.Println("#############6" + err.Error())
    }
    
    return err
}

/*
type MapReduce struct {
    Map      string      // Map Javascript function code (required)
    Reduce   string      // Reduce Javascript function code (required)
    Finalize string      // Finalize Javascript function code (optional)
    Out      interface{} // Output collection name or document. If nil, results are inlined into the result parameter.
    Scope    interface{} // Optional global scope for Javascript functions
    Verbose  bool
}
*/

/*
type resultValue struct{
    BrowserNameCount int64 `browser_name`
}

var result []struct { 
    Id string `_id`
    Value resultValue 
} dbName string, collName string, outCollName
*/