package controller

import (
	"github.com/gin-gonic/gin"
)

//还书：/book/return
func ReturnBook(ctx *gin.Context) {
	bookid := ctx.PostForm("id")
	//TODO:还书逻辑
}

//往图书馆添加一本书：路径/book/add
func AddBook(ctx *gin.Context) {
	bookname := ctx.PostForm("bookname")
	num := ctx.PostForm("num")
	//TODO:加书逻辑
}

//展示图书馆的图书状况：/book/all
func ShowBook(ctx *gin.Context) {
	page := ctx.Param("page")
	//TODO:展示逻辑
}

//借书：/book/borrow
func BorrowBook(ctx *gin.Context) {
	bookid := ctx.PostForm("bookid")
	userid := ctx.PostForm("userid")
	//TODO:借书逻辑
}
