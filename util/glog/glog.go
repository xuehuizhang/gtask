package glog

import (
	"fmt"
	"go.uber.org/zap"
)

func InitZap() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("初始化日志组件失败:%v", err))
	}
	zap.ReplaceGlobals(logger)
}
