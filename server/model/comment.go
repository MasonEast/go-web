package model

type CommentModel struct {
	User UserModel
	CreateTime string `json:"create_time" gorm:"comment: 发布时间"`
	Content string `json:"content" gorm:"comment: 文章内容"`
	Star int 
}