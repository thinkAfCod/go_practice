package repo

import (
	"fmt"
	"simple-explore/db"
	"simple-explore/model"
)

type ExploreItemRepository struct {
	BaseRepo
}

var eit *ExploreItemRepository

func init() {
	eit = new(ExploreItemRepository)
}

func NewExploreItemRepository() *ExploreItemRepository {
	return eit
}

func (eir *ExploreItemRepository) FindAllByTabIdAndName(
	page *model.PageModel, tabId int64, name string) (*[]model.ExploreItem, error) {
	groups := new([]model.ExploreItem)
	err := eir.QueryAll(groups, &model.ExploreItem{}, page, "tab_id = ? and parent_item_id = 0 and (name like '%?%')", tabId,
		name)
	return groups, err
}

func (eir *ExploreItemRepository) FindAllByParentId(page *model.PageModel, parentItemId string,
	name string) ([]*model.ExploreFile, error) {
	groups := new([]*model.ExploreFile)
	sql := `
         SELECT ei.id,ei.tab_id,ei.file_id,f.media_type,ei.parent_item_id,ei.name,ei.cover,ei.description,ei.keyword
         FROM explore_item ei,file f
         WHERE ei.file_id = f.id and ei.parent_item_id = ?  LIMIT ? OFFSET ?
    `
	countsql := `
         SELECT COUNT(*) as count
         FROM explore_item 
         WHERE parent_item_id = ?
    `
	tx := db.GetWithTx()
	tx.Raw(sql, parentItemId, page.PageSize, (page.PageNo-1)*page.PageSize).Scan(groups)
	tx.Raw(countsql, parentItemId).Scan(page)
	fmt.Println(page)
	page.PageTotal = page.Count / page.PageSize
	if page.Count%page.PageSize > 0 {
		page.PageTotal += 1
	}
	page.PageData = groups
	return *groups, tx.Commit().Error
}

func (eir *ExploreItemRepository) deleteByParentId(parentItemId string) error {
	tx := db.GetWithTx()
	tx.Delete(&model.ExploreItem{}, "parent_item_id = ?", parentItemId)
	return tx.Commit().Error
}
