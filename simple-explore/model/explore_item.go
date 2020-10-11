package model

type ExploreItem struct {
	TabId        int64     `gorm:"type:bigint(64);not null;" json:"tabId" form:"tabId"`
	FileId       int64     `gorm:"type:bigint(64);not null;" json:"fileId" form:"fileId"`
	File         FileModel `gorm:"references:file_id" json:"file" form:"file"`
	ParentItemId int64     `gorm:"type:bigint(64);default:0;" json:"parentItemId" form:"parentItemId"`
	Name         string    `gorm:"type:varchar(256);not null;" json:"name" form:"name"`
	Cover        string    `gorm:"type:varchar(1024);" json:"cover" form:"cover"`
	Description  string    `gorm:"type:varchar(512);" json:"description" form:"description"`
	Keyword      string    `gorm:"type:varchar(256);" json:"keyword" form:"keyword"`
	BaseModel
}

func (ei *ExploreItem) TableName() string {
	return "explore_item"
}
