package global

import (
	"fmt"
	"io/ioutil"
	"myapp/config"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"golang.org/x/sync/singleflight"

)

var (
	GB_CONFIG config.Config
	GB_DB *gorm.DB
	GB_Concurrency_Control = &singleflight.Group{}
	GB_Time_Format = "2006/1/2 15:04:05"
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