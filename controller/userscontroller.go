package controller

import (
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	name := ctx.Param("username")
	password := ctx.Param("password")
	//TODO:认证

}

func Register(ctx *gin.Context) {
	name := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//TODO:提交信息
}
