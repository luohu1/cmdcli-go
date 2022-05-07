package print

import (
	"github.com/spf13/cobra"
)

func NewCmdPrint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print",
		Short: "subcommand print",
	}

	cmd.AddCommand(newCmdPrintClone())

	return cmd
}
