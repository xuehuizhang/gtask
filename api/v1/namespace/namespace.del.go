package namespace

import (
	"codego/gtask/service/namespaceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type NamespaceDelBo struct {
	Id int64 `json:"id"`
}

// @Summary  删除Namespace
// @Tags namespace
// @Accept json
// @Produce  json
// @Param   body  body   NamespaceDelBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/namespace/del [POST]
func Del(ctx *gin.Context) {
	appG := app.AppG{C: ctx}

	bo := new(NamespaceDelBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Params")
		return
	}

	res := namespaceService.DelNamespace(bo.Id)
	if !res {
		appG.ResponseError("删除失败")
		return
	}
	appG.ResponseOk("success")
}
