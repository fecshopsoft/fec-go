package main

import(
    "github.com/fecshopsoft/fec-go/router"
)

func main() { 
    listenIp := "120.24.37.249:3000"
    router.Listen(listenIp);
}