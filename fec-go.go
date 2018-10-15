package main
/**
 * 服务端入口部分
 * 1.初始化log输出文件
 * 2.监听ip
 *
 */
import(
	"github.com/fecshopsoft/fec-go/router"
	"github.com/fecshopsoft/fec-go/initialization"
	"github.com/fecshopsoft/fec-go/config"
	"log"
	"time"
	"os"
	"os/signal"
	"syscall"
	"github.com/fecshopsoft/fec-go/db/mysqldb"
)

func main() {
	// 初始化log输出，log.Println("---") 输出的内容将输出到globalLog文件里面
	log.Println("------start：" + time.Now().String())
	initialization.InitGlobalLog()
	log.Println("------start：" + time.Now().String())
	log.SetFlags(log.LstdFlags | log.Llongfile)
	SetupCloseHandler()
	listenIp := config.Get("httpHost")
	router.Listen(listenIp);

}

func VerifyRemote(){

}

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := mysqldb.CloseEngine()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("\r- Ctrl+C pressed in Terminal")
		log.Println("------close：" + time.Now().String())
		os.Exit(0)
	}()
}