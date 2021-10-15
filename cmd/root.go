package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(produceCmd)
	rootCmd.AddCommand(consumeCmd)
	rootCmd.AddCommand(testPCmd)
	rootCmd.AddCommand(testCCmd)
}
