package router

import (
	"codego/gtask/api/v1/tenant"
	"github.com/gin-gonic/gin"
)

func initTenantRouter(r *gin.RouterGroup) {
	t := r.Group("/tenant")
	{
		t.POST("/save", tenant.Save)
		t.POST("/del", tenant.Del)
		t.GET("/info", tenant.Info)
		t.GET("/list", tenant.List)
	}
}
