package controller

import (
	"github.com/gin-gonic/gin"
	"restful/dao"
	"restful/service"
	"strconv"
)

//还书：/book/return
func ReturnBook(ctx *gin.Context) {
	bookid := ctx.PostForm("bookid")
	bid, _ := strconv.Atoi(bookid)
	dao.UpdateBookStatus(bid)
	ctx.String(200, "还书成功")
	//TODO:还书逻辑
}

//往图书馆添加一本书：路径/book/add
func AddBook(ctx *gin.Context) {
	//bookname := ctx.PostForm("bookname")
	//num := ctx.PostForm("num")
	//TODO:加书逻辑
}

//展示图书馆的图书状况：/book/show
func ShowBook(ctx *gin.Context) {
	page := ctx.Param("page")
	//TODO:展示逻辑
	p, _ := strconv.Atoi(page)
	bs := service.GetBooks(p)
	ctx.HTML(200, "books.html", gin.H{
		"books": bs,
	})

}

//借书：/book/borrow
func BorrowBook(ctx *gin.Context) {
	bookid := ctx.Param("bookid")
	userid := ctx.Param("userid")
	days := ctx.Param("days")
	bid, _ := strconv.Atoi(bookid)
	uid, _ := strconv.Atoi(userid)
	d, _ := strconv.Atoi(days)
	service.BorrowBook(bid, uid, d)
	ctx.String(200, "借书成功")
	//TODO:借书逻辑
}
