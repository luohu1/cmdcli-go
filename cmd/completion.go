package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(bashCompletionCmd)
	completionCmd.AddCommand(zshCompletionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion",
}

var bashCompletionCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate bash completion",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

var zshCompletionCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate zsh completion",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}
