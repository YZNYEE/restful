package dao

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestOpendatabase(t *testing.T) {
	db, err := openDatabase("root", "123456", "gomysql", "mysql")
	fmt.Println(db, err)
}
