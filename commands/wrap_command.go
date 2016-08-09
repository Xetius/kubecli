package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func wrap_command(c *cli.Context) {
	arguments := ""
	separator := ""

	for _, a := range c.Args() {
		arguments += separator + a
		separator = " "
	}

	fmt.Println("Executing " + c.Command.Name + " " + arguments)
}
