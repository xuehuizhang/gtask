package main

import (
	"codego/gtask/router"
	"codego/gtask/setting"
	"codego/gtask/task"
	"github.com/gin-gonic/gin"
)

func main() {
	setting.SetUp()
	//初始化定时任务
	go task.InitTask()

	//初始化web
	r := gin.Default()

	router.InitRouter(r)

	r.Run(":9091")
}
