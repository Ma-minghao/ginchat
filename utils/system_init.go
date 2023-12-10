package utils

import(
	"github.com/spf13/viper"
)


func InitConfig(){
	viper.SetConfigName( "app")
	viper.AddConfigPath("ginchat/config")

}

func InitMySQL(){

}