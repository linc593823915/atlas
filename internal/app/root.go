package app

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "atlas",
	Short: "atlas  is a ai agent os platform",
	Long:  `atlas  is a ai agent os platform`,
}

func Execute() error {
	return rootCmd.Execute()
}

func Run() int {
	if err := Execute(); err != nil {
		return 1
	}
	return 0
}
