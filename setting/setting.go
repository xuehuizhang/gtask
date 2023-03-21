package setting

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

var (
	DbConfigSetting = DbConfig{}
	TaskMap         = map[string]TaskRule{}
)

func SetUp() {
	initConfig()
}

type TaskFunc func()

type TaskRule struct {
	Name     string       `json:"name" yaml:"name"`
	Spec     string       `json:"spec" yaml:"spec"`
	Stop     bool         `json:"stop" yaml:"stop"`
	EntryId  cron.EntryID `json:"entry_id"`
	TaskFunc TaskFunc     `json:"task_func"`
}

func initConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("tasks", &TaskMap)
	if err != nil {
		fmt.Println(err)
	}
}
