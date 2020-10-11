package rest

import (
	"github.com/kataras/iris/v12"
	"simple-explore/common"
	"simple-explore/model"
	"simple-explore/repo"
	"strconv"
)

type ExploreItemResource struct {
}

var exploreRepository *repo.ExploreItemRepository

func init() {
	exploreRepository = repo.NewExploreItemRepository()
}

func (eir *ExploreItemResource) Get(ctx iris.Context) {
	page := &model.PageModel{}
	page.PageSize, _ = strconv.ParseInt(ctx.URLParamDefault("pageSize", "10"), 10, 64)
	page.PageNo, _ = strconv.ParseInt(ctx.URLParamDefault("pageNo", "1"), 10, 64)
	parentItemId := ctx.URLParamDefault("parentId", "0")
	name := ctx.URLParam("name")
	_, err := exploreRepository.FindAllByParentId(page, parentItemId, name)
	if err != nil {
		ctx.JSON(common.Error(err))
		return
	}
	ctx.JSON(common.Success(page))
}

func (eir *ExploreItemResource) Post(ctx iris.Context) {
	item := &model.ExploreItem{}
	err := ctx.ReadJSON(item)
	if err != nil {
		ctx.JSON(common.Error(err))
		return
	}
	err = exploreRepository.Insert(item)
	if err != nil {
		ctx.JSON(common.Error(err))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (eir *ExploreItemResource) Patch(ctx iris.Context) {
	//item model.ExploreItem
	item := &model.ExploreItem{}
	err := ctx.ReadJSON(item)
	if err != nil {
		ctx.JSON(common.Error(err))
		return
	}
	if item.Id == 0 {
		ctx.JSON(common.Error("id is invalid"))
		return
	}
	err = exploreRepository.Update(&item)
	if err != nil {
		ctx.JSON(common.Error("id is invalid"))
		return
	}
	ctx.JSON(common.Success(nil))
}

func (eir *ExploreItemResource) Delete(ctx iris.Context) {
	//item model.ExploreItem, physics bool
	id, _ := strconv.ParseInt(ctx.URLParam("id"), 10, 64)
	if id == 0 {
		ctx.JSON(common.Error("id is invalid"))
		return
	}
	err := exploreRepository.DeleteById(&model.ExploreItem{}, id)
	if err != nil {
		ctx.JSON(common.Error("id is invalid"))
		return
	}
	ctx.JSON(common.Success(nil))
}
