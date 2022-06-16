package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restful/service"
	"strconv"
)

//查询一个用户的借书记录：/record/search/:userid
func SearchByUser(ctx *gin.Context) {
	uid := ctx.Param("userid")
	fmt.Println("$$$$$$$$$$$$$$$4", uid)
	id, _ := strconv.Atoi(uid)
	r := service.Searchbyusertoreturn(id)
	ctx.JSON(200, gin.H{
		"records": r,
	})
}

//url:/record/all
func SearchAlltoReturn(ctx *gin.Context) {
	r := service.SearchtoReturn()
	fmt.Println("!!!!!!!!!!!!!!!!!!!!:", len(r))
	ctx.JSON(200, gin.H{
		"records": r,
	})
}

//url:/record/time
func Searchbydaysandmonth(ctx *gin.Context) {
	d := ctx.Query("day")
	m := ctx.Query("month")
	day, _ := strconv.Atoi(d)
	month, _ := strconv.Atoi(m)
	r := service.Searchbymonthandday(day, month)
	ctx.JSON(200, gin.H{
		"records": r,
	})
}

//初始化记录路由
func Initrecordrouter(engine *gin.Engine) {
	g := engine.Group("/record")
	//根据时间查询记录
	g.GET("/time", Searchbydaysandmonth)
	//查询所有为归还记录
	g.GET("/all", SearchAlltoReturn)
	//查询某人未归还的记录
	g.GET("/search/:userid", SearchByUser)
}
