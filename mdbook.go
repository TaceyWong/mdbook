package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TaceyWong/mdbook/cmd"
	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "mdbook",
		Usage:     "Creates a book from markdown files",
		Authors:   []*cli.Author{{Name: "Tacey Wong", Email: "xinyong.wang@qq.com"}},
		Copyright: "© 2023 TaceyWong under GPL3",
		Action: func(*cli.Context) error {
			s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
			s.Color("bgBlack", "bold", "fgRed")
			s.Suffix = "Building"
			s.Start()                   // Start the spinner
			time.Sleep(4 * time.Second) // Run for some time to simulate work
			s.Stop()
			fmt.Println("Completed")
			return nil
		},
	}
	app.EnableBashCompletion = true
	app.HideHelpCommand = false
	app.Version = "dev-0.0.1"
	app.HideVersion = false
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		cmd.InitCMD,
		cmd.BuildCMD,
		cmd.TestCMD,
		cmd.CleanCMD,
		cmd.CompletionsCMD,
		cmd.WatchCMD,
		cmd.ServeCMD,
	}
	// sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
