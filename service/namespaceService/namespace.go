package namespaceService

import (
	"codego/gtask/dao/namespaceDao"
	"codego/gtask/model"
	"codego/gtask/util/enum"
	"go.uber.org/zap"
)

func SaveNamespace(id int64, serviceId int64, name string) bool {
	info, _ := namespaceDao.GetOne(id)

	if info != nil {
		info.Name = name
	} else {
		info = &model.Namespace{
			Base: model.Base{
				Status: 1,
			},
			Name:      name,
			ServiceId: serviceId,
		}
	}
	err := namespaceDao.SaveOrUpdate(info)
	if err != nil {
		zap.S().Errorf("保存Namespace失败：%v", err.Error())
		return false
	}
	return true
}

func UpdateNamespace(id int64, name string) bool {
	err := namespaceDao.UpdateMap(id, map[string]interface{}{"name": name})
	if err != nil {
		zap.S().Errorf("更新Namespace：%v，失败：%v", id, err.Error())
		return false
	}
	return true
}

func GetNamespace(id int64) *model.Namespace {
	tenant, err := namespaceDao.GetOne(id)
	if err == enum.ErrRecordNotFound {
		return nil
	}

	if err != nil {
		zap.S().Errorf("获取Namespace失败：%v", err.Error())
		return nil
	}
	return tenant
}

func DelNamespace(id int64) bool {
	err := namespaceDao.Delete(id)
	if err != nil {
		zap.S().Errorf("删除Namespace失败：%v", err.Error())
		return false
	}
	return true
}

func ListNamespace(name string, offset int, pageSize int) []*model.Namespace {
	list, err := namespaceDao.Page([]string{"id", "name"}, "name like ?", []interface{}{"%" + name + "%"}, "id desc", offset, pageSize)
	if err != nil {
		zap.S().Errorf("获取Namespace列表失败：%v", err)
	}
	return list
}
