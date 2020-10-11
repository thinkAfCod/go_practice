package model

type FileModel struct {
	Name        string `gorm:"type:varchar(256);not null;" json:"fileName" form:"fileName"`
	AbsoluteUri string `gorm:"type:varchar(1024);not null;" json:"-" form:"absoluteUri"`
	MediaType   string `gorm:"type:varchar(64);" json:"mediaType" form:"mediaType"`
	Md5         string `gorm:"type:varchar(64);" json:"md5" form:"md5"`
	BaseModel
}

func (fm *FileModel) TableName() string {
	return "file"
}
