package model

import "github.com/google/uuid"

type ApointModel struct {
	UUID        uuid.UUID      `json:"uuid" gorm:"comment: 文章UUID"` 
	Data string `json:"data" gorm:"comment: 数据"`
	CreateTime string `json:"create_time" gorm:"comment: 创建时间"`
}

func (ApointModel) TableName() string {
	return "apoint"
}