package namespace

import (
	"codego/gtask/service/namespaceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type NamespaceListVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  namespace详情
// @Tags namespace
// @Accept json
// @Produce  json
// @Param name query string false "namespace名称"
// @Param pageNo query int64 true "页码"
// @Param pageSize query int64 false "每页数量"
// @Success 200 {object} NamespaceListVo true ""
// @Failure 400 {string} string
// @Router /api/namespace/list  [GET]
func List(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	name := appG.C.Query("name")

	offSet, pageSize := appG.GetPage()

	list := namespaceService.ListNamespace(name, offSet, pageSize)

	respList := make([]NamespaceListVo, 0)
	for _, t := range list {
		respList = append(respList, NamespaceListVo{
			Id:   t.Id,
			Name: t.Name,
		})
	}
	appG.ResponseOk(respList)
}
