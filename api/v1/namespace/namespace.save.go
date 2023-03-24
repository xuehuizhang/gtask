package namespace

import (
	"codego/gtask/service/namespaceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type NamespaceSaveBo struct {
	Id        int64  `json:"id"`
	ServiceId int64  `json:"service_id"`
	Name      string `json:"name"`
}

// @Summary  保存namespace
// @Tags namespace
// @Accept json
// @Produce  json
// @Param   body  body   NamespaceSaveBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/namespace/save  [POST]
func Save(ctx *gin.Context) {
	appG := app.AppG{C: ctx}

	bo := new(NamespaceSaveBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Params")
		return
	}

	res := namespaceService.SaveNamespace(bo.Id, bo.ServiceId, bo.Name)
	if !res {
		appG.ResponseError("保存失败")
		return
	}
	appG.ResponseOk("success")
}
