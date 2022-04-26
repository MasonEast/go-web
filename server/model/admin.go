package model

type AdminModel struct {
	Username        string `json:"username" gorm:"comment:api路径"`             // api路径
	Password string `json:"password" gorm:"comment:api中文描述"`    // api中文描述
	Nickname    string `json:"nickname" gorm:"comment:api组"`          // api组
	Mobile      string `json:"mobile" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (AdminModel) TableName() string {
	return "admin"
}
