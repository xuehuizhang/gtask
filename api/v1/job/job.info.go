package job

import (
	"codego/gtask/task"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type JobInfoVo struct {
	EntryId int    `json:"entry_id"`
	Name    string `json:"name"`
	Spec    string `json:"spec"`
	Stop    bool   `json:"stop"`
	Cate    int    `json:"cate"`
	Method  string `json:"method"`
	Params  string `json:"params"`
	Url     string `json:"url"`
}

// @Summary  Job详情
// @Tags JOB
// @Accept json
// @Produce  json
// @Param name query string false "job名称"
// @Success 200 {object} JobInfoVo true ""
// @Failure 400 {string} string
// @Router /api/job/info  [GET]
func Info(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	name := appG.C.Query("name")

	job, err := task.GetJob(name)
	if err != nil {
		appG.ResponseError("任务不存在")
		return
	}

	vo := JobInfoVo{
		EntryId: int(job.EntryId),
		Name:    job.JobRule.Name,
		Spec:    job.JobRule.Spec,
		Stop:    job.Stop,
		Cate:    job.JobRule.Cate,
		Method:  job.JobRule.Method,
		Url:     job.JobRule.Url,
		Params:  job.JobRule.Params,
	}

	appG.ResponseOk(vo)
}
