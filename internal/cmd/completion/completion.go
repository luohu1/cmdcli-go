package completion

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCompletionCmd() *cobra.Command {
	cc := &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completion",
	}

	cc.AddCommand(newBashCompletionCmd())
	cc.AddCommand(newZshCompletionCmd())

	return cc
}

func newBashCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bash",
		Short: "Generate bash completion",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd := cmd.Parent()
			rootCmd.GenBashCompletion(os.Stdout)
		},
	}
}

func newZshCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "zsh",
		Short: "Generate zsh completion",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd := cmd.Parent()
			rootCmd.GenZshCompletion(os.Stdout)
		},
	}
}
