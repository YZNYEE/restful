package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restful/dao"
	"restful/service"
	"strconv"
)

//往图书馆添加一本书：路径/book/add

//展示图书馆的图书状况：/book/show
func ShowBook(ctx *gin.Context) {
	page := ctx.Query("page")
	//TODO:展示逻辑
	p, _ := strconv.Atoi(page)
	bs := service.GetBooks(p)
	ctx.JSON(200, gin.H{
		"books": bs,
	})

}

func FindBook(ctx *gin.Context) {
	test := ctx.Query("bookname")
	bookname := test
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

func Add(ctx *gin.Context) {
	bookname := ctx.PostForm("bookname")
	dao.AddBook(bookname)
	ctx.JSON(200, gin.H{
		"msg": "添加成功",
	})
}

//书籍路由
func Initbookrouter(engine *gin.Engine) {
	g := engine.Group("/book")
	//查询书籍
	g.GET("/search", FindBook)
	//查询可借书籍
	g.GET("/borrow", Findcanborrowedbook)
	//添加书籍
	g.POST("/add", Add)
	//展示书籍
	g.GET("/show", ShowBook)
}
