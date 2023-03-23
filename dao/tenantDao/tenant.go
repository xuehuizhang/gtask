package tenantDao

import (
	"codego/gtask/dao"
	"codego/gtask/model"
	"fmt"
)

func Add(Tenant *model.Tenant) error {
	return dao.TDb.Create(Tenant).Error
}

func Delete(id int64) error {
	return dao.TDb.Delete(&model.Tenant{}, id).Error
}

func UpdateFiled(filedName string, value interface{}) error {
	return dao.TDb.Model(&model.Tenant{}).Update(filedName, value).Error
}

func UpdateMap(m map[string]interface{}) error {
	return dao.TDb.Model(&model.Tenant{}).Updates(m).Error
}

func UpdatesObj(c *model.Tenant) error {
	return dao.TDb.Model(&model.Tenant{}).Updates(c).Error
}

func GetOne(id int64) (*model.Tenant, error) {
	var Tenant *model.Tenant
	err := dao.TDb.First(&Tenant, id).Error
	if err != nil {
		return nil, err
	}
	return Tenant, nil
}

func GetByField(fieldName string, value interface{}) (*model.Tenant, error) {
	var tenant *model.Tenant
	err := dao.TDb.First(&tenant, fmt.Sprintf("%v= ?", fieldName), value).Error
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func Raw(sql string, params []interface{}) ([]*model.Tenant, error) {
	var list []*model.Tenant
	err := dao.TDb.Raw(sql, params...).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Count(sql string, params []interface{}) (int64, error) {
	var count int64
	err := dao.TDb.Model(&model.Tenant{}).Where(sql, params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
