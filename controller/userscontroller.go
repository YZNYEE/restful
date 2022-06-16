package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restful/service"
	"strconv"
)

//登入：路径/users/login
func Login(ctx *gin.Context) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!")
	name := ctx.Query("username")
	password := ctx.Query("password")
	//TODO:认证
	fmt.Println("name", name, "password", password)
	err := service.Check(name, password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "用户或密码错误",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg": "欢迎" + name,
		})
	}
}

func Return(ctx *gin.Context) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!")
	bookid := ctx.Query("bookid")
	bid, _ := strconv.Atoi(bookid)
	err := service.ReturnBook(bid)
	if err == nil {
		ctx.JSON(200, gin.H{
			"msg": "还书成功!!!",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg": "还书异常!!!",
		})
	}
}

func UserAction(ctx *gin.Context) {
	action := ctx.Param("action")
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!:", action)
	switch action {
	case "/login":
		Login(ctx)
	case "/register":
		Register(ctx)
	case "/borrow":
		Borrow(ctx)
	case "/return":
		Return(ctx)
	}
}
func Borrow(ctx *gin.Context) {
	bookid := ctx.Query("bookid")
	userid := ctx.Query("userid")
	days := ctx.Query("days")
	bid, _ := strconv.Atoi(bookid)
	uid, _ := strconv.Atoi(userid)
	d, _ := strconv.Atoi(days)
	err := service.BorrowBook(bid, uid, d)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "借书失败！！该书已被借走",
		})
	} else {
		ctx.JSON(400, gin.H{
			"msg": "借书成功",
		})
	}
}

//注册：/users/register
func Register(ctx *gin.Context) {
	name := ctx.Query("username")
	password := ctx.Query("password")
	email := ctx.Query("email")
	//TODO:提交信息
	err := service.Register(name, password, email)
	if err == nil {
		ctx.JSON(400, gin.H{"msg": "注册成功" + "欢迎:" + name})
	} else {
		//这里应该是ajax异步请求，没时间写
		ctx.String(400, "用户名已被注册!!!!!")
	}
}
