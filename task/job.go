package task

import (
	"errors"
	"github.com/robfig/cron/v3"
	"sync"
)

type JobPool struct {
	m    map[string]*Job
	lock sync.RWMutex
}

func InitJobPool() *JobPool {
	return &JobPool{
		m:    make(map[string]*Job),
		lock: sync.RWMutex{},
	}
}

func (j *JobPool) SaveOrUpdate(job *Job) {
	j.lock.Lock()
	defer j.lock.Unlock()
	j.m[job.JobRule.Name] = job
}

func (j *JobPool) Get(name string) (*Job, error) {
	j.lock.RLock()
	defer j.lock.RUnlock()
	job, ok := j.m[name]
	if ok {
		return job, nil
	}
	return nil, errors.New("un found")
}

func (j *JobPool) Del(name string) error {
	j.lock.Lock()
	defer j.lock.Unlock()
	delete(j.m, name)
	return nil
}

type Job struct {
	JobFunc JobFunc      `json:"task_func"`
	JobRule *Rule        `json:"job_rule"`
	EntryId cron.EntryID `json:"entry_id"`
	Stop    bool         `json:"stop" yaml:"stop"` //是否暂停
}

func (t *Job) Run() {
	if t.Stop {
		return
	}
	t.JobFunc(t.JobRule)
}
