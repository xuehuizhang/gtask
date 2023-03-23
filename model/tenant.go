package model

type Tenant struct {
	Base
	Name string `json:"name"`
}

func (Tenant) TableName() string {
	return "tenant"
}
