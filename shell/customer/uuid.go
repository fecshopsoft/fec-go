package customer

import(
    // "github.com/fecshopsoft/fec-go/shell/whole"
    // "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/fecshopsoft/fec-go/db/esdb"
    "github.com/fecshopsoft/fec-go/helper"
    model "github.com/fecshopsoft/fec-go/shell/customerModel"
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
            
            sku 			= this.sku;
            if(sku){
                b= {};
                b[sku] = 1;
                sku 	= b;
            }else{
                sku = null;
            }
            category 			= this.category;
            if(category){
                b= {};
                b[category] = 1;
                category 	= b;
            }else{
                category = null;
            }
            
            search 			= this.search;
            if(search && search.text){
                b= {};
                b[search.text] = search.result_qty;
                search 	= b;
            }else{
                search = null;
            }
            
            if (this.cart) {
                cart		= this.cart.length > 0 ? this.cart : null;  // last cart page
            } else {
                cart = null
            }
            
            if (this.order) {
                order		= this.order.invoice ? [this.order] : null;
            }else {
                order = null
            }
            
            customer_email		= [];
            if(login_email){
                customer_email.push(login_email);
            }
            if(register_email){
                customer_email.push(register_email);
            }
            
            // user_agent= this.user_agent ? this.user_agent : null;
            // service_timestamp= this.service_timestamp ? this.service_timestamp : null;
           

            ip 			= this.ip;
            if(ip){
                b= {};
                b[ip] = 1;
                ip 	= b;
            }else{
                ip = null;
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
            if(v_order = this.order && this.order.invoice){
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
            
            
            
            sku_cart = {};
            sku_order = {};
            sku_order_success = {};
            thiscart = this.cart;
            if(thiscart){
                ii = 0 ;
                for(x in thiscart){
                    ii++;
                    c_one = thiscart[x];
                    $c_sku = c_one['sku'];
                    if($c_sku){
                        sku_cart[$c_sku] = c_one['qty'];
                    }
                }
                if (ii == 0) {
                    sku_cart = null;
                }
            }
            thisorder = this.order;
            if(thisorder && this.order.invoice){ 
                c_products = thisorder.products;
                o_payment_status = thisorder.payment_status;
                ii = 0 ;
                jj = 0;
                for(x in c_products){
                    o_product = c_products[x];
                    o_sku = o_product['sku'];
                    if(o_sku){
                        if(o_payment_status == 'payment_confirmed'){
                            ii++;
                            sku_order_success[o_sku] = o_product['qty'];
                        }else{
                            jj++;
                            sku_order[o_sku] = o_product['qty'];
                        }
                    }
                }
                if (ii == 0) {
                    sku_order_success = null;
                }
                if (jj == 0) {
                    sku_order = null;
                }
            }
            
            fid 	= this.fid;
            if(fid){
                b= {};
                b[fid] = 1;
                fid 	= b;
            }else{
                fid = null;
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
            
            fec_campaign 	= this.fec_campaign;
            if(fec_campaign){
                b= {};
                b[fec_campaign] = 1;
                fec_campaign 	= b;
            }else{
                fec_campaign = null;
            }
            
            
            fec_source 	= this.fec_source;
            if(fec_source){
                b= {};
                b[fec_source] = 1;
                fec_source 	= b;
            }else{
                fec_source = null;
            }
            
            fec_medium 	= this.fec_medium;
            if(fec_medium){
                b= {};
                b[fec_medium] = 1;
                fec_medium 	= b;
            }else{
                fec_medium = null;
            }
            
            fec_design 	= this.fec_design;
            if(fec_design){
                b= {};
                b[fec_design] = 1;
                fec_design 	= b;
            }else{
                fec_design = null;
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
                    
                    // 改变成数组
                    fid:fid ? fid : null,  
                    fec_content:fec_content ? fec_content : null,
                    fec_market_group:fec_market_group ? fec_market_group : null,
                    fec_campaign:fec_campaign ? fec_campaign : null,
                    fec_source:fec_source ? fec_source : null,
                    fec_medium:fec_medium ? fec_medium : null,
                    fec_design:fec_design ? fec_design : null,
                    
                    // 新添加
                    ip: ip,
                    browser_name: browser_name,
                    devide: devide,
                    country_code: country_code,
                    operate: operate,
                    fec_app: fec_app,
                    resolution: resolution,
                    color_depth: color_depth,
                    language: language,
            
                
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
                    
                    refer_url:refer_url,
                    first_referrer_domain:first_referrer_domain,
                    
                    is_return:is_return,
                    
                    first_page_url:first_page_url,
                    out_page:out_page,
                    service_date_str:service_date_str,
                    
                    
                    data:[{
                        _id:this._id ? this._id.valueOf() : null,
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
                        
                        // 新添加
                        fec_app:this.fec_app ?  this.fec_app : null,
                        language:this.language ?  this.language : null,
                    
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
    
            //sku_cart
            //sku_order_success
            //sku_order
            //sku
            //category
  
    reduceStr := `
        function(key,emits){
            this_uuid 				= null;
            
            this_ip		        = {};
            this_browser_name	= {};
            this_devide	        = {};
            this_country_code   = {};
            this_operate	    = {};
            this_fec_app        = {};
            this_resolution     = {};
            this_color_depth    = {};
            this_language       = {};
            
            this_customer_email		= [];
            
            this_fid 				        = {};
            this_fec_content 				= {};
            this_fec_market_group 			= {};
            this_fec_campaign 				= {};
            this_fec_source 				= {};
            this_fec_medium 				= {};
            this_fec_design 				= {};

            this_customer_id		= null;
            this_identify 			= [];
            this_pv 				= 0;
            this_stay_seconds		= 0;
            this_service_date_str 	= null;
            this_register_email		= null;
            
            this_login_email		= null;
            
            this_sku				= {};
            this_category		    = {};
            this_sku_cart			= {};
            this_sku_order			= {};
            this_sku_order_success	= {};
            
            
            this_search		= {};
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

            this_refer_url	= null;
            this_first_referrer_domain	= null;
            this_is_return	= null;
            this_first_page_url	= null;
            this_out_page   = null;
            this_data = [];
            
            for(var i in emits){
                
                if(emits[i].uuid){
                    this_uuid 			=  emits[i].uuid;
                }
                
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
                
                if(emits[i].fec_source){
                    fec_source = emits[i].fec_source;
                    for(brower_ne in fec_source){
                        count = fec_source[brower_ne];
                        if(!this_fec_source[brower_ne]){
                            this_fec_source[brower_ne] = count;
                        }else{
                            this_fec_source[brower_ne] += count;
                        }
                    }
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
                
                if(emits[i].fid){
                    fid = emits[i].fid;
                    for(brower_ne in fid){
                        count = fid[brower_ne];
                        if(!this_fid[brower_ne]){
                            this_fid[brower_ne] = count;
                        }else{
                            this_fid[brower_ne] += count;
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
                
                if(emits[i].fec_campaign){
                    fec_campaign = emits[i].fec_campaign;
                    for(brower_ne in fec_campaign){
                        count = fec_campaign[brower_ne];
                        if(!this_fec_campaign[brower_ne]){
                            this_fec_campaign[brower_ne] = count;
                        }else{
                            this_fec_campaign[brower_ne] += count;
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
                
                if(emits[i].search){
                    search = emits[i].search;
                    for(brower_ne in search){
                        count = search[brower_ne];
                        if(!this_search[brower_ne]){
                            this_search[brower_ne] = count;
                        }else{
                            this_search[brower_ne] += count;
                        }
                    }
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
                if(emits[i].sku){
                    sku = emits[i].sku;
                    for(brower_ne in sku){
                        
                        count = sku[brower_ne];
                        if(!this_sku[brower_ne]){
                            this_sku[brower_ne] = count;
                        }else{
                            this_sku[brower_ne] += count;
                        }
                    }
                }
                
                if(emits[i].category){
                    category = emits[i].category;
                    for(brower_ne in category){
                        
                        count = category[brower_ne];
                        if(!this_category[brower_ne]){
                            this_category[brower_ne] = count;
                        }else{
                            this_category[brower_ne] += count;
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
                if(emits[i].sku_cart){
                    sku_cart = emits[i].sku_cart;
                    for(brower_ne in sku_cart){
                        count = sku_cart[brower_ne];
                        if(!this_sku_cart[brower_ne]){
                            this_sku_cart[brower_ne] = count;
                        }else{
                            this_sku_cart[brower_ne] += count;
                        }
                    }
                }
                
                // this_sku_order			= [];
                if(emits[i].sku_order){
                    sku_order = emits[i].sku_order;
                    for(brower_ne in sku_order){
                        count = sku_order[brower_ne];
                        if(!this_sku_order[brower_ne]){
                            this_sku_order[brower_ne] = count;
                        }else{
                            this_sku_order[brower_ne] += count;
                        }
                    }
                }
                
                // this_sku_order_success	= [];
                if(emits[i].sku_order_success){
                    sku_order_success = emits[i].sku_order_success;
                    for(brower_ne in sku_order_success){
                        count = sku_order_success[brower_ne];
                        if(!this_sku_order_success[brower_ne]){
                            this_sku_order_success[brower_ne] = count;
                        }else{
                            this_sku_order_success[brower_ne] += count;
                        }
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
                    ip = emits[i].ip;
                    for(brower_ne in ip){
                        
                        count = ip[brower_ne];
                        if(!this_ip[brower_ne]){
                            this_ip[brower_ne] = count;
                        }else{
                            this_ip[brower_ne] += count;
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
                fid:this_fid,
                
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
                language: this_language,
                sku_cart:this_sku_cart,
                sku_order:this_sku_order,
                sku_order_success:this_sku_order_success,
                
                ip:this_ip,
                browser_name:this_browser_name,
                devide:this_devide,
                country_code:this_country_code,
                operate:this_operate,
                fec_app: this_fec_app,
                resolution:this_resolution,
                color_depth:this_color_depth,
                
                
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
                
                
                

                
                
                
                
                
                refer_url:this_refer_url,
                first_referrer_domain:this_first_referrer_domain,
                is_return:this_is_return,
                first_page_url:this_first_page_url,
                out_page:this_out_page,
                
                data:this_data
            };
        }
    `
    
    finalizeStr := `
        function (key, reducedVal) {
            reducedVal.website_id        = "` + website_id + `"
            
            
            language = reducedVal.language;
            language_main = null;
            max_count = 0;
            for(language_name in language){
                language_count = language[language_name];
                if(!language_main){
                    max_count = language_count;
                    language_main = language_name;
                }else{
                    if(max_count < language_count){
                        language_main = language_name;
                        max_count = language_count;
                    }
                }
            }
            reducedVal.language_main = language_main;
            
            
            color_depth = reducedVal.color_depth;
            color_depth_main = null;
            max_count = 0;
            for(color_depth_name in color_depth){
                color_depth_count = color_depth[color_depth_name];
                if(!color_depth_main){
                    max_count = color_depth_count;
                    color_depth_main = color_depth_name;
                }else{
                    if(max_count < color_depth_count){
                        color_depth_main = color_depth_name;
                        max_count = color_depth_count;
                    }
                }
            }
            reducedVal.color_depth_main = color_depth_main;
            
            
            resolution = reducedVal.resolution;
            resolution_main = null;
            max_count = 0;
            for(resolution_name in resolution){
                resolution_count = resolution[resolution_name];
                if(!resolution_main){
                    max_count = resolution_count;
                    resolution_main = resolution_name;
                }else{
                    if(max_count < resolution_count){
                        resolution_main = resolution_name;
                        max_count = resolution_count;
                    }
                }
            }
            reducedVal.resolution_main = resolution_main;
            
            fec_app = reducedVal.fec_app;
            fec_app_main = null;
            max_count = 0;
            for(fec_app_name in fec_app){
                fec_app_count = fec_app[fec_app_name];
                if(!fec_app_main){
                    max_count = fec_app_count;
                    fec_app_main = fec_app_name;
                }else{
                    if(max_count < fec_app_count){
                        fec_app_main = fec_app_name;
                        max_count = fec_app_count;
                    }
                }
            }
            reducedVal.fec_app_main = fec_app_main;
            
            
            operate = reducedVal.operate;
            operate_main = null;
            max_count = 0;
            for(operate_name in operate){
                operate_count = operate[operate_name];
                if(!operate_main){
                    max_count = operate_count;
                    operate_main = operate_name;
                }else{
                    if(max_count < operate_count){
                        operate_main = operate_name;
                        max_count = operate_count;
                    }
                }
            }
            reducedVal.operate_main = operate_main;
            
            country_code = reducedVal.country_code;
            country_code_main = null;
            max_count = 0;
            for(country_code_name in country_code){
                country_code_count = country_code[country_code_name];
                if(!country_code_main){
                    max_count = country_code_count;
                    country_code_main = country_code_name;
                }else{
                    if(max_count < country_code_count){
                        country_code_main = country_code_name;
                        max_count = country_code_count;
                    }
                }
            }
            reducedVal.country_code_main = country_code_main;
            
            
            devide = reducedVal.devide;
            devide_main = null;
            max_count = 0;
            for(devide_name in devide){
                devide_count = devide[devide_name];
                if(!devide_main){
                    max_count = devide_count;
                    devide_main = devide_name;
                }else{
                    if(max_count < devide_count){
                        devide_main = devide_name;
                        max_count = devide_count;
                    }
                }
            }
            reducedVal.devide_main = devide_main;
            
            browser_name = reducedVal.browser_name;
            browser_name_main = null;
            max_count = 0;
            for(browser_name_name in browser_name){
                browser_name_count = browser_name[browser_name_name];
                if(!browser_name_main){
                    max_count = browser_name_count;
                    browser_name_main = browser_name_name;
                }else{
                    if(max_count < browser_name_count){
                        browser_name_main = browser_name_name;
                        max_count = browser_name_count;
                    }
                }
            }
            reducedVal.browser_name_main = browser_name_main;
            
            ip = reducedVal.ip;
            ip_main = null;
            max_count = 0;
            for(ip_name in ip){
                ip_count = ip[ip_name];
                if(!ip_main){
                    max_count = ip_count;
                    ip_main = ip_name;
                }else{
                    if(max_count < ip_count){
                        ip_main = ip_name;
                        max_count = ip_count;
                    }
                }
            }
            reducedVal.ip_main = ip_main;
            
            fec_design = reducedVal.fec_design;
            fec_design_main = null;
            max_count = 0;
            for(fec_design_name in fec_design){
                fec_design_count = fec_design[fec_design_name];
                if(!fec_design_main){
                    max_count = fec_design_count;
                    fec_design_main = fec_design_name;
                }else{
                    if(max_count < fec_design_count){
                        fec_design_main = fec_design_name;
                        max_count = fec_design_count;
                    }
                }
            }
            reducedVal.fec_design_main = fec_design_main;
            
            fec_medium = reducedVal.fec_medium;
            fec_medium_main = null;
            max_count = 0;
            for(fec_medium_name in fec_medium){
                fec_medium_count = fec_medium[fec_medium_name];
                if(!fec_medium_main){
                    max_count = fec_medium_count;
                    fec_medium_main = fec_medium_name;
                }else{
                    if(max_count < fec_medium_count){
                        fec_medium_main = fec_medium_name;
                        max_count = fec_medium_count;
                    }
                }
            }
            reducedVal.fec_medium_main = fec_medium_main;
            
            
            fec_source = reducedVal.fec_source;
            fec_source_main = null;
            max_count = 0;
            for(fec_source_name in fec_source){
                fec_source_count = fec_source[fec_source_name];
                if(!fec_source_main){
                    max_count = fec_source_count;
                    fec_source_main = fec_source_name;
                }else{
                    if(max_count < fec_source_count){
                        fec_source_main = fec_source_name;
                        max_count = fec_source_count;
                    }
                }
            }
            reducedVal.fec_source_main = fec_source_main;
            
            
            fec_campaign = reducedVal.fec_campaign;
            fec_campaign_main = null;
            max_count = 0;
            for(fec_campaign_name in fec_campaign){
                fec_campaign_count = fec_campaign[fec_campaign_name];
                if(!fec_campaign_main){
                    max_count = fec_campaign_count;
                    fec_campaign_main = fec_campaign_name;
                }else{
                    if(max_count < fec_campaign_count){
                        fec_campaign_main = fec_campaign_name;
                        max_count = fec_campaign_count;
                    }
                }
            }
            reducedVal.fec_campaign_main = fec_campaign_main;
            
            
            fec_market_group = reducedVal.fec_market_group;
            fec_market_group_main = null;
            max_count = 0;
            for(fec_market_group_name in fec_market_group){
                fec_market_group_count = fec_market_group[fec_market_group_name];
                if(!fec_market_group_main){
                    max_count = fec_market_group_count;
                    fec_market_group_main = fec_market_group_name;
                }else{
                    if(max_count < fec_market_group_count){
                        fec_market_group_main = fec_market_group_name;
                        max_count = fec_market_group_count;
                    }
                }
            }
            reducedVal.fec_market_group_main = fec_market_group_main;
            
            
            fec_content = reducedVal.fec_content;
            fec_content_main = null;
            max_count = 0;
            for(fec_content_name in fec_content){
                fec_content_count = fec_content[fec_content_name];
                if(!fec_content_main){
                    max_count = fec_content_count;
                    fec_content_main = fec_content_name;
                }else{
                    if(max_count < fec_content_count){
                        fec_content_main = fec_content_name;
                        max_count = fec_content_count;
                    }
                }
            }
            reducedVal.fec_content_main = fec_content_main;
            
            
            fid = reducedVal.fid;
            fid_main = null;
            max_count = 0;
            for(fid_name in fid){
                fid_count = fid[fid_name];
                if(!fid_main){
                    max_count = fid_count;
                    fid_main = fid_name;
                }else{
                    if(max_count < fid_count){
                        fid_main = fid_name;
                        max_count = fid_count;
                    }
                }
            }
            reducedVal.fid_main = fid_main;
            
            
            
            
            
            
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
                    //log.Println(customerUuid.Id_)
                    //log.Println(customerUuidValue)
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