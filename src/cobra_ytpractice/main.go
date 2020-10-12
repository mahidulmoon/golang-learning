package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	localRootFlag   bool
	persistRootFlag bool
	times           int
	rootCmd         = &cobra.Command{
		Use:   "example",
		Short: "An example of cobra program",
		Long: `This is a simple example of a cobra program.
		It will have several subcommand and flags.`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("hello from the room command")
		// },
	}
	echoCmd = &cobra.Command{
		Use:   "echo, [string to echo]",
		Short: "prints given string to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, ""))
		},
	}
	timesCmd = &cobra.Command{
		Use:   "times [strings to echo]",
		Short: "prints given string to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "apersistent root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
