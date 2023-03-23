package router

import (
	v1 "codego/gtask/api/v1/job"
	"github.com/gin-gonic/gin"
)

func InitTaskRouter(r *gin.RouterGroup) {
	t := r.Group("/job")
	{
		t.POST("/add", v1.Add)
		t.POST("/del", v1.Del)
		t.POST("/stop", v1.Stop)
		t.POST("/run", v1.Run)

		t.GET("/info", v1.Info)
	}
}
