package serviceService

import (
	"codego/gtask/dao/serviceDao"
	"codego/gtask/model"
	"codego/gtask/util/enum"
	"go.uber.org/zap"
)

func SaveService(id int64, tenantId int64, name string) bool {
	info, _ := serviceDao.GetOne(id)

	if info != nil {
		info.Name = name
	} else {
		info = &model.Service{
			Base: model.Base{
				Status: 1,
			},
			Name:     name,
			TenantId: tenantId,
		}
	}
	err := serviceDao.SaveOrUpdate(info)
	if err != nil {
		zap.S().Errorf("保存服务失败：%v", err.Error())
		return false
	}
	return true
}

func UpdateService(id int64, name string) bool {
	err := serviceDao.UpdateMap(id, map[string]interface{}{"name": name})
	if err != nil {
		zap.S().Errorf("更新服务：%v，失败：%v", id, err.Error())
		return false
	}
	return true
}

func GetService(id int64) *model.Service {
	tenant, err := serviceDao.GetOne(id)
	if err == enum.ErrRecordNotFound {
		return nil
	}

	if err != nil {
		zap.S().Errorf("获取服务失败：%v", err.Error())
		return nil
	}
	return tenant
}

func DelService(id int64) bool {
	err := serviceDao.Delete(id)
	if err != nil {
		zap.S().Errorf("删除服务失败：%v", err.Error())
		return false
	}
	return true
}

func ListService(name string, offset int, pageSize int) []*model.Service {
	list, err := serviceDao.Page([]string{"id", "name"}, "name like ?", []interface{}{"%" + name + "%"}, "id desc", offset, pageSize)
	if err != nil {
		zap.S().Errorf("获取服务列表失败：%v", err)
	}
	return list
}
