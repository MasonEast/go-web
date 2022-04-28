package api

import (
	"errors"
	"fmt"
	"myapp/global"
	"myapp/middleware"
	"myapp/model"
	"myapp/model/common/request"
	"myapp/model/common/response"
	"myapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicApi struct {}

// 用户注册
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

// 用户登录
func (p *PublicApi) Login(c *gin.Context) {
	var user model.UserModel
	c.ShouldBindJSON(&user)

	user.Password = utils.MD5V([]byte(user.Password))
	err := global.GB_DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error
	if err == nil {
		// 账号密码通过后，签发token
		p.tokenNext(c, user)
	} else {
		response.FailWithMessage("账号/密码错误", c)
	}
}

// 登录以后签发jwt
func (p *PublicApi) tokenNext(c *gin.Context, user model.UserModel) {
	j := &middleware.JWT{SigningKey: []byte(global.GB_CONFIG.JWT.SigningKey)} // 唯一签名

	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		NickName:    user.Nickname,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)

	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	
	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

}