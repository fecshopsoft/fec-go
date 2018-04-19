package initialization

import (
	"log"
	"github.com/fecshopsoft/fec-go/config"
	"os"
	"path/filepath"
)
// 初始化log
func InitGlobalLog() {
	if "false" == config.Get("output_log") {
		log.SetOutput(os.Stdout)
        return
	}
	globalLogUrl := config.Get("global_log")
	if globalLogUrl == "" {
		globalLogUrl = "logs/global.log"
	}
	path := filepath.Dir(globalLogUrl)
	os.MkdirAll(path, 0777)
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}



// 初始化 shell log
func InitShellLog() {
	if "false" == config.Get("shell_output_log") {
		log.SetOutput(os.Stdout)
        return
	}
	globalLogUrl := config.Get("shell_global_log")
	if globalLogUrl == "" {
		globalLogUrl = "logs/shell_global.log"
	}
	path := filepath.Dir(globalLogUrl)
	os.MkdirAll(path, 0777)
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}