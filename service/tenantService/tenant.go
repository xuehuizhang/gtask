package tenantService

import (
	"codego/gtask/dao/tenantDao"
	"codego/gtask/model"
	"codego/gtask/util/enum"
	"go.uber.org/zap"
)

func SaveTenant(id int64, name string) bool {
	info, _ := tenantDao.GetOne(id)

	if info != nil {
		info.Name = name
	} else {
		info = &model.Tenant{
			Base: model.Base{
				Status: 1,
			},
			Name: name,
		}
	}
	err := tenantDao.SaveOrUpdate(info)
	if err != nil {
		zap.S().Errorf("保存租户失败：%v", err.Error())
		return false
	}
	return true
}

func UpdateTenant(id int64, name string) bool {
	err := tenantDao.UpdateMap(id, map[string]interface{}{"name": name})
	if err != nil {
		zap.S().Errorf("更新租户：%v，失败：%v", id, err.Error())
		return false
	}
	return true
}

func GetTenant(id int64) *model.Tenant {
	tenant, err := tenantDao.GetOne(id)
	if err == enum.ErrRecordNotFound {
		return nil
	}

	if err != nil {
		zap.S().Errorf("获取租户失败：%v", err.Error())
		return nil
	}
	return tenant
}

func DelTenant(id int64) bool {
	err := tenantDao.Delete(id)
	if err != nil {
		zap.S().Errorf("删除租户失败：%v", err.Error())
		return false
	}
	return true
}

func ListTenant(name string, offset int, pageSize int) []*model.Tenant {
	list, err := tenantDao.Page([]string{"id", "name"}, "name like ?", []interface{}{"%" + name + "%"}, "id desc", offset, pageSize)
	if err != nil {
		zap.S().Errorf("获取租户列表失败：%v", err)
	}
	return list
}
