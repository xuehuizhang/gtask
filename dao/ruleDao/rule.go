package ruleDao

import (
	"codego/gtask/dao"
	"codego/gtask/model"
	"fmt"
)

func Add(Rule *model.Rule) error {
	return dao.TDb.Create(Rule).Error
}

func Delete(id int64) error {
	return dao.TDb.Delete(&model.Rule{}, id).Error
}

func UpdateFiled(filedName string, value interface{}) error {
	return dao.TDb.Model(&model.Rule{}).Update(filedName, value).Error
}

func UpdateMap(m map[string]interface{}) error {
	return dao.TDb.Model(&model.Rule{}).Updates(m).Error
}

func UpdatesObj(c *model.Rule) error {
	return dao.TDb.Model(&model.Rule{}).Updates(c).Error
}

func GetOne(id int64) (*model.Rule, error) {
	var rule *model.Rule
	err := dao.TDb.First(&rule, id).Error
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func GetByField(fieldName string, value interface{}) (*model.Rule, error) {
	var Rule *model.Rule
	err := dao.TDb.First(&Rule, fmt.Sprintf("%v= ?", fieldName), value).Error
	if err != nil {
		return nil, err
	}
	return Rule, nil
}

func Raw(sql string, params []interface{}) ([]*model.Rule, error) {
	var list []*model.Rule
	err := dao.TDb.Raw(sql, params...).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Count(sql string, params []interface{}) (int64, error) {
	var count int64
	err := dao.TDb.Model(&model.Rule{}).Where(sql, params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
