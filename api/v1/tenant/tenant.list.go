package tenant

import (
	"codego/gtask/service/tenantService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type TenantListVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  Job列表
// @Tags 租户
// @Accept json
// @Produce  json
// @Param name query string false "租户名称"
// @Param pageNo query int64 true "页码"
// @Param pageSize query int64 false "每页数量"
// @Success 200 {object} TenantListVo true ""
// @Failure 400 {string} string
// @Router /api/tenant/list  [GET]
func List(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	name := appG.C.Query("name")

	offSet, pageSize := appG.GetPage()

	list := tenantService.ListTenant(name, offSet, pageSize)

	respList := make([]TenantListVo, 0)
	for _, t := range list {
		respList = append(respList, TenantListVo{
			Id:   t.Id,
			Name: t.Name,
		})
	}
	appG.ResponseOk(respList)
}
