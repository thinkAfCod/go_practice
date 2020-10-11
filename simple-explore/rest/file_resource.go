package rest

import (
	"github.com/kataras/iris/v12"
	"simple-explore/common"
	"simple-explore/model"
	"simple-explore/repo"
)

var fr *repo.FileRepository

func init() {
	fr = repo.NewFileRepository()
}

type FileResource struct {
}

func (fres *FileResource) Get(ctx iris.Context) {
	id := ctx.URLParam("id")
	fileModel := &model.FileModel{}
	err := fr.QueryById(fileModel, id)
	common.PanicErr(err)
	err = ctx.SendFile(fileModel.AbsoluteUri, fileModel.Name)
	common.PanicErr(err)
}
