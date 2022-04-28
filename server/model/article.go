package model

import "github.com/google/uuid"

type ArticleModel struct {
	UUID        uuid.UUID      `json:"uuid" gorm:"comment: 文章UUID"` 
	Title string `json:"title" gorm:"comment: 标题"`
	Author string `json:"author" gorm:"comment: 作者"`
	Pageview string `json:"pageview" gorm:"default: 1;comment: 阅读数"`
	CreateTime string `json:"create_time" gorm:"comment: 发布时间"`
	Content string `json:"content" gorm:"comment: 文章内容"`
}

func (ArticleModel) TableName() string {
	return "article"
}