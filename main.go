package main

import (
	"github.com/gin-gonic/gin"
	"restful/controller"
	_ "restful/dao"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	g := r.Group("/")
	g.GET("/login", controller.Login)
	g.GET("/register", controller.Register)
	r.GET("/users/*action", controller.UserAction)
	//g = r.Group("/users")
	//g.GET("/login", controller.Loginusers)
	//g.POST("/register", controller.Registerusers)
	g = r.Group("/book")
	g.GET("/show/:page", controller.ShowBook)
	g.GET("/borrow/:userid/:bookid/:days", controller.BorrowBook)
	g.POST("/return", controller.ReturnBook)
	r.Run(":80")

}
