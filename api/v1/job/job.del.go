package job

import (
	"codego/gtask/task"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type JobDelBo struct {
	Name string `json:"name"`
}

// @Summary  删除Job
// @Tags JOB
// @Accept json
// @Produce  json
// @Param   body  body   JobDelBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/job/del  [POST]
func Del(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	bo := new(JobDelBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Prams")
		return
	}

	name := bo.Name
	//判断job是否存在
	job, err := task.GetJob(name)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	//判断是否在运行 todo

	//删除Rule
	task.DelRule(name)
	//删除Job
	task.DelJob(name)
	//删除Cron
	task.DelCronJob(job.EntryId)
}
