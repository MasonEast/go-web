package api

import (
	"errors"
	"fmt"
	"myapp/global"
	"myapp/model"
	"myapp/model/common/response"
	"myapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicApi struct {}

func (p *PublicApi) Register(c *gin.Context){
	var user model.UserModel
	c.ShouldBindJSON(&user)

	// 判断用户名是否注册
	if !errors.Is(global.GB_DB.Where("username = ?", user.Username).First(&user).Error, gorm.ErrRecordNotFound) { 
		response.FailWithMessage("用户名已注册", c)
		return 
	}
	// 否则 附加uuid 密码md5简单加密 注册
	user.Password = utils.MD5V([]byte(user.Password))
	user.UUID = uuid.New()

	err := global.GB_DB.Create(&user).Error

	if err != nil {
		fmt.Println(err)
		response.FailWithDetailed(err, "创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}