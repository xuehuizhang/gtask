package serviceDao

import (
	"codego/gtask/dao"
	"codego/gtask/model"
	"fmt"
)

func Add(Service *model.Service) error {
	return dao.TDb.Create(Service).Error
}

func Delete(id int64) error {
	return dao.TDb.Delete(&model.Service{}, id).Error
}

func UpdateFiled(filedName string, value interface{}) error {
	return dao.TDb.Model(&model.Service{}).Update(filedName, value).Error
}

func UpdateMap(m map[string]interface{}) error {
	return dao.TDb.Model(&model.Service{}).Updates(m).Error
}

func UpdatesObj(c *model.Service) error {
	return dao.TDb.Model(&model.Service{}).Updates(c).Error
}

func GetOne(id int64) (*model.Service, error) {
	var service *model.Service
	err := dao.TDb.First(&service, id).Error
	if err != nil {
		return nil, err
	}
	return service, nil
}

func GetByField(fieldName string, value interface{}) (*model.Service, error) {
	var Service *model.Service
	err := dao.TDb.First(&Service, fmt.Sprintf("%v= ?", fieldName), value).Error
	if err != nil {
		return nil, err
	}
	return Service, nil
}

func Raw(sql string, params []interface{}) ([]*model.Service, error) {
	var list []*model.Service
	err := dao.TDb.Raw(sql, params...).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Count(sql string, params []interface{}) (int64, error) {
	var count int64
	err := dao.TDb.Model(&model.Service{}).Where(sql, params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
