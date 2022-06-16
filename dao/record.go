package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"restful/entity"
	"time"
)

var RecordDB *gorm.DB

func AddRecord(bookid, userid int, days int) error {
	t := time.Now()
	et := t.AddDate(0, 0, days)
	r := entity.Record{bookid, userid, t, et}
	fmt.Println(r)
	res := RecordDB.Create(&r)
	return res.Error
}

func Findlatestamonth() entity.RecordSlice {
	r := make(entity.RecordSlice, 0)
	t := time.Now().AddDate(0, -1, 0)
	RecordDB.Where("borrow>=?", t).Find(&r)
	return r
}

func Findbyuser(userid int) entity.RecordSlice {
	r := make(entity.RecordSlice, 0)
	RecordDB.Table("records").Select("*").Joins("left join books on records.bookid=books.id").Where("userid=? and books.status=?", userid, true).Find(&r)
	return r

}

func init() {
	RecordDB, _ = openDatabase("root", "123456", "restful", "mysql")
}
