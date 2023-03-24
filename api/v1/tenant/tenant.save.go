package tenant

import (
	"codego/gtask/service/tenantService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type TenantSaveBo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  保存租户
// @Tags 租户
// @Accept json
// @Produce  json
// @Param   body  body   TenantSaveBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/tenant/save  [POST]
func Save(ctx *gin.Context) {
	appG := app.AppG{C: ctx}

	bo := new(TenantSaveBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Params")
		return
	}

	res := tenantService.SaveTenant(bo.Id, bo.Name)
	if !res {
		appG.ResponseError("保存失败")
		return
	}
	appG.ResponseOk("success")
}
