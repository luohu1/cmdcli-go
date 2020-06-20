package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Print(s string) {
	fmt.Println(s)
}

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A generator for Cobra based Applications",
	Long: `Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
}

func Execute() error {
	return rootCmd.Execute()
}
