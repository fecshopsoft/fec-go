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

// 浏览器统计计算部分
func BrowserMapReduct(dbName string, collName string, outCollName string, esIndexName string) error {
    var err error
    mapStr := `
        function() {  
            browser_name = this.browser_name ? this.browser_name : null;
            // login_email 	= this.login_email ? [this.login_email] : null;
            
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
            
            // first_page
            first_page 		= this.uuid_first_page ? this.uuid_first_page : 0;
            
            // 该处进行了更正，不应该 first_visit_this_url ，而应该使用 uuid_first_page
            is_return = 0;
            uv = 0;
            if(this.uuid_first_page == 1){
                uv = 1;
                is_return = this.is_return ? this.is_return : 0;
            }
            service_date_str = this.service_date_str ? this.service_date_str : null;
            
            cart = this.cart ? this.cart : null;
            order = this.order ? this.order : null;
            
            cart_count = 0;
            order_count = 0;
            success_order_count = 0;
            success_order_no_count = 0;
            if(cart){
                for(x in cart){
                    one 		= cart[x];
                    if(one && one['qty']){
                        //$sku 		= one['sku'];
                        cart_count += Number(one['qty']);
                    }
                }
            }
            
            if(order && order['products']){
                products = order['products'];
                payment_status = order['payment_status'];
                if(payment_status == 'payment_confirmed'){
                    success_order_no_count = 1;
                }
                for(x in products){
                    one = products[x];
                    if(one && one['qty']){
                        qty = Number(one['qty']);
                        order_count += qty;
                        if(payment_status == 'payment_confirmed'){
                            success_order_count += qty;
                            
                        }
                    }
                }
            }
            
            if(this.browser_name){
                emit(this.browser_name+"_"+service_date_str,{
                    browser_name:browser_name,
                    pv:1,
                    uv:uv,
                    rate_pv:rate_pv,
                    stay_seconds:stay_seconds,
                    //customer_id:customer_id,
                    jump_out_count:jump_out_count,
                    drop_out_count:drop_out_count,
                    devide:devide,
                    country_code:country_code,
                    
                    cart_count:cart_count,
                    order_count:order_count,
                    success_order_count:success_order_count,
                    success_order_no_count:success_order_no_count,
                    
                    operate:operate,
                    is_return:Number(is_return),
                    first_page:Number(first_page),
                    service_date_str:service_date_str
                    
                });
            }
            
        }
    `
    
    reduceStr := `
        function(key,emits){
            this_pv 				= 0;
            this_rate_pv			= 0;
            this_uv 				= 0;
            this_stay_seconds 		= 0;
            this_jump_out_count		= 0;
            this_drop_out_count		= 0 ;
            this_service_date_str 	= null;
            this_devide				= {};
            this_country_code		= {};
            this_browser_name		= null;
            this_operate			= {};
            this_is_return			= 0;
            this_first_page			= 0;
            
            this_cart_count				= 0;
            this_order_count			= 0;	
            this_success_order_count	= 0;
            this_success_order_no_count	= 0;
            
            for(var i in emits){
                
                if(emits[i].cart_count){
                    this_cart_count 			+= emits[i].cart_count;
                }
                if(emits[i].order_count){
                    this_order_count 			+= emits[i].order_count;
                }
                if(emits[i].success_order_count){
                    this_success_order_count 	+= emits[i].success_order_count;
                }
                if(emits[i].success_order_no_count){
                    this_success_order_no_count += emits[i].success_order_no_count;
                }
                
                if(emits[i].browser_name){
                    this_browser_name = emits[i].browser_name;
                }
                if(emits[i].pv){
                    this_pv 			+= emits[i].pv;
                }
                if(emits[i].uv){
                    this_uv 			+= emits[i].uv;
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
                
                if(emits[i].is_return){
                    this_is_return 			+= emits[i].is_return;
                }
                if(emits[i].first_page){
                    this_first_page 		+= emits[i].first_page;
                }
            }
            
            return {	
                browser_name:this_browser_name,
                pv:this_pv,
                uv:this_uv,
                rate_pv:this_rate_pv,
                jump_out_count:this_jump_out_count,
                drop_out_count:this_drop_out_count,
                stay_seconds:this_stay_seconds,
                //customer_id:this_customer_id,
                    
                devide:this_devide,
                country_code:this_country_code,
                
                operate:this_operate,
                
                is_return:this_is_return,
                first_page:this_first_page,
                
                cart_count:this_cart_count,
                order_count:this_order_count,	
                success_order_count:this_success_order_count,
                success_order_no_count:this_success_order_no_count,
                
                service_date_str:this_service_date_str
            };
        }
    `
    
    finalizeStr := `
        function (key, reducedVal) {
            uv = reducedVal.uv;
            pv = reducedVal.pv;
            // 平均pv
            if(uv){
                this_pv_rate = pv/uv;
                this_pv_rate = this_pv_rate*10000;
                this_pv_rate = Math.ceil(this_pv_rate);
                this_pv_rate = this_pv_rate/10000;
            }else{
                this_pv_rate = 0;
            }
            // 销售转换率
            if(uv){
                this_success_order_no_count = reducedVal.success_order_no_count;
                this_sku_sale_rate = this_success_order_no_count/uv;
                this_sku_sale_rate = this_sku_sale_rate*10000;
                this_sku_sale_rate = Math.ceil(this_sku_sale_rate);
                this_sku_sale_rate = this_sku_sale_rate/10000;
            }else{
                this_sku_sale_rate = 0;
            }
            
            // 跳出率
            if(uv){
                this_jump_out_count = reducedVal.jump_out_count;
                this_jump_out_rate = this_jump_out_count/uv;
                this_jump_out_rate = this_jump_out_rate*10000;
                this_jump_out_rate = Math.ceil(this_jump_out_rate);
                this_jump_out_rate = this_jump_out_rate/10000;
            }else{
                this_jump_out_rate = 0;
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
            // 老用户比率
            if(uv){
                this_is_return = reducedVal.is_return;
                this_is_return_rate = (this_is_return/uv); 
                this_is_return_rate = this_is_return_rate*10000;
                this_is_return_rate = Math.ceil(this_is_return_rate);
                this_is_return_rate = this_is_return_rate/10000;
            }else{
                this_is_return_rate = 0;
            }
            
            reducedVal.is_return_rate= this_is_return_rate;
            this_stay_seconds = reducedVal.stay_seconds;
            //平均停留时间
            rate_pv = reducedVal.rate_pv;
            if(rate_pv){
                stay_seconds_rate = (this_stay_seconds/rate_pv); 
                stay_seconds_rate = stay_seconds_rate*10000;
                stay_seconds_rate = Math.ceil(stay_seconds_rate);
                stay_seconds_rate = stay_seconds_rate/10000;
            }else{
                stay_seconds_rate = 0;
            }
            reducedVal.stay_seconds_rate = stay_seconds_rate;
            reducedVal.jump_out_rate= this_jump_out_rate;
            reducedVal.drop_out_rate= this_drop_out_rate;
            reducedVal.sku_sale_rate= this_sku_sale_rate;
            //reducedVal.country_code = this_country_code;
            //reducedVal.browser_name = this_browser_name;
            //reducedVal.operate 		= this_operate;
            reducedVal.pv_rate		= this_pv_rate;
            return reducedVal;
        }
    `
    
    outDoc := bson.M{"replace": outCollName}
    
    job := &mgo.MapReduce{
        Map:      mapStr,
        Reduce:   reduceStr,
        Finalize: finalizeStr,
        Out:      outDoc,
    }
    
    err = mongodb.MDC(dbName, collName, func(coll *mgo.Collection) error {
        _, err := coll.Find(nil).MapReduce(job, nil)
        return err
    })
    if err != nil {
        return err
    }
    
    // esIndexName
    esWholeBrowserTypeName :=  helper.GetEsWholeBrowserTypeName()
    esWholeBrowserTypeMapping := helper.GetEsWholeBrowserTypeMapping()
    // 删除index
    err =esdb.DeleteIndex(esIndexName)
    if err != nil {
        return err
    }
    // 初始化mapping
    err = esdb.InitMapping(esIndexName, esWholeBrowserTypeName, esWholeBrowserTypeMapping)
    if err != nil {
        return err
    }
    // 同步mongo数据到ES
    // 得到数据总数
    mCount := 0
    
    err = mongodb.MDC(dbName, outCollName, func(coll *mgo.Collection) error {
        var err error
        mCount, err = coll.Count() 
        return err
    })
    if err != nil {
        return err
    }
    numPerPage := 10
    pageNum := int(math.Ceil(float64(mCount) / float64(numPerPage)))
    for i:=0; i<pageNum; i++ {
        err = mongodb.MDC(dbName, outCollName, func(coll *mgo.Collection) error {
            var err error
            var wholeBrowsers []model.WholeBrowser
            coll.Find(nil).Skip(i*pageNum).Limit(numPerPage).All(&wholeBrowsers)
            log.Println("wholeBrowsers length:")
            log.Println(len(wholeBrowsers))
            for j:=0; j<len(wholeBrowsers); j++ {
                wholeBrowser := wholeBrowsers[j]
                wholeBrowserValue := wholeBrowser.Value
                // wholeBrowserValue.Devide = nil
                // wholeBrowserValue.CountryCode = nil
                ///wholeBrowserValue.Operate = nil
                log.Println("ID_:" + wholeBrowser.Id_)
                
                err := esdb.UpsertType(esIndexName, esWholeBrowserTypeName, wholeBrowser.Id_, wholeBrowserValue)
                
                if err != nil {
                    log.Println("11111" + err.Error())
                    return err
                }
            }
            return err
        })
        
        if err != nil {
            log.Println("2222" + err.Error())
            return err
        }
    }
    if err != nil {
        log.Println("3333" + err.Error())
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