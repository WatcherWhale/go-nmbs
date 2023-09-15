package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/go-nmbs/pkg/irail"
)

func disruptionCommand() *cli.Command {
	return &cli.Command{
		Name:    "disruptions",
		Aliases: []string{"dis"},
		Usage:   "Show all the disruptions on the rail network.",
		Action:  disruptionAction,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "id",
				Value:       -1,
				DefaultText: "All",
			},
		},
	}
}

func disruptionAction(cCtx *cli.Context) error {
	disruptions := irail.GetDisruptions(cCtx.String("lang"))

	if id := cCtx.Int("id"); id != -1 {
		dis := disruptions[id]
		fmt.Println(dis.Title)
		fmt.Println()
		fmt.Println(dis.Description)
	} else {
		for _, dis := range disruptions {
			fmt.Printf("[%s] %s", dis.Id, dis.Title)
			fmt.Println()
		}
	}

	return nil
}
