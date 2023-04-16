package main

import (
	"os"

	"github.com/spf13/cobra"
	logger "github.com/sirupsen/logrus"
)

func main() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error("command failed")
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rollctl",
		Short: "node-rotator is a cli-tool to rotate K8s cluster nodes",
	}
	rootCmd.AddCommand(newRotateCmd())
	return rootCmd
}