package api

import (
	"fmt"
	"myapp/global"
	"myapp/model"
	"myapp/model/common/response"

	"github.com/gin-gonic/gin"
)

type AdminApi struct {}

func (a *AdminApi)Get(c *gin.Context){
	var api model.AdminModel
	c.ShouldBindJSON(&api)

	err := global.GB_DB.First(&api).Error

	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("获取失败", c)
	}

	response.OkWithData(api, c)
}