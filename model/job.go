package model

type Job struct {
	Base
	RuleId int64  `json:"rule_id"`
	Name   string `json:"name"`
	Stop   int    `json:"stop"`
}

func (Job) TableName() string {
	return "job"
}
