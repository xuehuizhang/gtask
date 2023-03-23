package job

import (
	"codego/gtask/task"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type JobRunBo struct {
	Name string `json:"name"`
}

// @Summary  运行Job
// @Tags JOB
// @Accept json
// @Produce  json
// @Param   body  body   JobRunBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/job/run  [POST]
func Run(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	bo := new(JobRunBo)
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

	if job.Stop == true {
		appG.ResponseError("任务已经暂停")
		return
	}

	go job.JobFunc(job.JobRule)

	appG.ResponseOk("success")
}
