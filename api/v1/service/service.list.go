package service

import (
	"codego/gtask/service/serviceService"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type ServiceListVo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary  服务列表
// @Tags 服务
// @Accept json
// @Produce  json
// @Param name query string false "服务名称"
// @Param pageNo query int64 true "页码"
// @Param pageSize query int64 false "每页数量"
// @Success 200 {object} ServiceListVo true ""
// @Failure 400 {string} string
// @Router /api/service/list  [GET]
func List(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	name := appG.C.Query("name")

	offSet, pageSize := appG.GetPage()

	list := serviceService.ListService(name, offSet, pageSize)

	respList := make([]ServiceListVo, 0)
	for _, t := range list {
		respList = append(respList, ServiceListVo{
			Id:   t.Id,
			Name: t.Name,
		})
	}
	appG.ResponseOk(respList)
}
