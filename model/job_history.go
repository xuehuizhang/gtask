package model

type JobHistory struct {
	Base
	JobId   int64 `json:"job_id"`
	RuleId  int64 `json:"rule_id"`
	Code    int   `json:"code"`
	RunTime int64 `json:"run_time"`
	RunHost int64 `json:"run_host"`
}

func (JobHistory) TableName() string {
	return "job_history"
}
