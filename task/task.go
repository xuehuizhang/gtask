package task

import "C"
import (
	"codego/gtask/setting"
	"fmt"
	"github.com/robfig/cron/v3"
)

var (
	c *cron.Cron
)

type TaskJob struct {
	Name     string           `json:"name"`
	TaskFunc setting.TaskFunc `json:"task_func"`
}

func (t *TaskJob) Run() {
	t.TaskFunc()
}

func InitTask() {
	c = cron.New()

	for _, val := range setting.TaskMap {
		v := val

		f := func() {
			fmt.Println(v.Name)
		}

		job := &TaskJob{
			Name:     v.Name,
			TaskFunc: f,
		}

		entryId, _ := c.AddJob(v.Spec, job)

		rule := setting.TaskMap[v.Name]
		rule.EntryId = entryId
		rule.TaskFunc = f
		setting.TaskMap[v.Name] = rule
	}

	c.Run()
}

func Add(rule setting.TaskRule) cron.EntryID {
	entryId, _ := c.AddFunc(rule.Spec, rule.TaskFunc)
	return entryId
}

func Remove(entryId cron.EntryID) {
	c.Remove(entryId)
}
