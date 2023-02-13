package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var WatchCMD = &cli.Command{
	Name:  "watch",
	Usage: "Watches a book's files and rebuilds it on changes",
	Action: func(ctx *cli.Context) error {
		fmt.Println("watch")
		return nil
	},
}
