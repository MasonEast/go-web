package api

import (
	"fmt"
	"myapp/global"
	"myapp/model"
	"myapp/model/common/request"
	"myapp/model/common/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ArticleApi struct {}

// 创建文章
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

// 更新文章
func (a *ArticleApi) Update(c *gin.Context) {
	var article model.ArticleModel
	c.ShouldBindJSON(&article)

	err := global.GB_DB.Where("uuid = ?", article.UUID).Save(&article).Error
	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 删除文章
func (a *ArticleApi) Delete(c *gin.Context) {
	var article model.ArticleModel
	c.ShouldBindJSON(&article)

	err := global.GB_DB.Where("uuid = ?", article.UUID).Delete(&article).Error
	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 获取文章列表
func (a *ArticleApi) ArticleList(c *gin.Context) {
	var params request.PageInfo
	var list []model.ArticleModel
	var total int64

	c.ShouldBindQuery(&params)

	// 获取文章总条数
	global.GB_DB.Model(&model.ArticleModel{}).Count(&total)

	// 获取分页数据
	err := global.GB_DB.Limit(params.PageSize).Offset(params.PageSize * (params.Page - 1)).
				 Find(&list).Error

	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "获取文章列表失败", c)
		return
	}

	response.OkWithData(response.PageResult{
		List: list,
		Total: total,
		Page: params.Page,
		PageSize: params.PageSize,
	}, c)
}