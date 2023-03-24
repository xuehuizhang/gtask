package tenant

import (
	"codego/gtask/service/tenantService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type TenantDelBo struct {
	Id int64 `json:"id"`
}

// @Summary  删除租户
// @Tags 租户
// @Accept json
// @Produce  json
// @Param   body  body   TenantDelBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/tenant/del [POST]
func Del(ctx *gin.Context) {
	appG := app.AppG{C: ctx}

	bo := new(TenantDelBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Params")
		return
	}

	res := tenantService.DelTenant(bo.Id)
	if !res {
		appG.ResponseError("删除失败")
		return
	}
	appG.ResponseOk("success")
}
