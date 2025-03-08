package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "coop",
	Short: "Estellencs.coop resource management",
}

func Execute() error {
	return rootCmd.Execute()
}
