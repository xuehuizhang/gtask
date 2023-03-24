package router

import (
	_ "codego/gtask/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g := r.Group("/api")

	//Job
	initJobRouter(g)

	//租户
	initTenantRouter(g)

	//service
	initServiceRoute(g)

	//namespace
	initNamespaceRoute(g)
}
