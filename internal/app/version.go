package app

import (
	"context"

	"github.com/linc593823915/atlas/internal/logger"
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
		err := logger.WithOptions(logger.LevelInfo, logger.FormatText)
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		logger.Info(ctx, "runtime", "runtime", "starting")
		logger.Info(ctx, "version", "version", version.Version)
		logger.Info(ctx, "runtime", "runtime", "stopped")
	},
}
