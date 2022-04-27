package model

type AdminModel struct {
	Username        string `json:"username" gorm:"comment:用户名"`             // 用户名
	Password string `json:"password" gorm:"comment:密码"`    // 密码
	Nickname    string `json:"nickname" gorm:"comment:昵称"`          // 昵称
	Mobile      string `json:"mobile" gorm:"default:POST;comment:手机号"` // 手机号
}

func (AdminModel) TableName() string {
	return "admin"
}
