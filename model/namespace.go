package model

type Namespace struct {
	Base
	ServiceId int64  `json:"service_id"`
	Name      string `json:"name"`
}

func (Namespace) TableName() string {
	return "namespace"
}
