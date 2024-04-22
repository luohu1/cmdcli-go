package cmd

import (
	"github.com/luohu1/osctl/internal/cmd/completion"
	"github.com/luohu1/osctl/internal/cmd/print"
	"github.com/luohu1/osctl/internal/cmd/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "osctl",
		Short: "Gitlab command line tool.",
		Long:  `osctl is a command line tool to manage gitlab.`,
	}

	rootCmd.AddCommand(
		print.NewCmdPrint(),
		completion.NewCmdCompletion(),
		version.NewCmdVersion(),
	)

	return rootCmd
}
