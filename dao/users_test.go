package dao

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	username, password := "yangjhhy", "123456"
	AddUser(username, password, "test@test.com")
}
func TestFindByname(t *testing.T) {
	n := "test1"
	u := FindByname(n)
	fmt.Println(u)
}
