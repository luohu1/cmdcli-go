package main

import (
	"fmt"

	"github.com/luohu1/cmdcli/cmd"
)

func main() {
	fmt.Println("Go tools: CmdCli")

	cmd.Print("package: common, func: Print")

	cmd.Execute()
}
