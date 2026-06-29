package app

import (
	"fmt"

	"github.com/linc593823915/atlas/internal/version"
	"github.com/spf13/cobra"
)

const VersionTemplate = `Atlas
Version:%s
Commit:%s
BuildAt:%s
`

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of atlas",
	Long:  `All software has versions. This is atlas's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(VersionTemplate, version.Version, version.Commit, version.BuildAt)
	},
}
