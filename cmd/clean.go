package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var CleanCMD = &cli.Command{
	Name:      "clean",
	Usage:     "Deletes a built book",
	ArgsUsage: "[dir]  \n\n\t [dir] Root directory for the book (Defaults to the current directory when omitted)",

	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "dest-dir",
			Aliases: []string{"d"},
			Usage:   "Output `directory` for the book\n\tRelative paths are interpreted relative to the book's root directory.",
		},
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("clean")
		return nil
	},
}
