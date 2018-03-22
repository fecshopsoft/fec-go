package main

import(
    "github.com/fecshopsoft/fec-go/router"
)

func main() { 

    log.Println("------start：" + time.Now().String())
	initialization.InitGlobalLog()
	log.Println("------start：" + time.Now().String())
    log.SetFlags(log.LstdFlags | log.Llongfile)
    
    listenIp := "120.24.37.249:3000"
    router.Listen(listenIp);
}