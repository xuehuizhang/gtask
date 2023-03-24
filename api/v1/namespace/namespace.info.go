package namespace

import (
	"codego/gtask/service/namespaceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
	"strconv"
)

type NamespaceInfoVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  namespace详情
// @Tags namespace
// @Accept json
// @Produce  json
// @Param id query int64 false "namespace名称"
// @Success 200 {object} NamespaceInfoVo true ""
// @Failure 400 {string} string
// @Router /api/namespace/info  [GET]
func Info(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	strId := appG.C.Query("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		appG.ResponseError("参数异常")
		return
	}

	namespace := namespaceService.GetNamespace(int64(id))
	if namespace == nil {
		appG.ResponseError("namespace不存在")
		return
	}
	vo := NamespaceInfoVo{
		Id:   namespace.Id,
		Name: namespace.Name,
	}

	appG.ResponseOk(vo)
}
