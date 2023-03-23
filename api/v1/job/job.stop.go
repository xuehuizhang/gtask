package job

import (
	"codego/gtask/task"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type JobStopBo struct {
	Name string `json:"name"`
}

// @Summary  暂停Job
// @Tags JOB
// @Accept json
// @Produce  json
// @Param   body  body   JobStopBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/job/stop  [POST]
func Stop(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	bo := new(JobStopBo)
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Prams")
		return
	}

	name := bo.Name

	job, err := task.GetJob(name)
	if err != nil {
		appG.ResponseError("任务不存在")
		return
	}

	job.Stop = true

	//job pool中设置为暂停状态
	task.SaveJob(job)

	//cron中删除
	task.DelCronJob(job.EntryId)

	appG.ResponseOk("success")
}
