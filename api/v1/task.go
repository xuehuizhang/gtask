package v1

import (
	"codego/gtask/setting"
	"codego/gtask/task"
	"codego/gtask/util/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
)

type TaskVo struct {
	EntryId cron.EntryID `json:"entry_id"`
	Name    string       `json:"name"`
	Spec    string       `json:"spec"`
	Stop    bool         `json:"stop"`
}

type TaskBo struct {
	Name string `json:"name"`
	Spec string `json:"spec"`
	Stop bool   `json:"stop"`
}

func Add(ctx *gin.Context) {
	appG := app.AppG{ctx}
	bo := new(TaskBo)
	err := ctx.BindJSON(&bo)
	if err != nil {
		appG.ResponseError("Invalid Param")
		return
	}

	_, ok := setting.TaskMap[bo.Name]
	if !ok {
		rule := setting.TaskRule{
			Name:    bo.Name,
			Spec:    bo.Spec,
			Stop:    false,
			EntryId: 0,
			TaskFunc: func() {
				resp, _ := http.Get("https://baidu.com")
				defer resp.Body.Close()
				all, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(all))
			},
		}

		entryId := task.Add(rule)

		rule.EntryId = entryId

		setting.TaskMap[bo.Name] = rule
	}
}

func Del(ctx *gin.Context) {

}

func Info(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	name := appG.C.Query("name")

	rule, ok := setting.TaskMap[name]
	if !ok {
		appG.ResponseError("任务不存在")
		return
	}

	vo := TaskVo{
		EntryId: rule.EntryId,
		Name:    rule.Name,
		Spec:    rule.Spec,
		Stop:    false,
	}

	appG.ResponseOk(vo)
}

func Stop(ctx *gin.Context) {

}

func Run(ctx *gin.Context) {
	appG := app.AppG{C: ctx}
	name := appG.C.Query("name")

	rule, ok := setting.TaskMap[name]
	if !ok {
		appG.ResponseError("任务不存在")
		return
	}

	go rule.TaskFunc()

	appG.ResponseOk("success")
}
