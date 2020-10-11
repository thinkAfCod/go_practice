package repo

//import (
//	"simple-explore/db"
//	"simple-explore/model"
//)
//
//type CatalogueRepo struct {
//	BaseRepo
//}
//
//func (ddr *CatalogueRepo) QueryById(id string) (*model.Catalogue, error) {
//	cata := &model.Catalogue{}
//	err := ddr.BaseRepo.QueryById(cata, id)
//	return cata, err
//}
//
//func (ddr *CatalogueRepo) FindCateByParentId(parentId string) (*[]model.Catalogue, error) {
//	catas := new([]model.Catalogue)
//	tx := db.GetWithTx()
//	tx.Find(catas, "parent_id = ?", parentId)
//	return catas, tx.Commit().Error
//}
//
//func (ddr *CatalogueRepo) FindDetailByCataId(cataId string) (*[]model.Catalogue, error) {
//	details := new([]model.Catalogue)
//	tx := db.GetWithTx()
//	tx.Find(details, "parent_id = ?", cataId)
//	return details, tx.Commit().Error
//}
