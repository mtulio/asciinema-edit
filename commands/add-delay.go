package commands

import (
	"gopkg.in/urfave/cli.v1"
)

var AddDelay = cli.Command{
	Name:   "add-delay",
	Usage:  "Increases the delay of a given frame",
	Action: addDelayAction,
}

func addDelayAction(c *cli.Context) (err error) {
	return
}