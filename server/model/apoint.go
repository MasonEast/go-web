package model

import "github.com/google/uuid"

type ApointModel struct {
	UUID        uuid.UUID      `json:"uuid" gorm:"comment: 文章UUID"` 
	Data string `json:"data" gorm:"comment: 数据"`
	// Data ApointData `json:"data" gorm:"comment: 数据"`
	CreateTime string `json:"create_time" gorm:"comment: 创建时间"`
}

type ApointData struct {
	Browser Browser `json:"browser"`
}

type Browser struct {
	Size string `json:"size"`
	Network string `json:"network"`
	Language string `json:"language"`
	Timezone string `json:"timezone"`
	Ua string `json:"ua"`
	Os string `json:"os"`
	Engine string `json:"engine"`
}

func (ApointModel) TableName() string {
	return "apoint" 
}
