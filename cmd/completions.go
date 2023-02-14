package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var CompletionsCMD = &cli.Command{
	Name:      "completions",
	Usage:     "Generate shell completions for your shell to stdout",
	ArgsUsage: "<SHELL> \n\n\t <SHELL> the shell to generate completions for [possible values: bash, elvish, fish, powershell, zsh]",

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
		fmt.Println("completions")
		return nil
	},
}
