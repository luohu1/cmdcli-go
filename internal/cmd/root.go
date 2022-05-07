package cmd

import (
	"github.com/luohu1/glctl/internal/cmd/completion"
	"github.com/luohu1/glctl/internal/cmd/print"
	"github.com/luohu1/glctl/internal/cmd/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "glctl",
		Short: "Gitlab command line tool.",
		Long:  `glctl is a command line tool to manage gitlab.`,
	}

	rootCmd.AddCommand(
		print.NewCmdPrint(),
		completion.NewCmdCompletion(),
		version.NewCmdVersion(),
	)

	return rootCmd
}
