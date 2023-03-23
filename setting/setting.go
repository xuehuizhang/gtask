package setting

import (
	"github.com/spf13/viper"
)

var (
	DbConfigSetting = DbConfig{}
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//err = viper.UnmarshalKey("rules", &task.GT.RulePool.M)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
