package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"restful/entity"
)

var BookDB *gorm.DB

func AddBook(bookname string) {
	b := entity.Book{Bookname: bookname, Status: false}
	BookDB.Create(&b)
	//fmt.Println(res)
}

func SelectbyOffetandlimit(offset, limit int) entity.BookSlice {
	b := make(entity.BookSlice, 10)
	BookDB.Offset(offset).Limit(limit).Find(&b)
	return b
}

func UpdateBookStatus(bookid int) {
	b := FindbyId(bookid)
	b.Status = !b.Status
	BookDB.Save(b)
	//BookDB.Model(&entity.Book{}).Where("id", bookid).Update("status", true)
}

func FindbyId(bookid int) *entity.Book {
	b := entity.Book{}
	BookDB.Where("id=?", bookid).Find(&b)
	fmt.Println(bookid, b)
	return &b
}

func Findbyname(bookname string) entity.BookSlice {
	bs := make(entity.BookSlice, 0)
	BookDB.Where("bookname=?", bookname).Find(&bs)
	return bs
}

func init() {
	BookDB, _ = openDatabase("root", "123456", "restful", "mysql")

}
