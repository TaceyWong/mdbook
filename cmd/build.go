package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var BuildCMD = &cli.Command{
	Name:  "build",
	Usage: "Builds a book from its markdown files",
	Action: func(ctx *cli.Context) error {
		fmt.Println("build")
		return nil
	},
}
