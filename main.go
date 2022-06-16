package main

import (
	"github.com/gin-gonic/gin"
	"restful/controller"
	_ "restful/dao"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	controller.Inituserroute(r)
	controller.Initbookrouter(r)
	controller.Initrecordrouter(r)
	r.Run(":80")

}
