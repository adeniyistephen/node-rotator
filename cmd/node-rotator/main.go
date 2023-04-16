package main

import (
	"encoding/json"
	"os"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error("command failed")
		os.Exit(1)
	}
}

func printJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data)
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rollctl",
		Short: "node-rotator is a cli-tool to rotate K8s cluster nodes",
	}
	rootCmd.AddCommand(newRotateCmd())
	return rootCmd
}
