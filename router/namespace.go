package router

import (
	"codego/gtask/api/v1/namespace"
	"github.com/gin-gonic/gin"
)

func initNamespaceRoute(g *gin.RouterGroup) {
	t := g.Group("/namespace")
	{
		t.POST("/save", namespace.Save)
		t.POST("/del", namespace.Del)
		t.GET("/info", namespace.Info)
		t.GET("/list", namespace.List)
	}
}
