package app

import (
	"context"

	"github.com/linc593823915/atlas/internal/config"
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
		manager, err := configureRuntime()
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		logger.Info(ctx, "runtime", "runtime", "starting")
		logger.Info(ctx, "app", "name", manager.App().Name, "version", manager.App().Version)
		logger.Info(ctx, "build", "version", version.Version, "commit", version.Commit, "build_at", version.BuildAt)
		logger.Info(ctx, "runtime", "runtime", "stopped")
	},
}

func configureRuntime() (*config.Manager, error) {
	manager := config.NewManager()
	if err := manager.Load(); err != nil {
		return nil, err
	}
	if err := logger.WithOptions(manager.Logger().Level, manager.Logger().Format); err != nil {
		return nil, err
	}
	return manager, nil
}
