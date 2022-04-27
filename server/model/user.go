package model

import "github.com/google/uuid"

type UserModel struct {
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"` 
	Username    string `json:"username" gorm:"comment:用户名"`             // 用户名
	Password 		string `json:"password" gorm:"comment:密码"`    // 密码
	Nickname    string `json:"nickname" gorm:"comment:昵称"`          // 昵称
	Mobile      string `json:"mobile" gorm:"default:POST;comment:手机号"` // 手机号
}

func (UserModel) TableName() string {
	return "user"
}
