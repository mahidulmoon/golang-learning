package cmd

import (

	"fmt"
	"github.com/UpskillBD/BE-TrainerDash/api"
	"github.com/spf13/cobra"
	"os"
)


var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "initializes config by running viper",
	Long:  `users viper to get api keys and db credientials`,
	Run: func(cmd *cobra.Command, args []string) {

		api.RunServer(":9004")
	},
}

func init() {

	//rootCmd.AddCommand(dbup)
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
