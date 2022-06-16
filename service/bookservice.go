package service

import (
	"errors"
	"restful/dao"
	"restful/entity"
)

func BorrowBook(bookid, userid, days int) error {
	b := dao.FindbyId(bookid)
	if b == nil || b.Status {
		return errors.New("借书失败")
	}
	//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!", bookid, userid, days)
	res := dao.AddRecord(bookid, userid, days)
	dao.UpdateBookStatus(bookid)
	return res
}
func ReturnBook(bookid int) error {
	b := dao.FindbyId(bookid)
	if b == nil || !b.Status {
		return errors.New("还书异常！请找管理员解决")
	} else {
		dao.UpdateBookStatus(bookid)
		return nil
	}
}

func GetBooks(page int) entity.BookSlice {
	bs := dao.SelectbyOffetandlimit(page*10, 10)
	return bs
}
