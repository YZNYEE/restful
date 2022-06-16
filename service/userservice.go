package service

import (
	"errors"
	"fmt"
	"restful/dao"
	"restful/utils"
)

func Check(username, password string) error {
	u := dao.FindByname(username)
	fmt.Println("!!!!!!!!!!")
	fmt.Println(username, u)
	if u == nil {
		return errors.New("用户名或者密码错误")
	}
	s := u.Salt
	np := password + s
	cnp := utils.Cryptomd5(np)
	if cnp[0:15] == u.Password {
		return nil
	}
	return errors.New("用户名或者密码错误")
}

func Register(username, password, email string) error {
	return dao.AddUser(username, password, email)
}
