package job

import (
	"codego/gtask/task"
	"codego/gtask/util/app"
	"github.com/gin-gonic/gin"
)

type JobAddBo struct {
	Name   string `json:"name"`
	Spec   string `json:"spec"`
	Cate   int    `json:"cate"`
	Url    string `json:"url"`
	Method string `json:"method"`
	Params string `json:"params"`
}

// @Summary  新增Job
// @Tags JOB
// @Accept json
// @Produce  json
// @Param   body  body   JobAddBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/job/add  [POST]
func Add(ctx *gin.Context) {
	appG := app.AppG{ctx}
	bo := new(JobAddBo)
	err := ctx.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Param")
		return
	}

	_, err = task.GT.JobPool.Get(bo.Name)
	if err != nil {
		rule := &task.Rule{
			Name:   bo.Name,
			Spec:   bo.Spec,
			Url:    bo.Url,
			Method: bo.Method,
			Cate:   bo.Cate,
			Params: bo.Params,
		}

		var f task.JobFunc
		if rule.Cate == task.HTTP_JOB {
			f = task.HttpJobFunc()
		} else if rule.Cate == task.TCP_JOB {
			f = task.TcpJobFunc()
		} else if rule.Cate == task.SH_JOB {
			f = task.ShellJobFunc()
		}

		j := &task.Job{
			JobFunc: f,
			JobRule: rule,
			Stop:    false,
		}

		entryId, err := task.SaveCronJob(j)
		if err != nil {
			appG.ResponseError("添加Job失败")
			return
		}
		j.EntryId = entryId

		task.SaveJob(j)
		task.SaveRule(rule)
	}

	appG.ResponseOk("success")
}
