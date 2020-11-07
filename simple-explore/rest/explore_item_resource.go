package rest

import (
	"github.com/kataras/iris/v12"
	"simple-explore/common"
	"simple-explore/model"
	"simple-explore/repo"
	"strconv"
	"strings"
)

type ExploreItemResource struct {
}

var exploreRepository *repo.ExploreItemRepository

func init() {
	exploreRepository = repo.NewExploreItemRepository()
}

func (eir *ExploreItemResource) Get(ctx iris.Context) {
	page, parentItemId, name := eir.parseHttpParam(ctx)
	_, err := exploreRepository.FindAllByParentId(page, parentItemId, name)
	eir.handleResult(ctx, page, err, err)
}

func (eir *ExploreItemResource) GetParentId(ctx iris.Context) {
	id := ctx.URLParamDefault("id", "0")
	var selectParentId string
	var parentItem *model.ExploreFile
	var err error
	if !strings.EqualFold("0", id) {
		parentItem, err = exploreRepository.FindParentIdById(id)
		selectParentId = strconv.FormatInt(parentItem.ParentItemId, 10)
	} else {
		selectParentId = id
	}
	eir.handleResult(ctx, selectParentId, err, err)
}

func (eir *ExploreItemResource) handleResult(ctx iris.Context, data interface{}, errMsg interface{}, err error) {
	if err != nil {
		ctx.JSON(common.Error(errMsg))
		return
	}
	ctx.JSON(common.Success(data))
}

func (eir *ExploreItemResource) parseHttpParam(ctx iris.Context) (*model.PageModel, string, string) {
	page := &model.PageModel{}
	page.PageSize, _ = strconv.ParseInt(ctx.URLParamDefault("pageSize", "10"), 10, 64)
	page.PageNo, _ = strconv.ParseInt(ctx.URLParamDefault("pageNo", "1"), 10, 64)
	parentItemId := ctx.URLParamDefault("parentId", "0")
	name := ctx.URLParam("name")
	return page, parentItemId, name
}

func (eir *ExploreItemResource) Post(ctx iris.Context) {
	item := &model.ExploreItem{}
	err := ctx.ReadJSON(item)
	if err != nil {
		ctx.JSON(common.Error(err))
		return
	}
	err = exploreRepository.Insert(item)
	eir.handleResult(ctx, nil, "id is invalid", err)
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
	eir.handleResult(ctx, nil, "id is invalid", err)
}

func (eir *ExploreItemResource) Delete(ctx iris.Context) {
	//item model.ExploreItem, physics bool
	id, _ := strconv.ParseInt(ctx.URLParam("id"), 10, 64)
	if id == 0 {
		ctx.JSON(common.Error("id is invalid"))
		return
	}
	err := exploreRepository.DeleteById(&model.ExploreItem{}, id)
	eir.handleResult(ctx, nil, "id is invalid", err)
}
