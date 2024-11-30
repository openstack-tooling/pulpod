package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "pulpod",
	Short: "A CI orchestrator tool for RHOSO (Red Hat OpenStack Services on OpenShift)",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
