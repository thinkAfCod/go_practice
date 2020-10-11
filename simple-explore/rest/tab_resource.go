package rest

import (
	"github.com/kataras/iris/v12"
	"simple-explore/common"
	"simple-explore/model"
	"simple-explore/repo"
	"strconv"
	"strings"
)

type TabResource struct {
}

var tabRepository repo.TabRepository

func init() {
	tabRepository = repo.TabRepository{}
}

func (cr *TabResource) Get(ctx *iris.Context) {
	page := &model.PageModel{}
	page.PageSize, _ = strconv.ParseInt((*ctx).URLParamDefault("pageSize", "10"), 10, 64)
	page.PageNo, _ = strconv.ParseInt((*ctx).URLParamDefault("pageNo", "1"), 10, 64)
	name := (*ctx).URLParam("name")
	out := new([]*model.Tab)
	var err error
	if strings.EqualFold(name, "") {
		err = tabRepository.QueryAll(out, &model.Tab{}, page)
	} else {
		err = tabRepository.QueryAll(out, &model.Tab{}, page, "tab_name like '%?%'", name)
	}
	if err != nil {
		(*ctx).JSON(common.Error(err))
	}
	(*ctx).JSON(common.Success(page))
}

func (cr *TabResource) Post(ctx *iris.Context) {
	tab := model.Tab{}
	(*ctx).ReadJSON(&tab)
	err := tabRepository.Insert(&tab)
	if err != nil {
		(*ctx).JSON(common.Error(err))
	}
	(*ctx).JSON(common.Success(nil))
}

func (cr *TabResource) Patch(ctx *iris.Context) {
	tab := model.Tab{}
	(*ctx).ReadJSON(&tab)
	if tab.Id == 0 {
		(*ctx).JSON(common.Error("id is invalid"))
		return
	}
	err := tabRepository.Update(&tab)
	if err != nil {
		(*ctx).JSON(common.Error(err))
	}
	(*ctx).JSON(common.Success(nil))
}

func (cr *TabResource) DeleteTab(ctx *iris.Context) {
	//tab model.Tab, physics bool
	id, _ := strconv.ParseInt((*ctx).URLParam("id"), 10, 64)
	if id == 0 {
		(*ctx).ReadJSON(common.Success("id is invalid"))
	}
	err := tabRepository.DeleteById(&model.Tab{}, id)
	if err != nil {
		(*ctx).JSON(common.Error(err))
	}
	(*ctx).JSON(common.Success(nil))
}
