package model

import "time"

type BaseModel struct {
	Id       int64     `gorm:"primary_key;AUTO_INCREMENT;" json:"id" form:"id"`
	CreateBy int64     `gorm:"type:bigint(64)" json:"createBy" form:"createBy"`
	CreateAt time.Time `json:"createAt" form:"createAt"`
	UpdateBy int64     `gorm:"type:bigint(64)" json:"updateBy" form:"updateBy"`
	UpdateAt time.Time `json:"updateAt" form:"updateAt"`
	Version  int32     `gorm:"type:int(32);default:0;" json:"version" from:"version"`
	IsDelete int8      `gorm:"type:char(1);default:'0';" json:"isDelete" form:"isDelete"`
}
