package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/name", func(ctx *gin.Context) {
		time.Sleep(time.Second * 35)
		ctx.JSON(http.StatusOK, gin.H{"name": "张三"})
		return
	})
	r.Run(":9098")
}
