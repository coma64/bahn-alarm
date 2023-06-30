package cmd

import (
	"github.com/coma64/bahn-alarm-backend/cmd/serve"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bahn-alarm",
	Short: "Bahn Alarm Backend",
	Long:  "The bahn alarm API server and management commands",
}

func Execute() {
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serve.Cmd)
}
