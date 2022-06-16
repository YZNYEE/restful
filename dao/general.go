package dao

import "fmt"
import "github.com/jinzhu/gorm"

func openDatabase(user, password, dbname, dialect string) (*gorm.DB, error) {
	str := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)
	db, err := gorm.Open(dialect, str)
	return db, err
}
