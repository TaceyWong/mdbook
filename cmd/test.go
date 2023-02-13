package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var TestCMD = &cli.Command{
	Name:  "test",
	Usage: "Tests that a book's Rust code samples compile",
	Action: func(ctx *cli.Context) error {
		fmt.Println("test")
		return nil
	},
}
