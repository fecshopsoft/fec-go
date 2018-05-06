package whole

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

// Search 统计计算部分
func SearchMapReduct(dbName string, collName string, outCollName string, website_id string) error {
    var err error
    
    mapStr := `
        function() {  
            website_id = "` + website_id + `";
            
            search = this.search ? this.search : null;
            search_text = null;
            search_qty  = 0;
            if(search && search.text ){
                search_text = search['text'] ? search['text'] : null;
                search_qty = search['result_qty'] ? search['result_qty'] : 0;
                if(search_text.length > 260){
                    search_text = search_text.substring(0,260);
                }
            }
            
            uv = this.uuid ? [this.uuid] : [];
            if(this.is_return && this.uuid){
                is_return = [this.uuid + this.is_return];
            }else{
                is_return = [];
            }
            stay_seconds = this.stay_seconds ? this.stay_seconds : 0;
            // 跳出个数和退出个数
            jump_out_count = 0;   //跳出
            drop_out_count = 0;		//退出
            // 最后一个页面
            rate_pv = 0;
            if(this.stay_seconds == 0){
                // 如果这个页面是第一个访问的页面，则代表跳出
                if(this.uuid_first_page == 1 ){
                    jump_out_count = 1;
                }else{
                    drop_out_count = 1;
                }
            }else{
                rate_pv = 1;
            }
    
            devide 			= this.devide;
            if(devide){
                b= {};
                b[devide] = 1;
                devide 	= b;
            }else{
                devide = null;
            }
            
            browser_name 			= this.browser_name;
            if(browser_name){
                b= {};
                b[browser_name] = 1;
                browser_name 	= b;
            }else{
                browser_name = null;
            }
            
            // country_code
            country_code 	= this.country_code;
            if(country_code){
                b= {};
                b[country_code] = 1;
                country_code 	= b;
            }else{
                country_code = null;
            }
            
            // operate
            operate 		= this.operate ;
            if(operate){
                b= {};
                b[operate] = 1;
                operate 	= b;
            }else{
                operate = null;
            }
            // fec_app
            fec_app 		= this.fec_app ;
            if(fec_app){
                b= {};
                b[fec_app] = 1;
                fec_app 	= b;
            }else{
                fec_app = null;
            }
            
            
            resolution 		= this.resolution ;
            if(resolution){
                b= {};
                b[resolution] = 1;
                resolution 	= b;
            }else{
                resolution = null;
            }
            
            color_depth 		= this.color_depth ;
            if(color_depth){
                b= {};
                b[color_depth] = 1;
                color_depth 	= b;
            }else{
                color_depth = null;
            }
            
            language 		= this.fec_lang;
            if(language){
                b= {};
                b[language] = 1;
                language 	= b;
            }else{
                language = null;
            }
            // first_page
            first_page 		= this.uuid_first_page ? this.uuid_first_page : 0;
            sku = this.sku ? this.sku : null;
            
            service_date_str = this.service_date_str ? this.service_date_str : null;
            
            search_sku_click = this.search_sku_click ? this.search_sku_click : 0;
            search_login_email = this.search_login_email ? this.search_login_email : 0;

            search_sku_cart = this.search_sku_cart ? this.search_sku_cart  : null;
            search_sku_order = this.search_sku_order ? this.search_sku_order  : null;
            search_sku_order_success = null;
            // 如果订单支付状态为成功状态，那么 search_sku_order_success = search_sku_order
            order = this.order ? this.order : null;
            if( search_sku_order && order && order['invoice'] && order['payment_status'] ){
                if (order['payment_status'] === 'payment_confirmed') {
                    search_sku_order_success = search_sku_order;
                }
            }
            
            if(search_text){
                if(sku){
                    emit(search_text+"_"+service_date_str+"_"+website_id,{
                        search_text:search_text,
                        website_id:website_id,
                        pv:0,
                        rate_pv:0,
                        search_qty:0,
                        uv:[],
                        stay_seconds:0,
                        jump_out_count:0,
                        drop_out_count:0,
                        devide:'',
                        country_code:'',
                        browser_name:'',
                        operate:'',
                        is_return:[],
                        first_page:0,

                        search_sku_click:search_sku_click,
                        search_login_email:0,
                        search_sku_cart:0,
                        search_sku_order:0,
                        search_sku_order_success:0,
                        
                        fec_app:'',
                        resolution:'',
                        color_depth:'',
                        language:'',
                        website_id:website_id,
                    
                        service_date_str:service_date_str

                    });

                }else{
                    emit(search_text+"_"+service_date_str+"_"+website_id,{
                        search_text:search_text,
                        pv:1,
                        rate_pv:rate_pv,
                        search_qty:Number(search_qty),
                        uv:uv,
                        stay_seconds:stay_seconds,
                        jump_out_count:jump_out_count,
                        drop_out_count:drop_out_count,
                        devide:devide,
                        country_code:country_code,
                        browser_name:browser_name,
                        operate:operate,
                        is_return:is_return,
                        first_page:Number(first_page),


                        search_sku_click:0,
                        search_login_email:search_login_email,
                        search_sku_cart:0,
                        search_sku_order:0,
                        search_sku_order_success:0,
                        
                        fec_app:fec_app,
                        resolution:resolution,
                        color_depth:color_depth,
                        language:language,
                        website_id:website_id,
                        
                        service_date_str:service_date_str
                    });
                }
            }
            
            
            if(search_sku_cart && (search_sku_cart  instanceof Object )){
                for(search_text in search_sku_cart){
                    k = search_sku_cart[search_text];
                    qty = 0;
                    if( ! Number(k) ){  //操蛋的地方，不知道为什么  (typeof k) == 'object'  和 k instanceof Object    数组也成立，不知道为什么艹！只能用这么托比的办法了，来辨别之前的一些错误格式的数据。
                        for(sku in k){
                            qty += k[sku];
                        }
                    }
                    if(qty){
                        if(search_text.length > 260){
                            search_text = search_text.substring(0,260);
                        }
                        emit(search_text+"_"+service_date_str+"_"+website_id,{
                            search_text:search_text,
                            pv:0,
                            rate_pv:0,
                            search_qty:0,
                            uv:[],
                            stay_seconds:0,
                            jump_out_count:0,
                            drop_out_count:0,
                            devide:'',
                            country_code:'',
                            browser_name:'',
                            operate:'',
                            is_return:[],
                            first_page:0,


                            search_sku_click:0,
                            search_login_email:0,
                            search_sku_cart:qty,
                            search_sku_order:0,
                            search_sku_order_success:0,
                            
                            fec_app:'',
                            resolution:'',
                            color_depth:'',
                            language:'',
                            website_id:website_id,
                            
                            service_date_str:service_date_str

                        });
                    }
                }
            }
            
            if(search_sku_order && (search_sku_order  instanceof Object )){
                for(search_text in search_sku_order){
                    k = search_sku_order[search_text];
                    qty = 0;
                    if( ! Number(k) ){  //操蛋的地方，不知道为什么  (typeof k) == 'object'  和 k instanceof Object    数组也成立，不知道为什么艹！只能用这么托比的办法了，来辨别之前的一些错误格式的数据。
                        for(sku in k){
                            qty += k[sku];
                        }
                    }
                    if(qty){
                        if(search_text.length > 260){
                            search_text = search_text.substring(0,260);
                        }
                        emit(search_text+"_"+service_date_str+"_"+website_id,{
                            search_text:search_text,
                            pv:0,
                            rate_pv:0,
                            search_qty:0,
                            uv:[],
                            stay_seconds:0,
                            jump_out_count:0,
                            drop_out_count:0,
                            devide:'',
                            country_code:'',
                            browser_name:'',
                            operate:'',
                            is_return:[],
                            first_page:0,

                            search_sku_click:0,
                            search_login_email:0,
                            search_sku_cart:0,
                            search_sku_order:qty,
                            search_sku_order_success:0,
                            
                            fec_app:'',
                            resolution:'',
                            color_depth:'',
                            language:'',
                            website_id:website_id,
                            
                            service_date_str:service_date_str

                        });
                    }
                }
            }
            
            //成功订单
            if(search_sku_order_success && ( search_sku_order_success  instanceof Object )){
                for(search_text in search_sku_order_success){

                    k = search_sku_order_success[search_text];
                    qty = 0;
                    if( ! Number(k) ){  //操蛋的地方，不知道为什么  (typeof k) == 'object'  和 k instanceof Object    数组也成立，不知道为什么艹！只能用这么托比的办法了，来辨别之前的一些错误格式的数据。
                        for(sku in k){
                            qty += k[sku];
                        }
                    }
                    if(qty){
                        if(search_text.length > 260){
                            search_text = search_text.substring(0,260);
                        }
                        emit(search_text+"_"+service_date_str+"_"+website_id,{
                            search_text:search_text,
                            pv:0,
                            rate_pv:0,
                            search_qty:0,
                            uv:[],
                            stay_seconds:0,
                            jump_out_count:0,
                            drop_out_count:0,
                            devide:'',
                            country_code:'',
                            browser_name:'',
                            operate:'',
                            is_return:[],
                            first_page:0,

                            search_sku_click:0,
                            search_login_email:0,
                            search_sku_cart:0,
                            search_sku_order:0,
                            search_sku_order_success:1,
                            
                            fec_app:'',
                            resolution:'',
                            color_depth:'',
                            language:'',
                            website_id:website_id,
                            
                            service_date_str:service_date_str

                        });
                    }
                }
            }
            
            
        }
    `
    
    reduceStr := `
        function(key,emits){
            this_search_text		= null;
            this_search_qty			= 0;
            this_pv 				= 0;
            this_rate_pv			= 0;
            this_uv 				= [];
            this_stay_seconds 		= 0;
            //this_ 		= [];
            this_jump_out_count		= 0;
            this_drop_out_count		= 0 ;
            this_service_date_str 	= null;
            this_devide				= {};
            this_country_code		= {};
            this_browser_name		= {};
            this_operate			= {};
            this_is_return			= [];
            this_first_page			= 0;
            
            this_search_sku_click 			= 0;
            this_search_login_email 		= 0;
            this_search_sku_cart 			= 0;
            this_search_sku_order			= 0;
            this_search_sku_order_success 	= 0;
            
            this_fec_app      			= {};
            this_resolution			= {};
            this_color_depth		= {};
            this_language			= {};
            
            for(var i in emits){
                if(emits[i].search_text){
                    this_search_text	=  emits[i].search_text;
                }
                if(emits[i].search_qty){
                    this_search_qty		+=  emits[i].search_qty;
                }
                if(emits[i].pv){
                    this_pv 			+= emits[i].pv;
                }
            
                if(emits[i].uv){
                    //this_uv 			+= emits[i].uv;
                    this_uv = this_uv.concat(emits[i].uv);
                }
                if(emits[i].is_return){
                    //this_uv 			+= emits[i].uv;
                    this_is_return = this_uv.concat(emits[i].is_return);
                }
            
                if(emits[i].stay_seconds){
                    this_stay_seconds 	+= emits[i].stay_seconds;
                }
                if(emits[i].service_date_str){
                    this_service_date_str = emits[i].service_date_str;
                }
                //this_customer_id = this_customer_id.concat(emits[i].customer_id);
                if(emits[i].jump_out_count){
                    this_jump_out_count += emits[i].jump_out_count;
                }
                if(emits[i].drop_out_count){
                    this_drop_out_count += emits[i].drop_out_count;
                }
                if(emits[i].rate_pv){
                    this_rate_pv 		+= emits[i].rate_pv;
                }
                
                if(emits[i].devide){
                    devide = emits[i].devide;
                    for(brower_ne in devide){
                        
                        count = devide[brower_ne];
                        if(!this_devide[brower_ne]){
                            this_devide[brower_ne] = count;
                        }else{
                            this_devide[brower_ne] += count;
                        }
                    }
                }
                
                
                
                
                if(emits[i].country_code){
                    country_code = emits[i].country_code;
                    for(brower_ne in country_code){
                        
                        count = country_code[brower_ne];
                        if(!this_country_code[brower_ne]){
                            this_country_code[brower_ne] = count;
                        }else{
                            this_country_code[brower_ne] += count;
                        }
                    }
                }
                
                if(emits[i].browser_name){
                    browser_name = emits[i].browser_name;
                    for(brower_ne in browser_name){
                        
                        count = browser_name[brower_ne];
                        if(!this_browser_name[brower_ne]){
                            this_browser_name[brower_ne] = count;
                        }else{
                            this_browser_name[brower_ne] += count;
                        }
                    }
                }
                
                
                
                if(emits[i].operate){
                    operate = emits[i].operate;
                    for(brower_ne in operate){
                        
                        count = operate[brower_ne];
                        if(!this_operate[brower_ne]){
                            this_operate[brower_ne] = count;
                        }else{
                            this_operate[brower_ne] += count;
                        }
                    }
                }
                
                 if(emits[i].fec_app){
                    fec_app = emits[i].fec_app;
                    for(brower_ne in fec_app){
                        
                        count = fec_app[brower_ne];
                        if(!this_fec_app[brower_ne]){
                            this_fec_app[brower_ne] = count;
                        }else{
                            this_fec_app[brower_ne] += count;
                        }
                    }
                }
                
                if(emits[i].resolution){
                    resolution = emits[i].resolution;
                    for(brower_ne in resolution){
                        count = resolution[brower_ne];
                        if(!this_resolution[brower_ne]){
                            this_resolution[brower_ne] = count;
                        }else{
                            this_resolution[brower_ne] += count;
                        }
                    }
                }
                if(emits[i].color_depth){
                    color_depth = emits[i].color_depth;
                    for(brower_ne in color_depth){
                        
                        count = color_depth[brower_ne];
                        if(!this_color_depth[brower_ne]){
                            this_color_depth[brower_ne] = count;
                        }else{
                            this_color_depth[brower_ne] += count;
                        }
                    }
                }
                if(emits[i].language){
                    language = emits[i].language;
                    for(brower_ne in language){
                        
                        count = language[brower_ne];
                        if(!this_language[brower_ne]){
                            this_language[brower_ne] = count;
                        }else{
                            this_language[brower_ne] += count;
                        }
                    }
                }
                
                if(emits[i].first_page){
                    this_first_page 		+= emits[i].first_page;
                }
                
                if(emits[i].search_sku_click){
                    this_search_sku_click		+=  emits[i].search_sku_click;
                }
                
                if(emits[i].search_login_email){
                    this_search_login_email		+=  emits[i].search_login_email;
                }
                
                if(emits[i].search_sku_cart){
                    this_search_sku_cart		+=  emits[i].search_sku_cart;
                }
                
                if(emits[i].search_sku_order){
                    this_search_sku_order		+=  emits[i].search_sku_order;
                }
                
                if(emits[i].search_sku_order_success){
                    this_search_sku_order_success+=  emits[i].search_sku_order_success;
                }
            }
            
            
            return {	
                search_text:this_search_text,
                search_qty:this_search_qty,
                pv:this_pv,
                rate_pv:this_rate_pv,
                uv:this_uv,
                jump_out_count:this_jump_out_count,
                drop_out_count:this_drop_out_count,
                stay_seconds:this_stay_seconds,
                //customer_id:this_customer_id,
                    
                devide:this_devide,
                country_code:this_country_code,
                browser_name:this_browser_name,
                operate:this_operate,
                
                is_return:this_is_return,
                first_page:this_first_page,
                
                search_sku_click:this_search_sku_click,
                search_login_email:this_search_login_email,
                search_sku_cart:this_search_sku_cart,
                search_sku_order:this_search_sku_order,
                search_sku_order_success:this_search_sku_order_success,
                fec_app:this_fec_app,
                resolution:this_resolution,
                color_depth:this_color_depth,
                language:this_language,
                service_date_str:this_service_date_str
            };
        }
    `
    
    finalizeStr := `
        function (key, reducedVal) {
            search_qty = reducedVal.search_qty;
            uv = reducedVal.uv;
            pv = reducedVal.pv;
            if (pv) {
                search_qty = Math.ceil(search_qty/pv);
            } else {
                search_qty = 0;
            }
            
            
            //
            uv.sort();
            var re=[uv[0]];
            for(var i = 1; i < uv.length; i++)
            {
                if( uv[i] !== re[re.length-1])
                {
                    re.push(uv[i]);
                }
            }
            uv = re.length;
            reducedVal.uv = uv;
            
            // is_return 
            is_return = reducedVal.is_return;
            is_return.sort();
            var re=[is_return[0]];
            for(var i = 1; i < is_return.length; i++)
            {
                if( is_return[i] !== re[re.length-1])
                {
                    re.push(is_return[i]);
                }
            }
            is_return = re.length;
            reducedVal.is_return = is_return;
            
            
            // 平均pv
            if(uv){
                this_pv_rate = pv/uv;
                this_pv_rate = this_pv_rate*10000;
                this_pv_rate = Math.ceil(this_pv_rate);
                this_pv_rate = this_pv_rate/10000;
            }else{
                this_pv_rate = 0;
            }
            
            // 跳出率
            this_jump_out_count = reducedVal.jump_out_count;
            if(uv){
                this_jump_out_rate = this_jump_out_count/uv;
                this_jump_out_rate = this_jump_out_rate*10000;
                this_jump_out_rate = Math.ceil(this_jump_out_rate);
                this_jump_out_rate = this_jump_out_rate/10000;
            }else{
                this_jump_out_rate = 0;
            }
            
            
            // 点击率
            this_search_sku_click = reducedVal.search_sku_click;
            if(pv){
                this_search_sku_click_rate = this_search_sku_click/pv;
                this_search_sku_click_rate = this_search_sku_click_rate*10000;
                this_search_sku_click_rate = Math.ceil(this_search_sku_click_rate);
                this_search_sku_click_rate = this_search_sku_click_rate/10000;
            }else{
                this_search_sku_click_rate = 0;
            }
            
            
            // 转化率
            this_search_sku_order_success = reducedVal.search_sku_order_success;
            if(uv){
                this_search_sale_rate = this_search_sku_order_success/uv;
                this_search_sale_rate = this_search_sale_rate*10000;
                this_search_sale_rate = Math.ceil(this_search_sale_rate);
                this_search_sale_rate = this_search_sale_rate/10000;
            }else{
                this_search_sale_rate = 0;
            }
            
            
            // 退出率
            if(uv){
                this_drop_out_count = reducedVal.drop_out_count;
                this_drop_out_rate = this_drop_out_count/uv;
                this_drop_out_rate = this_drop_out_rate*10000;
                this_drop_out_rate = Math.ceil(this_drop_out_rate);
                this_drop_out_rate = this_drop_out_rate/10000;
            }else{
                this_drop_out_rate = 0;
            }
            
            
            
            //平均停留时间
            rate_pv = reducedVal.rate_pv;
            this_stay_seconds = reducedVal.stay_seconds;
            if(rate_pv){
                stay_seconds_rate = (this_stay_seconds/rate_pv); 
                stay_seconds_rate = stay_seconds_rate*10000;
                stay_seconds_rate = Math.ceil(stay_seconds_rate);
                stay_seconds_rate = stay_seconds_rate/10000;
            }else{
                stay_seconds_rate = 0;
            }
            
            this_is_return = reducedVal.is_return;
            if(pv){
                this_is_return_rate = (this_is_return/pv); 
                this_is_return_rate = this_is_return_rate*10000;
                this_is_return_rate = Math.ceil(this_is_return_rate);
                this_is_return_rate = this_is_return_rate/10000;
            }else{
                this_is_return_rate = 0;
            }
            reducedVal.search_qty = search_qty;
            reducedVal.is_return_rate= this_is_return_rate;
            reducedVal.stay_seconds_rate = stay_seconds_rate;
            reducedVal.jump_out_rate= this_jump_out_rate;
            reducedVal.drop_out_rate= this_drop_out_rate;
            
            reducedVal.search_sku_click_rate = this_search_sku_click_rate;
            reducedVal.search_sale_rate = this_search_sale_rate;
            
            //reducedVal.devide 		= this_devide;
            //reducedVal.country_code = this_country_code;
            //reducedVal.browser_name = this_browser_name;
            //reducedVal.operate 		= this_operate;
            reducedVal.pv_rate		= this_pv_rate;
            reducedVal.website_id        = "` + website_id + `";
            reducedVal.ip_count = reducedVal.uv;
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
    esWholeSearchTypeName :=  helper.GetEsWholeSearchTypeName()
    esIndexName := helper.GetEsIndexNameByType(esWholeSearchTypeName)
    // es index 的type mapping
    esWholeSearchTypeMapping := helper.GetEsWholeSearchTypeMapping()
    // 删除index，如果mapping建立的不正确，可以执行下面的语句删除重建mapping
    //err = esdb.DeleteIndex(esIndexName)
    //if err != nil {
    //    return err
    //}
    // 初始化mapping
    err = esdb.InitMapping(esIndexName, esWholeSearchTypeName, esWholeSearchTypeMapping)
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
            var wholeSearchs []model.WholeSearch
            coll.Find(nil).Skip(i*pageNum).Limit(numPerPage).All(&wholeSearchs)
            log.Println("wholeSearchs length:")
            log.Println(len(wholeSearchs))
            
            /* 这个代码是upsert单行数据
            for j:=0; j<len(wholeSearchs); j++ {
                wholeBrowser := wholeSearchs[j]
                wholeBrowserValue := wholeBrowser.Value
                // wholeBrowserValue.Devide = nil
                // wholeBrowserValue.CountryCode = nil
                ///wholeBrowserValue.Operate = nil
                log.Println("ID_:" + wholeBrowser.Id_)
                wholeBrowserValue.Id = wholeBrowser.Id_
                err := esdb.UpsertType(esIndexName, esWholeSkuTypeName, wholeBrowser.Id_, wholeBrowserValue)
                
                if err != nil {
                    log.Println("11111" + err.Error())
                    return err
                }
            }
            */
            if len(wholeSearchs) > 0 {
                // 使用bulk的方式，将数据批量插入到elasticSearch
                bulkRequest, err := esdb.Bulk()
                if err != nil {
                    log.Println("444" + err.Error())
                    return err
                }
                for j:=0; j<len(wholeSearchs); j++ {
                    wholeSearch := wholeSearchs[j]
                    wholeSearchValue := wholeSearch.Value
                    wholeSearchValue.Id = wholeSearch.Id_
                    log.Println("888")
                    log.Println(esIndexName)
                    log.Println(esWholeSearchTypeName)
                    log.Println(wholeSearch.Id_)
                    log.Println(wholeSearchValue)
                    req := esdb.BulkUpsertTypeDoc(esIndexName, esWholeSearchTypeName, wholeSearch.Id_, wholeSearchValue)
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
