package repo

import (
	"simple-explore/db"
	"simple-explore/model"
)

type BaseRepo struct {
	//PrimaryIdName  string `cond:"id,"`
	//CreateTimeName string `cond:"createtime,createdate"`
	//UpdateTimeName string `cond:"createtime,createdate"`
	//ParamType reflect.Type
}

func (br *BaseRepo) QueryAll(out interface{}, argType interface{}, page *model.PageModel, where ...interface{}) error {
	tx := db.GetWithTx()
	whereDb := tx.Model(argType).Where(where[0], where[1:]...)
	whereDb.Count(&page.Count)
	whereDb.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(out)
	page.PageTotal = page.Count / page.PageSize
	if page.Count%page.PageSize > 0 {
		page.PageTotal += 1
	}
	page.PageData = out
	return tx.Commit().Error
}

func (br *BaseRepo) QueryById(out interface{}, id string) error {
	tx := db.GetWithTx()
	tx.Find(out, " id = ? ", id)
	return tx.Commit().Error
}

func (br *BaseRepo) Insert(arg interface{}) error {
	tx := db.GetWithTx()
	tx.Create(arg)
	return tx.Commit().Error
}

func (br *BaseRepo) Save(arg interface{}) error {
	tx := db.GetWithTx()
	tx.Save(arg)
	return tx.Commit().Error
}

func (br *BaseRepo) Update(arg interface{}) error {
	tx := db.GetWithTx()
	tx.Update(arg)
	return tx.Commit().Error
}

func (br *BaseRepo) Delete(arg interface{}) error {
	tx := db.GetWithTx()
	tx.Delete(arg)
	return tx.Commit().Error
}

func (br *BaseRepo) DeleteById(arg interface{}, id int64) error {
	tx := db.GetWithTx()
	tx.Delete(arg, "id = ?", id)
	return tx.Commit().Error
}

func (br *BaseRepo) DeleteBatchInIds(arg interface{}, ids []int64) error {
	tx := db.GetWithTx()
	tx.Delete(arg, "id in(?)", ids)
	return tx.Commit().Error
}

//func (br *BaseRepo) findPrimary(arg interface{}) (error) {
//	reflectType := reflect.ValueOf(arg).Type()
//	for reflectType.Kind() == reflect.Slice || reflectType.Kind() == reflect.Ptr {
//		reflectType = reflectType.Elem()
//	}
//	if reflectType.Kind() != reflect.Struct {
//		return errors.New("can not find struct type")
//	}
//	field, _ := reflect.ValueOf(br).Elem().Type().FieldByName("PrimaryIdName")
//	tags := field.Tag.Get("cond")
//	for i := 0; i < reflectType.NumField(); i++ {
//		structField := reflectType.Field(i)
//		fieldTag := structField.Tag.Get("gorm")
//		if strings.Contains(strings.ToLower(fieldTag), "primary") || strings.Contains(strings.ToLower(tags), strings.ToLower(structField.Name)) {
//			br.PrimaryIdName = structField.Name
//			break
//		}
//	}
//	if strings.EqualFold(br.PrimaryIdName, "") {
//		return errors.New("Primary id is not been set")
//	} else {
//		return nil
//	}
//}
