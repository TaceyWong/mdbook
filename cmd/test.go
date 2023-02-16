package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var TestCMD = &cli.Command{
	Name:      "test",
	Usage:     "Tests that a book's code samples compile",
	ArgsUsage: "[dir]  \n\n\t [dir] Root directory for the book (Defaults to the current directory when omitted)",

	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "dest-dir",
			Aliases: []string{"d"},
			Usage:   "Output `directory` for the book\n\tRelative paths are interpreted relative to the book's root directory.",
		},
		&cli.StringFlag{
			Name:    "chapter",
			Aliases: []string{"c"},
			Usage:   "",
		},
		&cli.StringFlag{
			Name:    "library-path",
			Aliases: []string{"L"},
			Usage:   "A comma-separated list of `directories` to add to the crate search path when building tests",
		},
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("test")
		return nil
	},
}
