package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run pulpod",
	Run: func(_ *cobra.Command, _ []string) {
		slog.Info("run pulpod, run")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
