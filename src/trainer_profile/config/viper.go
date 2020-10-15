package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperHelper struct {
	DbString string
	DbName   string
}

var Config ViperHelper
func ViperSetup() {

	viper.SetConfigName("config.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find a the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println("err")
	}
	//changed := false
	//viper.WatchConfig()
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	err = viper.Unmarshal(&Config)
	//	if err != nil {
	//		fmt.Println("err")
	//	}
	//	changed = true
	//})

}