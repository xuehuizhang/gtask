package service

import (
	"codego/gtask/service/serviceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ServiceInfoVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  服务详情
// @Tags 服务
// @Accept json
// @Produce  json
// @Param id query int64 false "服务id"
// @Success 200 {object} ServiceInfoVo true ""
// @Failure 400 {string} string
// @Router /api/service/info  [GET]
func Info(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	strId := appG.C.Query("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		appG.ResponseError("参数异常")
		return
	}

	service := serviceService.GetService(int64(id))
	if service == nil {
		appG.ResponseError("服务不存在")
		return
	}
	vo := ServiceInfoVo{
		Id:   service.Id,
		Name: service.Name,
	}

	appG.ResponseOk(vo)
}
