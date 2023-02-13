package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var InitCMD = &cli.Command{
	Name:  "init",
	Usage: "Creates the boilerplate structure and files for a new book",
	Action: func(ctx *cli.Context) error {
		fmt.Println("init")
		return nil
	},
}
