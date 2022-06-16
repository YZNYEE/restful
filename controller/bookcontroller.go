package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restful/dao"
	"restful/service"
	"strconv"
)

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

func FindBook(ctx *gin.Context) {
	bookname := ctx.Query("name")
	test := ctx.Query("test")
	fmt.Println("test:", test)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!", bookname)
	bs := dao.Findbyname(bookname)
	if len(bs) == 0 {
		j := gin.H{
			"msg": "未找到书籍:" + bookname,
		}
		ctx.JSON(400, j)
	} else {
		s := fmt.Sprintf("找到书籍%s:%d本", bookname, len(bs))
		j := gin.H{
			"msg": s,
		}
		ctx.JSON(200, j)
	}
}
func Findcanborrowedbook(ctx *gin.Context) {
	bookname := ctx.Query("bookname")
	bs := dao.Findbynameandstatus(bookname, false)
	if len(bs) == 0 {
		j := gin.H{
			"msg":   "无书籍" + bookname + "可借",
			"books": bs,
		}
		ctx.JSON(400, j)
	} else {
		s := fmt.Sprintf("找到书籍%s:%d本", bookname, len(bs))
		j := gin.H{
			"msg":   s,
			"books": bs,
		}
		ctx.JSON(200, j)
	}
}

func Initbookrouter(engine *gin.Engine) {
	g := engine.Group("/book")
	g.GET("/search", FindBook)
	g.GET("/borrow", Findcanborrowedbook)
}
