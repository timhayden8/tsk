// Package cmd holds the tsk command-line interface.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// version is overridden at build time: go build -ldflags "-X github.com/timhayden8/tsk/cmd.version=v0.1.0"
var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "tsk",
	Short: "tsk is a small task tracker",
	Long:  "tsk is a small task tracker, built as a playground for learning the Claude Code harness.",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the tsk version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// Execute runs the root command and exits non-zero on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
