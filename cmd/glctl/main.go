package main

import (
	"os"

	"github.com/luohu1/glctl/internal/cmd"
)

func main() {
	command := cmd.NewRootCmd()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
