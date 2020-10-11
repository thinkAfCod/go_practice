package model

type Tab struct {
	TabName     string `gorm:"type:varchar(32);not null;" json:"tabName" form:"tabName"`
	TabParentId int64  `gorm:"type:bigint(64);default:0;" json:"tabParentId" form:"tabParentId"`
	Icon        string `gorm:"type:varchar(512);default:'';" json:"icon" form:"icon"`
	Description string `gorm:"type:varchar(256);default:'';" json:"description" form:"description"`
	Remark      string `gorm:"type:varchar(128);default:'';" json:"remark" form:"remark"`
	BaseModel
}

func (t *Tab) TableName() string {
	return "tab"
}

func NewTab() {

}
