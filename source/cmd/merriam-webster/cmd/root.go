package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set by main.go
var Version string

// RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   "merriam-webster",
	Short: "A CLI for querying the Merriam-Webster dictionary",
}

// Execute runs the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
