package main

import (
	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "postgres",
		"password": "1234",
		"host":     5432,
		"database": "test",
	}
)

func main() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}
