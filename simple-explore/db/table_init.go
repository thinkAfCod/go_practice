package db

import (
	"mime"
	"os"
	"path"
	"simple-explore/common"
	"simple-explore/model"
	"strings"
	"time"
)

func InitTable() {
	Db.AutoMigrate(&model.ChangeLog{}, &model.Tab{}, &model.ExploreItem{}, &model.FileModel{})
	PrepareData()
	//PrepareData2()
}

func PrepareData() {
	now := time.Now()
	user := int64(0)
	changeLog := &model.ChangeLog{
		ChangeCode: "1",
		BaseModel: model.BaseModel{
			CreateBy: user,
			CreateAt: now,
			UpdateBy: user,
			UpdateAt: now,
		},
	}
	existChangeLog := &model.ChangeLog{}
	Db.Find(existChangeLog, "change_code = ?", "1")

	if strings.EqualFold(existChangeLog.ChangeCode, changeLog.ChangeCode) {
		return
	}

	movie := &model.Tab{
		BaseModel: model.BaseModel{
			Id:       1,
			CreateBy: user,
			CreateAt: now,
			UpdateBy: user,
			UpdateAt: now,
		},
		TabName: "电影",
	}
	photoAlbum := &model.Tab{
		BaseModel: model.BaseModel{
			Id:       2,
			CreateBy: user,
			CreateAt: now,
			UpdateBy: user,
			UpdateAt: now,
		},
		TabName: "写真集",
	}
	comic := &model.Tab{
		BaseModel: model.BaseModel{
			Id:       3,
			CreateBy: user,
			CreateAt: now,
			UpdateBy: user,
			UpdateAt: now,
		},
		TabName: "漫画",
	}
	asmr := &model.Tab{
		BaseModel: model.BaseModel{
			Id:       4,
			CreateBy: user,
			CreateAt: now,
			UpdateBy: user,
			UpdateAt: now,
		},
		TabName: "音频",
	}
	novel := &model.Tab{
		BaseModel: model.BaseModel{
			Id:       5,
			CreateBy: user,
			CreateAt: now,
			UpdateBy: user,
			UpdateAt: now,
		},
		TabName: "小说",
	}
	orginCatas := []*model.Tab{movie, photoAlbum, comic, asmr, novel}
	for _, cata := range orginCatas {
		Db.Save(cata)
	}
	Db.Save(changeLog)
}

func PrepareData2() {
	filePathList := []string{
		"F:\\databackup\\小电影",
		"F:\\databackup\\图集",
		"F:\\databackup\\漫画",
		"F:\\databackup\\AMSR",
		"F:\\databackup\\书籍",
	}
	for index, filePath := range filePathList {
		go loop(int64(index+1), 0, filePath)
	}
}

func loop(tabId int64, parentId int64, filePath string) {
	subId := InsertRec(tabId, parentId, filePath)
	lstat, err := os.Lstat(filePath)
	common.PanicErr(err)
	if lstat.IsDir() {
		f, err := os.Open(filePath)
		common.PanicErr(err)
		names, err := f.Readdirnames(-1)
		f.Close()
		common.PanicErr(err)
		for _, name := range names {
			subFilePath := path.Join(filePath, name)
			loop(tabId, subId, subFilePath)
		}
	}
}

func InsertRec(tabId int64, parentId int64, filePath string) int64 {
	mimeType := ""
	lstat, err := os.Lstat(filePath)
	common.PanicErr(err)
	now := time.Now()
	base := &model.BaseModel{
		CreateBy: 0,
		CreateAt: now,
		UpdateBy: 0,
		UpdateAt: now,
	}
	lastIndex := strings.LastIndex(filePath, ".")
	if !lstat.IsDir() && lastIndex != -1 {
		mimeType = mime.TypeByExtension(filePath[lastIndex:])
	}
	fileModel := &model.FileModel{
		Name:        lstat.Name(),
		AbsoluteUri: filePath,
		MediaType:   mimeType,
		BaseModel:   *base,
	}
	Db.Save(fileModel)
	item := &model.ExploreItem{
		TabId:        tabId,
		FileId:       fileModel.Id,
		ParentItemId: parentId,
		Name:         lstat.Name(),
		Description:  "",
		Keyword:      "",
		BaseModel:    *base,
	}
	Db.Save(item)
	return item.Id
}
