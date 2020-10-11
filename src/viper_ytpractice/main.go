package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "postgres",
		"password": "1234",
		"host":     5432,
		"database": "test",
	}
	configName  = "config"
	configPaths = []string{
		".",
	}
)

func main() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}
	fmt.Println("Username from viper: ", viper.GetString("username"))
	fmt.Println("Password from viper: ", viper.GetString("password"))
	fmt.Println("Host from viper: ", viper.GetString("host"))
	fmt.Println("Port from viper: ", viper.GetInt("port"))
}
