package router

import (
	v1 "codego/gtask/api/v1"
	"github.com/gin-gonic/gin"
)

func InitTaskRouter(r *gin.Engine) {
	t := r.Group("/task")
	{
		t.POST("/add", v1.Add)
		t.POST("/del", v1.Del)
		t.GET("/info", v1.Info)
		t.POST("/stop", v1.Stop)
		t.GET("/run", v1.Run)
	}
}
