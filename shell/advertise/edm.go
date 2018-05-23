package advertise

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

// Edm 统计计算部分
func EdmMapReduct(dbName string, collName string, outCollName string, website_id string) error {
    var err error
    mapStr := `
        function() {  
            website_id = "` + website_id + `";
            fec_source = this.fec_source ? this.fec_source : null;
            fec_campaign = this.fec_campaign ? this.fec_campaign : null;
            fid 	= this.fid ? this.fid : null;
            if (fid && fec_source == 'EDM') {
                fec_edm = fid + '_' + fec_campaign
                fec_medium 	= this.fec_medium;
                if(fec_medium){
                    b= {};
                    b[fec_medium] = 1;
                    fec_medium 	= b;
                }else{
                    fec_medium = null;
                }
                
                
                fec_content 	= this.fec_content;
                if(fec_content){
                    b= {};
                    b[fec_content] = 1;
                    fec_content 	= b;
                }else{
                    fec_content = null;
                }
                
                fec_market_group 	= this.fec_market_group;
                if(fec_market_group){
                    b= {};
                    b[fec_market_group] = 1;
                    fec_market_group 	= b;
                }else{
                    fec_market_group = null;
                }
                
                fec_design 	= this.fec_design;
                if(fec_design){
                    b= {};
                    b[fec_design] = 1;
                    fec_design 	= b;
                }else{
                    fec_design = null;
                }
                
                
                
                first_referrer_domain = this.first_referrer_domain ? this.first_referrer_domain : null;
                if(first_referrer_domain){
                    first_referrer_domain.replace(/./, "##");
                    b= {};
                    b[first_referrer_domain] = 1;
                    first_referrer_domain 	= b;
                }else{
                    first_referrer_domain = null;
                }
        
        
                stay_seconds = this.stay_seconds ? this.stay_seconds : 0;
                
                // 跳出个数和退出个数
                jump_out_count = 0;   //跳出
                drop_out_count = 0;		//退出
                // 最后一个页面
                rate_pv = 0;
                drop_out_page_info = null;
                if(this.stay_seconds == 0){
                    // 如果这个页面是第一个访问的页面，则代表跳出
                    if(this.uuid_first_page == 1 ){
                        jump_out_count = 1;
                    }else{
                        drop_out_count = 1;
                        b= {};
                        url_new = this.url_new;
                        b[url_new] = 1;
                        drop_out_page_info 	= b;
                    }
                }else{
                    rate_pv = 1;
                }
                
                browser_name 			= this.browser_name;
                if(browser_name){
                    b= {};
                    b[browser_name] = 1;
                    browser_name 	= b;
                }else{
                    browser_name = null;
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
                
                sku_visit_info 		= this.sku ;
                if(sku_visit_info){
                    b= {};
                    b[sku_visit_info] = 1;
                    sku_visit_info 	= b;
                }else{
                    sku_visit_info = null;
                }
                
                category_visit_info 		= this.category ;
                if(category_visit_info){
                    b= {};
                    b[category_visit_info] = 1;
                    category_visit_info    = b;
                }else{
                    category_visit_info = null;
                }
                
                // search 
                search_visit_info = null;
                search 		= this.search ;
                if(search && search['text']){
                    b= {};
                    b[search['text']] = 1;
                    search_visit_info 	= b;
                }else{
                    search_visit_info = null;
                }
                
                
                first_page 		= this.uuid_first_page ? this.uuid_first_page : 0;
                is_return = 0;
                if(this.uuid_campaign_first_fid == 1){
                    is_return 		= this.is_return ? this.is_return : 0;
                }
                
                if(this.uuid_campaign_first_fid == 1){
                    uv = 1;
                }else{
                    uv = 0;
                }
                if(this.ip_first_fid == 1){
                    ip_count = 1;
                }else{
                    ip_count = 0;
                }
                
                service_date_str = this.service_date_str ? this.service_date_str : null;
                
                // 注册email
                register_email = this.register_email;
                if(register_email){
                    register_count = 1;
                }else{
                    register_count = 0;
                }
                
                // 登录email
                login_email = this.login_email;
                if(login_email){
                    login_count = 1;
                }else{
                    login_count = 0;
                }
                
                // category_count
                category = this.category;
                if(category){
                    category_count = 1;
                }else{
                    category_count = 0;
                }
                // sku_count
                sku = this.sku;
                if(sku){
                    sku_count = 1;
                }else{
                    sku_count = 0;
                }
                
                // search_count
                search = this.search;
                if(search && search['text']){
                    search_count = 1;
                }else{
                    search_count = 0;
                }
                
                
                cart = this.cart ? this.cart : null;
                order = this.order ? this.order : null;
                
                cart_count = 0;
                order_count = 0;
                order_no_count = 0;
                success_order_count = 0;
                success_order_no_count = 0;
                order_amount = 0;
                success_order_amount = 0;
                
                cart_sku_info = null;
                
                if(cart){
                    for(x in cart){
                        one 		= cart[x];
                        if(one && one['qty']){
                            $sku 		= one['sku'];
                            var skuqty = Number(one['qty'])
                            skuqty = isNaN(skuqty) ? 0 : skuqty
                            cart_count += skuqty;
                            if($sku && skuqty){
                                if(!cart_sku_info){
                                    cart_sku_info = {};
                                }
                                cart_sku_info[$sku] = skuqty;
                            }
                        }
                    }
                }
                
                order_sku_info = null;
                success_order_sku_info = null;
                
                order_increment_id = null;
                fail_order_increment_id = null;
                success_order_increment_id = null;
                
                success_order_info = null;
                fail_order_info = null;
        
                if(order && order['invoice'] && order['products']){
                    products = order['products'];
                    payment_status = order['payment_status'];
                    amount = order['amount'];
                    if(amount){
                        order_amount = amount;
                    }
                    if(order['invoice']){
                        order_increment_id = [order['invoice']];
                    }
                    order_no_count = 1;
                    if(payment_status == 'payment_confirmed'){
                        success_order_no_count = 1;
                        if(amount){
                            success_order_amount = amount;
                        }
                        success_order_increment_id = [order['invoice']];
                        success_order_info = [order];
                    } else {
                        fail_order_increment_id = [order['invoice']];
                        fail_order_info = [order];
                    }
                    for(x in products){
                        one = products[x];
                        if(one && one['qty'] && one['sku']){
                            sku = one['sku']
                            qty = Number(one['qty']);
                            qty = isNaN(qty) ? 0 : qty
                            order_count += qty;
                            if(!order_sku_info){
                                order_sku_info = {};
                            }
                            order_sku_info[sku] = qty;
                            if(payment_status == 'payment_confirmed'){
                                success_order_count += qty;
                                if(!success_order_sku_info){
                                    success_order_sku_info = {};
                                }
                                success_order_sku_info[sku] = qty;
                            }
                        }
                    }
                }
                
                if(fec_edm){
                    is_return = Number(is_return);
                    is_return = isNaN(is_return) ? 0 : is_return
                    first_page = Number(first_page);
                    first_page = isNaN(first_page) ? 0 : first_page
                    emit(fec_edm+"_"+service_date_str+"_"+website_id,{
                        fid: fid,
                        fec_edm: fec_edm,
                        fec_design: fec_design,
                        fec_content: fec_content,
                        fec_source: fec_source,
                        fec_market_group: fec_market_group,
                        fec_campaign: fec_campaign,
                        fec_medium: fec_medium,
                        first_referrer_domain: first_referrer_domain,
                        // drop_out_page_info: drop_out_page_info,
                        sku_visit_info: sku_visit_info,
                        category_visit_info: category_visit_info,
                        search_visit_info: search_visit_info,
                        
                        cart_sku_info: cart_sku_info,
                        order_sku_info: order_sku_info,
                        success_order_sku_info: success_order_sku_info,
                        order_increment_id: order_increment_id,
                        fail_order_increment_id: fail_order_increment_id,
                        success_order_increment_id: success_order_increment_id,
                        success_order_info: success_order_info,
                        fail_order_info: fail_order_info,
                
                        register_count: register_count,
                        login_count: login_count,
                        category_count: category_count,
                        sku_count: sku_count,
                        search_count: search_count,
                        
                        
                        browser_name:browser_name,
                        pv:1,
                        uv:uv,
                        ip_count:ip_count,
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
                        order_no_count:order_no_count,
                        order_amount:order_amount,
                        success_order_amount:success_order_amount,
                        operate:operate,
                        fec_app:fec_app,
                        resolution:resolution,
                        color_depth:color_depth,
                        language:language,
                        is_return: is_return,
                        first_page: first_page,
                        service_date_str:service_date_str
                        
                    });
                }
            }
        }
    `
    
    reduceStr := `
        function(key,emits){
            this_pv 				= 0;
            this_rate_pv			= 0;
            this_uv 				= 0;
            this_ip_count 			= 0;
            this_stay_seconds 		= 0;
            this_jump_out_count		= 0;
            this_drop_out_count		= 0 ;
            this_service_date_str 	= null;
            this_devide				= {};
            this_country_code		= {};
            this_browser_name		= {};
            this_operate			= {};
            this_fec_app      			= {};
            this_is_return			= 0;
            this_first_page			= 0;
            this_resolution			= {};
            this_color_depth		= {};
            this_language			= {};
            this_cart_count				= 0;
            this_order_count			= 0;	
            this_success_order_count	= 0;
            this_order_amount			= 0;
            this_success_order_amount	= 0;
            this_order_no_count	= 0;
            this_success_order_no_count	= 0;
            
            this_fec_source 	= null;
            this_fec_campaign 	= null;
            this_fec_edm    	= null;
            this_fid 	= null;
            
            this_fec_medium 	= {};
            this_fec_content 	= {};
            this_fec_market_group 	= {};
            this_fec_design 	= {};
            
            this_first_referrer_domain	 		= {};
            
            this_sku_visit_info			= {};
            this_category_visit_info	= {};
            this_search_visit_info		= {};
            this_cart_sku_info			= {};
            this_success_order_sku_info	= {};
            this_order_sku_info			= {};
            
            this_order_increment_id = null;	
            this_fail_order_increment_id= null;
            this_success_order_increment_id	= null;
            this_success_order_info	= null;
            this_fail_order_info	= null;

            this_register_count  	= 0;
            this_login_count 		= 0;
            this_category_count 	= 0;
            this_sku_count 			= 0;
            this_search_count 		= 0;        
            
            for(var i in emits){
                
                //if( emits[i].fec_market_group){
                //    this_fec_market_group 		=  emits[i].fec_market_group;
                //}
                if(emits[i].fec_market_group){
                    fec_market_group = emits[i].fec_market_group;
                    for(brower_ne in fec_market_group){
                        count = fec_market_group[brower_ne];
                        if(!this_fec_market_group[brower_ne]){
                            this_fec_market_group[brower_ne] = count;
                        }else{
                            this_fec_market_group[brower_ne] += count;
                        }
                    }
                }
                //if( emits[i].fec_content){
                //    this_fec_content 		=  emits[i].fec_content;
                //}
                if(emits[i].fec_content){
                    fec_content = emits[i].fec_content;
                    for(brower_ne in fec_content){
                        count = fec_content[brower_ne];
                        if(!this_fec_content[brower_ne]){
                            this_fec_content[brower_ne] = count;
                        }else{
                            this_fec_content[brower_ne] += count;
                        }
                    }
                }
                
                if( emits[i].fid){
                    this_fid 		=  emits[i].fid;
                }
                if( emits[i].fec_edm){
                    this_fec_edm 		=  emits[i].fec_edm;
                }
                if( emits[i].fec_source){
                    this_fec_source 		=  emits[i].fec_source;
                }
                if( emits[i].fec_campaign){
                    this_fec_campaign 		=  emits[i].fec_campaign;
                }
                
                if(emits[i].fec_medium){
                    fec_medium = emits[i].fec_medium;
                    for(brower_ne in fec_medium){
                        count = fec_medium[brower_ne];
                        if(!this_fec_medium[brower_ne]){
                            this_fec_medium[brower_ne] = count;
                        }else{
                            this_fec_medium[brower_ne] += count;
                        }
                    }
                }
                
                if(emits[i].fec_design){
                    fec_design = emits[i].fec_design;
                    for(brower_ne in fec_design){
                        count = fec_design[brower_ne];
                        if(!this_fec_design[brower_ne]){
                            this_fec_design[brower_ne] = count;
                        }else{
                            this_fec_design[brower_ne] += count;
                        }
                    }
                }
                
                if(emits[i].first_referrer_domain){
                    first_referrer_domain = emits[i].first_referrer_domain;
                    for(brower_ne in first_referrer_domain){
                        
                        count = first_referrer_domain[brower_ne];
                        if(!this_first_referrer_domain[brower_ne]){
                            this_first_referrer_domain[brower_ne] = count;
                        }else{
                            this_first_referrer_domain[brower_ne] += count;
                        }
                    }
                }
                if(emits[i].sku_visit_info){
                    sku_visit_info = emits[i].sku_visit_info;
                    for(brower_ne in sku_visit_info){
                        
                        count = sku_visit_info[brower_ne];
                        if(!this_sku_visit_info[brower_ne]){
                            this_sku_visit_info[brower_ne] = count;
                        }else{
                            this_sku_visit_info[brower_ne] += count;
                        }
                    }
                }
                if(emits[i].category_visit_info){
                    category_visit_info = emits[i].category_visit_info;
                    for(brower_ne in category_visit_info){
                        
                        count = category_visit_info[brower_ne];
                        if(!this_category_visit_info[brower_ne]){
                            this_category_visit_info[brower_ne] = count;
                        }else{
                            this_category_visit_info[brower_ne] += count;
                        }
                    }
                }
                if(emits[i].search_visit_info){
                    search_visit_info = emits[i].search_visit_info;
                    for(brower_ne in search_visit_info){
                        
                        count = search_visit_info[brower_ne];
                        if(!this_search_visit_info[brower_ne]){
                            this_search_visit_info[brower_ne] = count;
                        }else{
                            this_search_visit_info[brower_ne] += count;
                        }
                    }
                }
                //this_cart_sku_info			= {};
                if(emits[i].cart_sku_info){
                    cart_sku_info = emits[i].cart_sku_info;
                    for(brower_ne in cart_sku_info){
                        
                        count = cart_sku_info[brower_ne];
                        if(!this_cart_sku_info[brower_ne]){
                            this_cart_sku_info[brower_ne] = count;
                        }else{
                            this_cart_sku_info[brower_ne] += count;
                        }
                    }
                }
                //this_success_order_sku_info	= {};
                if(emits[i].success_order_sku_info){
                    success_order_sku_info = emits[i].success_order_sku_info;
                    for(brower_ne in success_order_sku_info){
                        
                        count = success_order_sku_info[brower_ne];
                        if(!this_success_order_sku_info[brower_ne]){
                            this_success_order_sku_info[brower_ne] = count;
                        }else{
                            this_success_order_sku_info[brower_ne] += count;
                        }
                    }
                }
                //this_order_sku_info			= {};
                if(emits[i].order_sku_info){
                    order_sku_info = emits[i].order_sku_info;
                    for(brower_ne in order_sku_info){
                        
                        count = order_sku_info[brower_ne];
                        if(!this_order_sku_info[brower_ne]){
                            this_order_sku_info[brower_ne] = count;
                        }else{
                            this_order_sku_info[brower_ne] += count;
                        }
                    }
                }
                if( emits[i].order_increment_id){
                    if(!this_order_increment_id){
                        this_order_increment_id 		=  emits[i].order_increment_id;
                    }else{
                        this_order_increment_id = this_order_increment_id.concat(emits[i].order_increment_id);
                    }
                }
                if( emits[i].fail_order_increment_id){
                    if(!this_fail_order_increment_id){
                        this_fail_order_increment_id 		=  emits[i].fail_order_increment_id;
                    }else{
                        this_fail_order_increment_id = this_fail_order_increment_id.concat(emits[i].fail_order_increment_id);
                    }
                }
                if( emits[i].success_order_increment_id){
                    if(!this_success_order_increment_id){
                        this_success_order_increment_id 		=  emits[i].success_order_increment_id;
                    }else{
                        this_success_order_increment_id = this_success_order_increment_id.concat(emits[i].success_order_increment_id);
                    }
                }
                if( emits[i].success_order_info){
                    if(!this_success_order_info){
                        this_success_order_info 		=  emits[i].success_order_info;
                    }else{
                        this_success_order_info = this_success_order_info.concat(emits[i].success_order_info);
                    }
                }
                if( emits[i].fail_order_info){
                    if(!this_fail_order_info){
                        this_fail_order_info 		=  emits[i].fail_order_info;
                    }else{
                        this_fail_order_info = this_fail_order_info.concat(emits[i].fail_order_info);
                    }
                }
                if(emits[i].register_count){
                    this_register_count			+= emits[i].register_count;
                }
                if(emits[i].login_count){
                    this_login_count			+= emits[i].login_count;
                }
                if(emits[i].category_count){
                    this_category_count			+= emits[i].category_count;
                }
                if(emits[i].sku_count){
                    this_sku_count			+= emits[i].sku_count;
                }
                if(emits[i].search_count){
                    this_search_count			+= emits[i].search_count;
                }
                
                
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
                if(emits[i].order_no_count){
                    this_order_no_count += emits[i].order_no_count;
                }
                if(emits[i].order_amount){
                    this_order_amount 			+= emits[i].order_amount;
                }
                if(emits[i].success_order_amount){
                    this_success_order_amount 	+= emits[i].success_order_amount;
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
                if(emits[i].pv){
                    this_pv 			+= emits[i].pv;
                }
                if(emits[i].uv){
                    this_uv 			+= emits[i].uv;
                }
                if(emits[i].ip_count){
                    this_ip_count 			+= emits[i].ip_count;
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
               
                fid: this_fid,
                fec_edm: this_fec_edm,
                fec_market_group: this_fec_market_group,
                fec_content: this_fec_content,
                fec_source: this_fec_source,
                fec_design: this_fec_design,
                fec_campaign: this_fec_campaign,
                fec_medium: this_fec_medium,
                first_referrer_domain: this_first_referrer_domain,
                
                sku_visit_info: this_sku_visit_info,
                category_visit_info: this_category_visit_info,
                search_visit_info: this_search_visit_info,
                cart_sku_info: this_cart_sku_info,
                success_order_sku_info: this_success_order_sku_info,
                order_sku_info: this_order_sku_info,
                
                order_increment_id: this_order_increment_id,
                fail_order_increment_id: this_fail_order_increment_id,
                success_order_increment_id: this_success_order_increment_id,
                success_order_info: this_success_order_info,
                fail_order_info: this_fail_order_info,
                
                register_count: this_register_count,
                login_count: this_login_count,
                category_count: this_category_count,
                sku_count: this_sku_count,
                search_count: this_search_count,  
                
                devide:this_devide,
                country_code:this_country_code,
                ip_count:this_ip_count,
                operate:this_operate,
                fec_app:this_fec_app,
                is_return:this_is_return,
                first_page:this_first_page,
                resolution:this_resolution,
                color_depth:this_color_depth,
                language:this_language,
                cart_count:this_cart_count,
                order_count:this_order_count,	
                success_order_count:this_success_order_count,
                success_order_no_count:this_success_order_no_count,
                order_no_count:this_order_no_count,
                order_amount:this_order_amount,
                success_order_amount:this_success_order_amount,
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
            // 订单支付率
            order_no_count = reducedVal.order_no_count;
            if(order_no_count){
                this_success_order_no_count = reducedVal.success_order_no_count;
                this_order_payment_rate = this_success_order_no_count/order_no_count;
                this_order_payment_rate = this_order_payment_rate*10000;
                this_order_payment_rate = Math.ceil(this_order_payment_rate);
                this_order_payment_rate = this_order_payment_rate/10000;
            }else{
                this_order_payment_rate = 0;
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
            
            // success_order_c_success_uv_rate
            // 平均订单金额(客单价) = 订单总金额/订单数
            success_order_no_count = reducedVal.success_order_no_count;
            success_order_amount = reducedVal.success_order_amount;
            if(success_order_no_count){
                success_order_c_success_uv_rate = (success_order_amount/success_order_no_count); 
                success_order_c_success_uv_rate = success_order_c_success_uv_rate*10000;
                success_order_c_success_uv_rate = Math.ceil(success_order_c_success_uv_rate);
                success_order_c_success_uv_rate = success_order_c_success_uv_rate/10000;
            }else{
                success_order_c_success_uv_rate = 0;
            }
            reducedVal.success_order_c_success_uv_rate = success_order_c_success_uv_rate;
            
            // 广告订单平均金额 =   订单总金额/所有uv数
            if(uv){
                success_order_c_all_uv_rate = (success_order_amount/uv); 
                success_order_c_all_uv_rate = success_order_c_all_uv_rate*10000;
                success_order_c_all_uv_rate = Math.ceil(success_order_c_all_uv_rate);
                success_order_c_all_uv_rate = success_order_c_all_uv_rate/10000;
            }else{
                success_order_c_all_uv_rate = 0;
            }
            reducedVal.success_order_c_all_uv_rate = success_order_c_all_uv_rate;
            
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
            reducedVal.jump_out_rate     = this_jump_out_rate;
            reducedVal.drop_out_rate     = this_drop_out_rate;
            reducedVal.sku_sale_rate     = this_sku_sale_rate;
            reducedVal.order_payment_rate= this_order_payment_rate;
            //reducedVal.country_code    = this_country_code;
            //reducedVal.browser_name    = this_browser_name;
            //reducedVal.operate 		 = this_operate;
            reducedVal.pv_rate		     = this_pv_rate;
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
    esAdvertiseEdmTypeName :=  helper.GetEsAdvertiseEdmTypeName()
    esIndexName := helper.GetEsIndexNameByType(esAdvertiseEdmTypeName)
    // es index 的type mapping
    esAdvertiseEdmTypeMapping := helper.GetEsAdvertiseEdmTypeMapping()
    // 删除index，如果mapping建立的不正确，可以执行下面的语句删除重建mapping
    //err = esdb.DeleteIndex(esIndexName)
    //if err != nil {
    //    return err
    //}
    // 初始化mapping
    err = esdb.InitMapping(esIndexName, esAdvertiseEdmTypeName, esAdvertiseEdmTypeMapping)
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
            var advertiseEdms []model.AdvertiseEdm
            coll.Find(nil).Skip(i*pageNum).Limit(numPerPage).All(&advertiseEdms)
            log.Println("advertiseEdms length:")
            log.Println(len(advertiseEdms))
            
            /* 这个代码是upsert单行数据
            for j:=0; j<len(advertiseEdms); j++ {
                AdvertiseEdm := advertiseEdms[j]
                advertiseEdmValue := advertiseEdm.Value
                // advertiseEdmValue.Devide = nil
                // advertiseEdmValue.CountryCode = nil
                ///advertiseEdmValue.Operate = nil
                log.Println("ID_:" + advertiseEdm.Id_)
                advertiseEdmValue.Id = advertiseEdm.Id_
                err := esdb.UpsertType(esIndexName, esAdvertiseEdmTypeName, advertiseEdm.Id_, advertiseEdmValue)
                
                if err != nil {
                    log.Println("11111" + err.Error())
                    return err
                }
            }
            */
            if len(advertiseEdms) > 0 {
                // 使用bulk的方式，将数据批量插入到elasticSearch
                bulkRequest, err := esdb.Bulk()
                if err != nil {
                    log.Println("444" + err.Error())
                    return err
                }
                for j:=0; j<len(advertiseEdms); j++ {
                    advertiseEdm := advertiseEdms[j]
                    advertiseEdmValue := advertiseEdm.Value
                    advertiseEdmValue.Id = advertiseEdm.Id_
                    log.Println("888")
                    log.Println(esIndexName)
                    log.Println(esAdvertiseEdmTypeName)
                    log.Println(advertiseEdm.Id_)
                    log.Println(advertiseEdmValue)
                    req := esdb.BulkUpsertTypeDoc(esIndexName, esAdvertiseEdmTypeName, advertiseEdm.Id_, advertiseEdmValue)
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