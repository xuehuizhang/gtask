package service

import (
	"codego/gtask/service/serviceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type ServiceDelBo struct {
	Id int64 `json:"id"`
}

// @Summary  删除服务
// @Tags 服务
// @Accept json
// @Produce  json
// @Param   body  body   ServiceDelBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/service/del [POST]
func Del(ctx *gin.Context) {
	appG := app.AppG{C: ctx}

	bo := new(ServiceDelBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Params")
		return
	}

	res := serviceService.DelService(bo.Id)
	if !res {
		appG.ResponseError("删除失败")
		return
	}
	appG.ResponseOk("success")
}
