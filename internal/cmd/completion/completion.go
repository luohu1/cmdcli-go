package completion

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCmdCompletion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completion",
	}

	cmd.AddCommand(newCmdBashCompletion())
	cmd.AddCommand(newCmdZshCompletion())

	return cmd
}

func newCmdBashCompletion() *cobra.Command {
	return &cobra.Command{
		Use:   "bash",
		Short: "Generate bash completion",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd := cmd.Parent()
			rootCmd.GenBashCompletion(os.Stdout)
		},
	}
}

func newCmdZshCompletion() *cobra.Command {
	return &cobra.Command{
		Use:   "zsh",
		Short: "Generate zsh completion",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd := cmd.Parent()
			rootCmd.GenZshCompletion(os.Stdout)
		},
	}
}
