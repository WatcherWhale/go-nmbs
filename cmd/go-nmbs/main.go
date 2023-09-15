package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/go-nmbs/pkg/commands"
)

func main() {
	cmds := commands.LoadComands()
	app := &cli.App{
		Name:  "go-nmbs",
		Usage: "The cli for when you are on the go with the NMBS.",
		Authors: []*cli.Author{
			{
				Name:  "Mathias Maes",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "en",
				Usage:   "en|nl|fr|de",
			},
		},
		Commands:             cmds,
		EnableBashCompletion: true,
		BashComplete: func(cCtx *cli.Context) {
			// This will complete if no args are passed
			if cCtx.NArg() > 0 {
				return
			}
			for _, cmd := range cmds {
				fmt.Println(cmd.Name)
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
