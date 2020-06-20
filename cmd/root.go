package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmdcli",
	Short: "A command line tool.",
	Long:  `CmdCli is a command line tool based on go language.`,
}

func Execute() error {
	return rootCmd.Execute()
}
