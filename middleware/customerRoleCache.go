package middleware

import(
	"sync"
    "log"
    customerH "github.com/fecshopsoft/fec-go/handler/customer"
)
// https://studygolang.com/articles/7973

var customerResourceCache CustomerResourceCache

type CustomerResourceCache struct {
	//count int
	// keys  []string
	resources  map[int64][]customerH.ResourceRole
	lock  sync.RWMutex
}


// 添加kv键值对
func (this *CustomerResourceCache) Set(customer_id int64, resources []customerH.ResourceRole) {
	log.Println("###################1")
    log.Println(customer_id)
    log.Println(resources)
    this.lock.Lock()
    if this.resources == nil {
        this.resources = make(map[int64][]customerH.ResourceRole)
    }
	this.resources[customer_id] = resources
	this.lock.Unlock()
}

// 由key检索value
func (this *CustomerResourceCache) Get(customer_id int64) ([]customerH.ResourceRole, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	resources, ok := this.resources[customer_id]
	return resources, ok
}


// 获取数据长度
/*
func (this *CustomerResourceCache) Count() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.count
}
*/



// 根据key排序，返回有序的vaule切片
/*
func (this *Cache) Values() []interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	vals := make([]interface{}, this.count)
	for i := 0; i < this.count; i++ {
		vals[i] = this.hash[this.keys[i]]
	}
	return vals
}
*/