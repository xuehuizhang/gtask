package initialize

import (
	"codego/gtask/setting"
	"codego/gtask/task"
	"codego/gtask/util/client"
	"codego/gtask/util/glog"
)

func Init() {
	//初始化日志
	glog.InitZap()

	//初始化配置
	setting.InitConfig()

	//初始化Client
	client.InitClient()

	//初始化RulePool
	task.InitRulePool()

	//初始化定时任务
	go task.InitTask()
}
