package model

type Rule struct {
	Base
	Name   string `json:"name"`
	Spec   string `json:"spec"`
	Cate   int    `json:"cate"`
	Url    string `json:"url"`
	Method string `json:"method"`
	Params string `json:"params"`
	Dcs    int    `json:"dcs"`
}

func (Rule) TableName() string {
	return "rule"
}
