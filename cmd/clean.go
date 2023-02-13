package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var CleanCMD = &cli.Command{
	Name:  "clean",
	Usage: "Deletes a built book",
	Action: func(ctx *cli.Context) error {
		fmt.Println("clean")
		return nil
	},
}
