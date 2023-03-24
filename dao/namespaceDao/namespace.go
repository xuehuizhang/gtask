package namespaceDao

import (
	"codego/gtask/dao"
	"codego/gtask/model"
	"fmt"
)

func Add(Namespace *model.Namespace) error {
	return dao.TDb.Create(Namespace).Error
}

func Delete(id int64) error {
	return dao.TDb.Delete(&model.Namespace{}, id).Error
}

func UpdateFiled(filedName string, value interface{}) error {
	return dao.TDb.Model(&model.Namespace{}).Update(filedName, value).Error
}

func UpdateMap(id int64, m map[string]interface{}) error {
	return dao.TDb.Model(&model.Namespace{}).Where("id=?", id).Updates(m).Error
}

func UpdatesObj(c *model.Namespace) error {
	return dao.TDb.Model(&model.Namespace{}).Updates(c).Error
}

func SaveOrUpdate(namespace *model.Namespace) error {
	return dao.TDb.Save(namespace).Error
}

func GetOne(id int64) (*model.Namespace, error) {
	var namespace *model.Namespace
	err := dao.TDb.First(&namespace, id).Error
	if err != nil {
		return nil, err
	}
	return namespace, nil
}

func GetByField(fieldName string, value interface{}) (*model.Namespace, error) {
	var namespace *model.Namespace
	err := dao.TDb.First(&namespace, fmt.Sprintf("%v= ?", fieldName), value).Error
	if err != nil {
		return nil, err
	}
	return namespace, nil
}

func Raw(sql string, params []interface{}) ([]*model.Namespace, error) {
	var list []*model.Namespace
	err := dao.TDb.Raw(sql, params...).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Count(sql string, params []interface{}) (int64, error) {
	var count int64
	err := dao.TDb.Model(&model.Namespace{}).Where(sql, params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func Page(selectField []string, sql string, params []interface{}, order string, offset, pageSize int) ([]*model.Namespace, error) {
	var posts []*model.Namespace
	err := dao.TDb.Select(selectField).Where(sql, params...).Order(order).Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, err
}
