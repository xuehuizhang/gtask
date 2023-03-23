package model

type Service struct {
	Base
	TenantId int64  `json:"tenant_id"`
	Name     string `json:"name"`
}

func (Service) TableName() string {
	return "service"
}
