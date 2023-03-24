package router

import (
	"codego/gtask/api/v1/service"
	"github.com/gin-gonic/gin"
)

func initServiceRoute(g *gin.RouterGroup) {
	t := g.Group("/service")
	{
		t.POST("/save", service.Save)
		t.POST("/del", service.Del)
		t.GET("/info", service.Info)
		t.GET("/list", service.List)
	}
}
