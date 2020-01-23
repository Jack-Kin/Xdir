package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Name					string
	StockDictUrl			string
	AppServerAddr			string
}

func GetConfig()(conf Config) {
	// use viper to read the config file
	viper.SetConfigName("config")
	viper.AddConfigPath("./req_resp/conf")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	conf.Name = viper.GetString("name")
	conf.StockDictUrl = viper.GetString("stockdict_url")
	conf.AppServerAddr = viper.GetString("app_server_addr")

	return conf
}