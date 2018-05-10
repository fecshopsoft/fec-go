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

// Category 统计计算部分
func CategoryMapReduct(dbName string, collName string, outCollName string, website_id string) error {
    var err error
    
    mapStr := `
        function() {  
            website_id = "` + website_id + `";
            category = this.category ? this.category : null;
            if (category) {
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
                
                // 
                is_return = 0;
                uv = 0;
                ip_count = 0;
                if(this.uuid_first_category == 1){
                    uv = 1;
                    is_return 		= this.is_return ? this.is_return : 0;
                }
                if(this.ip_first_category == 1){
                    ip_count = 1;
                }
                service_date_str = this.service_date_str ? this.service_date_str : null;
                
                is_return = Number(is_return);
                is_return = isNaN(is_return) ? 0 : is_return
                first_page = Number(first_page);
                first_page = isNaN(first_page) ? 0 : first_page
                
                emit(category+"_"+service_date_str+"_"+website_id,{
                    category: category,
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
    `
    
    reduceStr := `
        function(key,emits){
            this_category                = 0;
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
            
            for(var i in emits){
                
                if(emits[i].category){
                    this_category = emits[i].category;
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
                
                if(emits[i].is_return){
                    this_is_return 			+= emits[i].is_return;
                }
                if(emits[i].first_page){
                    this_first_page 		+= emits[i].first_page;
                }
            }
            
            return {	
                category: this_category,
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
                ip_count:this_ip_count,
                operate:this_operate,
                fec_app:this_fec_app,
                is_return:this_is_return,
                first_page:this_first_page,
                resolution:this_resolution,
                color_depth:this_color_depth,
                language:this_language,
                
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
            reducedVal.jump_out_rate     = this_jump_out_rate;
            reducedVal.drop_out_rate     = this_drop_out_rate;
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
    esWholeCategoryTypeName :=  helper.GetEsWholeCategoryTypeName()
    esIndexName := helper.GetEsIndexNameByType(esWholeCategoryTypeName)
    // es index 的type mapping
    esWholeCategoryTypeMapping := helper.GetEsWholeCategoryTypeMapping()
    // 删除index，如果mapping建立的不正确，可以执行下面的语句删除重建mapping
    //err = esdb.DeleteIndex(esIndexName)
    //if err != nil {
    //    return err
    //}
    // 初始化mapping
    err = esdb.InitMapping(esIndexName, esWholeCategoryTypeName, esWholeCategoryTypeMapping)
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
            var WholeCategorys []model.WholeCategory
            coll.Find(nil).Skip(i*pageNum).Limit(numPerPage).All(&WholeCategorys)
            log.Println("WholeCategorys length:")
            log.Println(len(WholeCategorys))
            
            /* 这个代码是upsert单行数据
            for j:=0; j<len(WholeCategorys); j++ {
                wholeBrowser := WholeCategorys[j]
                wholeBrowserValue := wholeBrowser.Value
                // wholeBrowserValue.Devide = nil
                // wholeBrowserValue.CountryCode = nil
                ///wholeBrowserValue.Operate = nil
                log.Println("ID_:" + wholeBrowser.Id_)
                wholeBrowserValue.Id = wholeBrowser.Id_
                err := esdb.UpsertType(esIndexName, esWholeCategoryTypeName, wholeBrowser.Id_, wholeBrowserValue)
                
                if err != nil {
                    log.Println("11111" + err.Error())
                    return err
                }
            }
            */
            if len(WholeCategorys) > 0 {
                // 使用bulk的方式，将数据批量插入到elasticSearch
                bulkRequest, err := esdb.Bulk()
                if err != nil {
                    log.Println("444" + err.Error())
                    return err
                }
                for j:=0; j<len(WholeCategorys); j++ {
                    WholeCategory := WholeCategorys[j]
                    WholeCategoryValue := WholeCategory.Value
                    WholeCategoryValue.Id = WholeCategory.Id_
                    log.Println("888")
                    log.Println(esIndexName)
                    log.Println(esWholeCategoryTypeName)
                    log.Println(WholeCategory.Id_)
                    log.Println(WholeCategoryValue)
                    req := esdb.BulkUpsertTypeDoc(esIndexName, esWholeCategoryTypeName, WholeCategory.Id_, WholeCategoryValue)
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
