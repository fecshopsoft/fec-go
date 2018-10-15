package config

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Mymap map[string]string
}

var myConfig *Config
var once sync.Once

func Get(key string) string {
	v, found := GetInstance().Mymap[key]
	if !found {
		return ""
	}
	return v
}

func GetInstance() *Config {
	once.Do(func() {
		myConfig = new(Config)
        	myConfig.InitConfig("/www/fec-go/etc/config.ini")
		//myConfig.InitConfig("resource/config")
		//myConfig.InitConfig("C:\\work\\Workspaces\\goWorkspace20161022\\src\\mimi\\djq\\resource\\config")
	})
	return myConfig
}



func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}
		c.Mymap[frist] = strings.TrimSpace(second)
	}
}