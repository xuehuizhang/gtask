package main

import (
	"codego/gtask/initialize"
	"codego/gtask/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//项目初始化
	initialize.Init()

	//初始化web
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":9091")
}
