package customerModel

import(
    "github.com/globalsign/mgo/bson"
)

// type MapStringInt64 map[string]int64
type UuidCustomer struct{
    Id_  bson.ObjectId `form:"_id" json:"_id" bson:"_id"` 
    CustomerId string `form:"customer_id" json:"customer_id" bson:"customer_id"`
    Uuids []string `form:"uuids" json:"uuids" bson:"uuids"`
    Emails []string `form:"emails" json:"emails" bson:"emails"`
    UpdatedAt int64 `form:"updated_at" json:"updated_at" bson:"updated_at"`
}

type UuidCustomerEmail struct{
    Id_ string `form:"_id" json:"_id" bson:"_id"`
    Value UuidCustomerEmailValue `form:"value" json:"value" bson:"value"`
}

type UuidCustomerEmailValue struct{
    Email string `form:"email" json:"email" bson:"email"`
    Count string `form:"count" json:"count" bson:"count"`
}