package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var CompletionsCMD = &cli.Command{
	Name:  "completions",
	Usage: "Generate shell completions for your shell to stdout",
	Action: func(ctx *cli.Context) error {
		fmt.Println("completions")
		return nil
	},
}
