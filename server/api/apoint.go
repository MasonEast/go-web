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

type ApointApi struct {}

// 创建埋点
func (a *ApointApi) Send(c *gin.Context) {
	var apoint model.ApointModel
	c.ShouldBindJSON(&apoint)

	apoint.CreateTime = time.Now().Format(global.GB_Time_Format)
	apoint.UUID = uuid.New()

	err := global.GB_DB.Create(&apoint).Error
	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
