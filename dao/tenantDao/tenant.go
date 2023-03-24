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

func UpdateMap(id int64, m map[string]interface{}) error {
	return dao.TDb.Model(&model.Tenant{}).Where("id=?", id).Updates(m).Error
}

func SaveOrUpdate(tenant *model.Tenant) error {
	return dao.TDb.Save(tenant).Error
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

func Page(selectField []string, sql string, params []interface{}, order string, offset, pageSize int) ([]*model.Tenant, error) {
	var posts []*model.Tenant
	err := dao.TDb.Select(selectField).Where(sql, params...).Order(order).Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, err
}
