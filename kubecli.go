package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/xetius/kubecli/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kube-CLI"
	app.Version = "0.1.0"
	app.Usage = "Safety wrapper for kubectl command"
	app.Commands = commands.GetCommands()

	app.Run(os.Args)
}
