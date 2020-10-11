package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"simple-explore/common"
	"simple-explore/conf"
)

var Db *gorm.DB

func InitByConfig(config conf.DataSource) {
	url := fmt.Sprintf("%s:%s@(%s)%s", config.Username, config.Password, config.Domain, config.Path)
	db, err := gorm.Open(config.DriverName, url)
	common.PanicErr(err)
	db.LogMode(true)
	Db = db
	InitTable()
}

func Get() *gorm.DB {
	if Db == nil {
		panic(errors.New("instance of Db is init"))
	}
	return Db
}

func GetWithTx() *gorm.DB {
	if Db == nil {
		panic(errors.New("instance of Db is init"))
	}
	return Db.Begin()
}

//func Create(arg interface{}) error {
//	//allowAction := []string{"INSERT", "UPDATE", "DELETE"}
//	//err := check(statement, allowAction)
//	//common.PanicErr(err)
//	defer Db.Commit()
//	tx := Db.Begin()
//	tx.Create(arg)
//	return nil
//}
//
//func Update(arg interface{}) error {
//	//allowAction := []string{"INSERT", "UPDATE", "DELETE"}
//	//err := check(statement, allowAction)
//	//common.PanicErr(err)
//	defer Db.Commit()
//	tx := Db.Begin()
//	tx.Update(arg)
//	return nil
//}
//
//func Delete(arg interface{}) error {
//	//allowAction := []string{"INSERT", "UPDATE", "DELETE"}
//	//err := check(statement, allowAction)
//	//common.PanicErr(err)
//	defer Db.Commit()
//	tx := Db.Begin()
//	tx.Delete(arg)
//	return nil
//}
//
//func Query(statement string, ,args ...interface{}) {
//	allowAction := []string{"SELECT"}
//	err := check(statement, allowAction)
//	common.PanicErr(err)
//	defer Db.Close()
//	err = Db.Ping()
//	common.PanicErr(err)
//	tx, err := Db.Begin()
//	common.PanicErr(err)
//	row, err := tx.Query(statement, args)
//	common.PanicErr(err)
//	row.
//	return id, rows, nil
//}
//
//func check(statement string, allowAction []string) error {
//	if len(statement) == 0 {
//		return errors.New("statement is an empty string")
//	}
//	action := strings.ToUpper(strings.Fields(statement)[0])
//	for _, aclAction := range allowAction {
//		if strings.EqualFold(action, aclAction) {
//			return nil
//		}
//	}
//	return errors.New("action is not allowed here:" + action)
//}
