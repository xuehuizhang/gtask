package jobDao

import (
	"codego/gtask/dao"
	"codego/gtask/model"
	"fmt"
)

func Add(job *model.Job) error {
	return dao.TDb.Create(job).Error
}

func Delete(id int64) error {
	return dao.TDb.Delete(&model.Job{}, id).Error
}

func UpdateFiled(filedName string, value interface{}) error {
	return dao.TDb.Model(&model.Job{}).Update(filedName, value).Error
}

func UpdateMap(m map[string]interface{}) error {
	return dao.TDb.Model(&model.Job{}).Updates(m).Error
}

func UpdatesObj(c *model.Job) error {
	return dao.TDb.Model(&model.Job{}).Updates(c).Error
}

func GetOne(id int64) (*model.Job, error) {
	var job *model.Job
	err := dao.TDb.First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func GetByField(fieldName string, value interface{}) (*model.Job, error) {
	var Job *model.Job
	err := dao.TDb.First(&Job, fmt.Sprintf("%v= ?", fieldName), value).Error
	if err != nil {
		return nil, err
	}
	return Job, nil
}

func Raw(sql string, params []interface{}) ([]*model.Job, error) {
	var list []*model.Job
	err := dao.TDb.Raw(sql, params...).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Count(sql string, params []interface{}) (int64, error) {
	var count int64
	err := dao.TDb.Model(&model.Job{}).Where(sql, params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
