package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var WatchCMD = &cli.Command{
	Name:      "watch",
	Usage:     "Watches a book's files and rebuilds it on changes",
	ArgsUsage: "[dir]  \n\n\t [dir] Root directory for the book (Defaults to the current directory when omitted)",

	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "dest-dir",
			Aliases: []string{"d"},
			Usage:   "Output `directory` for the book\n\tRelative paths are interpreted relative to the book's root directory.",
		},
		&cli.BoolFlag{
			Name:    "open",
			Aliases: []string{"o"},
			Usage:   "Opens the compiled book in a web browser",
		},
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("watch")
		return nil
	},
}
