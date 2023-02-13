package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var ServeCMD = &cli.Command{
	Name:  "serve",
	Usage: "Serves a book at http://localhost:3000, and rebuilds it on changes",
	Action: func(ctx *cli.Context) error {
		fmt.Println("serve")
		return nil
	},
}
