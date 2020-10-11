package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
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

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

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

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Couldnot decode config into struct: %v", err)
	}
	fmt.Println("Username from struct: ", config.Username)
	fmt.Println("Config struct: ", config)
	changed := false
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&config)
		if err != nil {
			log.Printf("could not decode config after changed : %v", err)
		}
		changed = true
	})
	for !changed {
		time.Sleep(time.Second)
		fmt.Println("Config struct: ", config)
	}
}
