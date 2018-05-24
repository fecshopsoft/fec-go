package customer

import(
    // "github.com/fecshopsoft/fec-go/shell/whole"
    // "github.com/fecshopsoft/fec-go/util"
    "github.com/fecshopsoft/fec-go/db/mongodb"
    "github.com/fecshopsoft/fec-go/shell/customerModel"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "github.com/fecshopsoft/fec-go/helper"
    
)

// 将emailCollName 中count 大于2的数据，合并
func CustomerMergeByEmail(traceDbName string, traceCollName string, customerDbName string, customerCollName string, emailCollName string, website_id string) error {
    var err error
    err = mongodb.MDC(customerDbName, emailCollName, func(coll *mgo.Collection) error {
        var customerEmails []customerModel.UuidCustomerEmail
        _ = coll.Find(bson.M{"value.count": bson.M{"$gte":2}}).All(&customerEmails)
        for i:=0; i<len(customerEmails); i++ {
            customerEmail := customerEmails[i]
            email := customerEmail.Value.Email
            if email == "" {
                continue
            }
            // 数据合并
            var deleteCustomerIds []string
            var customerUuids []string
            var customerEmails []string
            var primaryCustomerId string
            var customers []customerModel.UuidCustomer
            err = mongodb.MDC(customerDbName, customerCollName, func(coll *mgo.Collection) error {
                _ = coll.Find(bson.M{"emails": email}).Sort("+_id").All(&customers)
                if len(customers) > 0 {
                    
                    primaryCustomerId = customers[0].CustomerId
                    for i:=0; i<len(customers); i++ {
                        customer := customers[i]
                        uuids := customer.Uuids
                        emails := customer.Emails
                        customerId := customer.CustomerId
                        if primaryCustomerId != customerId {
                            deleteCustomerIds = append(deleteCustomerIds, customerId)
                        }
                        if len(uuids) > 0 {
                            customerUuids = helper.ArrayMergeAndUnique(customerUuids, uuids)
                        }
                        if len(emails) > 0 {
                            customerEmails = helper.ArrayMergeAndUnique(customerEmails, emails)
                        }
                    }
                    // 更新primary customer ，然后，将其他的删除
                    selector := bson.M{"customer_id": primaryCustomerId}
                    updateData := bson.M{"$set": bson.M{"uuids": customerUuids, "emails": customerEmails}}
                    err = coll.Update(selector, updateData)
                    deleteSelector := bson.M{"customer_id": bson.M{"$in": deleteCustomerIds}}
                    err = coll.Remove(deleteSelector)
                }
                return err
            })
            // trace表，更新trace表的数据。
            err = mongodb.MDC(traceDbName, traceCollName, func(coll *mgo.Collection) error {
                selector := bson.M{"uuid": bson.M{"$in": customerUuids}}
                updateData := bson.M{"$set": bson.M{"customer_id": primaryCustomerId}}
                err = coll.Update(selector, updateData)
                return err
            })
            
        }
        return err
    })

    return err
}

// uuid 统计计算部分
func EmailMapReduct(dbName string, collName string, outCollName string, website_id string) error {
    var err error
    mapStr := `
        function() {  
            emails = this.emails;
            if(emails){
                for(x in emails){
                    email = emails[x];
                    if(email){
                        if( 
                            email != "1001@qq.com"
                        &&	email != "3001258674@qq.com"
                        &&  email != "sunlight426@126.com"
                        
                        ){
                            emit(email,{
                                email:email,
                                count:1
                            });
                        }
                    }
                }
                
            }
        }
    `
    
    reduceStr := `
        function(key,emits){
            this_email	 		= null;
            this_count = 0;
            for(var i in emits){
                if(emits[i].email){
                    this_email = emits[i].email;
                }
                if(emits[i].count){
                    this_count += emits[i].count;
                }
            }
            
            return {	
                email:this_email,
                count:this_count
            };
        }
    `
    
    finalizeStr := `
        function (key, reducedVal) {
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
    
    return err
}
