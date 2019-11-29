package main

import (
	"fmt"
	"log"
	"os"

	internal "github.com/papey/yckms/internal/app"
	"github.com/urfave/cli"
)

func main() {

	// Declare app
	app := cli.NewApp()
	// Basic config
	app.Name = "YCKMS"
	app.Usage = "Sync playlists from frech metal podcasts shows to Spotify"
	app.Version = "0.1.0"

	// Flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "to, t",
			Usage: "Sync podcast to date YYYY-MM-DD",
		},
		cli.StringFlag{
			Name:  "from, f",
			Usage: "Sync podcast from date YYYY-MM-DD",
		},
		cli.BoolFlag{
			Name:  "date, d",
			Usage: "Enable dates flags",
		},
		cli.BoolFlag{
			Name:  "last, l",
			Usage: "Sync last podcast show",
		},
	}

	// Action
	app.Action = func(c *cli.Context) error {
		// Args check
		if c.NArg() != 1 {
			log.Fatal("Error: RSS feed URL argument is missing")
		}

		if c.Bool("date") && c.Bool("last") {
			fmt.Println("Warning: from, to and last flags set, last will be used")
		}

		// Flags check
		if c.Bool("last") {
			return internal.Sync(c.Args().First(), true, "", "")
		}

		// Check dates
		if c.Bool("date") {
			if c.String("from") != "" && c.String("to") != "" {
				return internal.Sync(c.Args().First(), false, c.String("from"), c.String("to"))
			}
			log.Fatal("Error: one of the dates is missing")
		}

		if !c.Bool("last") && !c.Bool("date") {
			return internal.Sync(c.Args().First(), false, "", "")
		}

		return nil

	}

	// Run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
