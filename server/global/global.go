package global

import (
	"fmt"
	"io/ioutil"
	"myapp/config"

	"golang.org/x/sync/singleflight"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

var (
	GB_CONFIG config.Config
	GB_DB *gorm.DB
	GB_Concurrency_Control = &singleflight.Group{}
	GB_Time_Format = "2006/1/2 15:04:05"

	TOPIC     = "looklook-log" //主题
	PARTITION = 0          //partition ID

	SERVER_LIST = "localhost:9094"

)

func init () {
	
	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Print(err)
	}

	//yaml文件内容影射到结构体中
	yaml.Unmarshal(config, &GB_CONFIG)
	fmt.Println(GB_CONFIG)

}