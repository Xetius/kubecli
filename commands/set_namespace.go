package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/xetius/kubecli/config"
)

func set_namespace(c *cli.Context) {
	namespace, err := getArg(c, 1)
	if err != nil {
		panic(err.Error())
	}

	config.Config.Namespace = namespace
}

func getArg(c *cli.Context, count int) (string, error) {
	if len(c.Args()) != 1 {
		return "", fmt.Errorf("Command requires only one parameter")
	}

	return c.Args()[0], nil
}
