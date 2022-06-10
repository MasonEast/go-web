package api

import (
	"context"
	"fmt"
	"myapp/global"
	"myapp/kafka"
	"myapp/model"
	"myapp/model/common/response"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olivere/elastic/v7"
)

type ApointApi struct {}

var producer *kafka.Producer


func init () {
	// 初始化kafka
	producer = new(kafka.Producer)
	producer.InitProducer()
}

// 创建埋点
func (a *ApointApi) Send(c *gin.Context) {
	var apoint model.ApointModel
	c.ShouldBindJSON(&apoint)

	apoint.CreateTime = time.Now().Format(global.GB_Time_Format)
	apoint.UUID = uuid.New()

	producer.SendMessage(apoint.Data)

	err := global.GB_DB.Create(&apoint).Error
	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}


// 查询数据
func (a *ApointApi) Search(c *gin.Context) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		fmt.Println("create elastic client err:",err)
	}
	// res, err := client.Get().Index("logstash-2022.06.10").Id("tlOQS4EBh3aiP3pGr_qn").Do(context.Background())

	rangeQuery := elastic.NewRangeQuery("@timestamp").Gt("2022-06-09").Lt("now")
	res, err := client.Search().
	Index("logstash-2022.06.10").
	Query(rangeQuery).
	// From(0).
	// Size(10).
	Do(context.Background())
	
	if err != nil {
		fmt.Println("get elastic client err:",err)
	}

	response.OkWithData(res, c)

}