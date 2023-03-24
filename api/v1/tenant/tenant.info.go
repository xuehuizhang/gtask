package tenant

import (
	"codego/gtask/service/tenantService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TenantInfoVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  Job详情
// @Tags 租户
// @Accept json
// @Produce  json
// @Param id query int64 false "租户id"
// @Success 200 {object} TenantInfoVo true ""
// @Failure 400 {string} string
// @Router /api/tenant/info  [GET]
func Info(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	strId := appG.C.Query("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		appG.ResponseError("参数异常")
		return
	}

	tenant := tenantService.GetTenant(int64(id))
	if tenant == nil {
		appG.ResponseError("租户不存在")
		return
	}
	vo := TenantInfoVo{
		Id:   tenant.Id,
		Name: tenant.Name,
	}

	appG.ResponseOk(vo)
}
