package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
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
	Action: InitAction,
}

func InitAction(ctx *cli.Context) error {
	fmt.Println("init")
	vcs := false
	prompt := &survey.Confirm{
		Message: "Do you want a .gitignore to be created?",
	}
	survey.AskOne(prompt, &vcs)
	vcs_file := ".gitignore"
	if vcs {
		prompt_vcs_file := &survey.Select{
			Message: "Choose a VCS ignore file :",
			Options: []string{"git", "svn", "none"},
		}
		survey.AskOne(prompt_vcs_file, &vcs_file)
	}
	title := ""
	prompt_title := &survey.Input{
		Message: "What title would you like to give the book?",
	}
	survey.AskOne(prompt_title, &title)
	return nil
}
