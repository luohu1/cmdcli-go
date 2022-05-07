package cmd

import (
	"github.com/luohu1/glctl/internal/cmd/completion"
	"github.com/luohu1/glctl/internal/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glctl",
	Short: "Gitlab command line tool.",
	Long:  `glctl is a command line tool to manage gitlab.`,
}

func init() {
	rootCmd.AddCommand(
		completion.NewCompletionCmd(),
		version.NewVersionCmd(),
	)
}

func Execute() error {
	return rootCmd.Execute()
}
