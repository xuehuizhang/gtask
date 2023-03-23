package task

import "github.com/robfig/cron/v3"

var (
	//global task
	GT *Task
)

type Task struct {
	GCron    *GCron    `json:"g_cron"`
	JobPool  *JobPool  `json:"job_pool"`
	RulePool *RulePool `json:"rule_pool"`
}

func InitTask() {
	GT = &Task{
		GCron:    InitGCron(),
		JobPool:  InitJobPool(),
		RulePool: InitRulePool(),
	}

	//初始化Task并运行
	GT.GCron.Run()
}

func GetJob(name string) (*Job, error) {
	return GT.JobPool.Get(name)
}

func DelJob(name string) error {
	return GT.JobPool.Del(name)
}

func SaveJob(j *Job) {
	GT.JobPool.SaveOrUpdate(j)
}

func GetRule(name string) (*Rule, error) {
	return GT.RulePool.Get(name)
}

func DelRule(name string) error {
	return GT.RulePool.Del(name)
}

func SaveRule(r *Rule) {
	GT.RulePool.SaveOrUpdate(r)
}

func SaveCronJob(j *Job) (cron.EntryID, error) {
	return GT.GCron.AddJob(j)
}

func DelCronJob(id cron.EntryID) {
	GT.GCron.Del(id)
}
