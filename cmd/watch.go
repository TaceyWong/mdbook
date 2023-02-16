package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fsnotify/fsnotify"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli/v2"
)

var WatchCMD = &cli.Command{
	Name:      "watch",
	Usage:     "Watches a book's files and rebuilds it on changes",
	ArgsUsage: "[dir]  \n\n\t [dir] Root directory for the book (Defaults to the current directory when omitted)",

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
	Action: WatchAction,
}

func WatchAction(ctx *cli.Context) error {
	root, _ := os.Getwd()
	watchPath := filepath.Join(root, "book")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event, event.Name)
				// if event.Has(fsnotify.Write) {
				// log.Println("modified file:", event.Name)
				// }
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	// Add path to watch.
	err = watcher.Add(watchPath)
	if err != nil {
		return err
	}

	if ctx.Bool("open") { // 打开浏览器
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
		s.Color("bgBlack", "bold", "fgRed")
		s.Suffix = " Openning Browser,wait a momment……"
		s.Start()                   // Start the spinner
		time.Sleep(4 * time.Second) // Run for some time to simulate work
		browseres := []string{"google-chrome", "firefox"}
		browserS := "default"
		for index, browser := range browseres {
			if err := open.RunWith("https://baidu.com/", browser); err == nil {
				browserS = browser
				break
			}
			if index >= len(browseres)-1 {
				if err := open.Run("https://baidu.com/"); err != nil {
					return err
				}
			}
		}
		s.Stop()
		fmt.Printf("open in %s browser\n", browserS)
	}
	// Block main goroutine forever.
	<-make(chan struct{})
	return nil
}
