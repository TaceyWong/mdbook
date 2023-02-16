package cmd

import (
	"os"
	"path/filepath"

	"github.com/TaceyWong/mdbook/pkg/utils"
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
	Action: CleanAction,
}

func CleanAction(ctx *cli.Context) error {
	root, _ := os.Getwd()
	if ctx.NArg() > 0 {
		root = ctx.Args().First()
	}
	bookDir := filepath.Join(root, "book")
	utils.RemoveDir(bookDir)
	return nil
}
