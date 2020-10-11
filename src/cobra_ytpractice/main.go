package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "example",
		Short: "An example of cobra program",
		Long: `This is a simple example of a cobra program.
		It will have several subcommand and flags.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from the room command")
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
