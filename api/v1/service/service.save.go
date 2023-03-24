package service

import (
	"codego/gtask/service/serviceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type ServiceSaveBo struct {
	Id       int64  `json:"id"`
	TenantId int64  `json:"tenant_id"`
	Name     string `json:"name"`
}

// @Summary  保存服务
// @Tags 服务
// @Accept json
// @Produce  json
// @Param   body  body   ServiceSaveBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/service/save  [POST]
func Save(ctx *gin.Context) {
	appG := app.AppG{C: ctx}

	bo := new(ServiceSaveBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Params")
		return
	}

	res := serviceService.SaveService(bo.Id, bo.TenantId, bo.Name)
	if !res {
		appG.ResponseError("保存失败")
		return
	}
	appG.ResponseOk("success")
}
