package task

import "github.com/robfig/cron/v3"

type GCron struct {
	c *cron.Cron
}

func InitGCron() *GCron {
	return &GCron{c: cron.New()}
}

func (g *GCron) AddJob(job *Job) (cron.EntryID, error) {
	return g.c.AddJob(job.JobRule.Spec, job)
}

func (g *GCron) Del(entryId cron.EntryID) {
	g.c.Remove(entryId)
}

func (g *GCron) Run() {
	g.c.Run()
}
