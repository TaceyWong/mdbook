package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var InitCMD = &cli.Command{
	Name:      "init",
	Usage:     "Creates the boilerplate structure and files for a new book",
	ArgsUsage: "[dir]  \n\n\t [dir] Directory to create the book in (Defaults to the current directory when omitted)",

	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "theme",
			Usage: "Copies the default theme into your source folder",
		},
		&cli.BoolFlag{
			Name:  "force",
			Usage: "Skips confirmation prompts",
		},
		&cli.StringFlag{
			Name:  "title",
			Usage: "Set the book `TITLE`",
		},
		&cli.StringFlag{
			Name:  "ignore",
			Usage: "Creates a VCS `IGNORE` file (i.e. .gitignore) [possible values: none, git]",
		},
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("init")
		return nil
	},
}
