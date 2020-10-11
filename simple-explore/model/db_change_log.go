package model

type ChangeLog struct {
	ChangeCode string `gorm:"type:varchar(64);unique"`
	BaseModel
}

func (cl *ChangeLog) TableName() string {
	return "change_log"
}
