package main

import(
    "github.com/fecshopsoft/fec-go/initialization"
    "log"
    "time"
    "github.com/fecshopsoft/fec-go/shell"
)

func main() { 
    // 初始化log输出，log.Println("---") 输出的内容将输出到globalLog文件里面
    log.Println("------start：" + time.Now().String())
	initialization.InitShellLog()
	
    log.SetFlags(log.LstdFlags | log.Llongfile)
    shell.GoShell()
    // shell.TestEs()
    log.Println("------end：" + time.Now().String())
}