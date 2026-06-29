package app

import (
	"context"
	"fmt"

	"github.com/linc593823915/atlas/internal/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve atlas",
	Long:  `serve atlas`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrapInstance, err := bootstrap.New(context.Background())
		if err != nil {
			fmt.Printf("bootstrap new error: %v", err)
			return
		}
		err = bootstrapInstance.Start()
		if err != nil {
			fmt.Printf("bootstrap start error: %v", err)
			return
		}
		fmt.Println("bootstrap start success")
		defer func() {
			err = bootstrapInstance.Stop()
			if err != nil {
				fmt.Printf("bootstrap stop error: %v", err)
				return
			}
			fmt.Println("bootstrap stop success")
		}()
	},
}
