package cmd

import (
	"fmt"
	"github.com/UpskillBD/BE-TrainerDash/api"
	"github.com/fsnotify/fsnotify"
	cfg "github.com/UpskillBD/BE-TrainerDash/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)


var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "initializes config by running viper",
	Long:  `users viper to get api keys and db credientials`,
	Run: func(cmd *cobra.Command, args []string) {
		//set up viper
		fmt.Println("Setup viper")
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("..")     // optionally look for config in the working directory
		err := viper.ReadInConfig()   // Find a the config file
		if err != nil {               // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		err = viper.Unmarshal(&cfg.Config)
		if err != nil {
			fmt.Println("err")
		}
		changed := false
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			err = viper.Unmarshal(&cfg.Config)
			if err != nil {
				fmt.Println("err")
			}
			changed = true
		})
		api.RunServer(":9004")
	},
}

func init() {
	rootCmd.AddCommand(dbup)
}

var dbup = &cobra.Command{
	Use:   "dbup",
	Short: "upload schema or schema and seed",
	Long:  "upload schema or schema and seed",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("schema up")
		if len(args) != 0 && args[0] == "seed" {
			fmt.Println("schema and seed up")

		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
