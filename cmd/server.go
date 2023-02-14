package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var ServeCMD = &cli.Command{
	Name:      "serve",
	Usage:     "Serves a book at http://localhost:3000, and rebuilds it on changes",
	ArgsUsage: "[dir]  \n\n\t [dir] Root directory for the book (Defaults to the current directory when omitted)",

	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "dest-dir",
			Aliases: []string{"d"},
			Usage:   "Output `directory` for the book\n\tRelative paths are interpreted relative to the book's root directory.",
		},
		&cli.StringFlag{
			Name:    "hostname",
			Aliases: []string{"n"},
			Value:   "localhost",
			Usage:   "`HOSTNAME` to listen on for HTTP connections ",
		},
		&cli.Int64Flag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   3000,
			Usage:   "`PORT` to use for HTTP connections [default: 3000]",
		},
		&cli.BoolFlag{
			Name:    "open",
			Aliases: []string{"o"},
			Usage:   "Opens the compiled book in a web browser",
		},
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("serve")
		return nil
	},
}
