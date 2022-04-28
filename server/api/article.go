package api

import (
	"fmt"
	"myapp/global"
	"myapp/model"
	"myapp/model/common/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ArticleApi struct {}

func (a *ArticleApi) Create(c *gin.Context) {
	var article model.ArticleModel
	c.ShouldBindJSON(&article)

	article.CreateTime = time.Now().Format("2006/1/2 15:04:05")
	article.UUID = uuid.New()

	err := global.GB_DB.Create(&article).Error
	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}