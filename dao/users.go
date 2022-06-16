package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
import "restful/entity"
import "restful/utils"

var UserDb *gorm.DB
var err error

func AddUser(username, password, email string) error {
	salt := utils.RandString(5)
	np := password + salt
	p := utils.Cryptomd5(np)

	u := entity.User{Username: username, Password: p[0:15], Salt: salt, Email: email}
	fmt.Println("!!!!!!######", u)
	//UserDb, _ = openDatabase("root", "123456", "restful", "mysql")
	res := UserDb.Create(&u)
	return res.Error
}

func FindByname(name string) *entity.User {
	u := entity.User{}
	//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", UserDb)
	//fmt.Println("########################", name)
	res := UserDb.Where("username=?", name).Find(&u)
	if res.Error != nil {
		return nil
	}
	return &u
}

func init() {
	UserDb, err = openDatabase("root", "123456", "restful", "mysql")
}
