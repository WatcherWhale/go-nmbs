package commands

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/watcherwhale/go-nmbs/pkg/irail"
	"github.com/watcherwhale/go-nmbs/pkg/util"
)

func departureCommand() *cli.Command {
	return &cli.Command{
		Name:    "departures",
		Aliases: []string{"dep", "d"},
		Usage:   "Query a connection.",
		Action:  departureAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "from",
				Aliases:  []string{"f"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "to",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.BoolFlag{
				Name: "first",
			},
		},
	}
}

func departureAction(cCtx *cli.Context) error {

	connections := irail.GetConnections(cCtx.String("lang"), cCtx.String("from"), cCtx.String("to"))

	for _, conn := range connections {
		if conn.Departure.Canceled == "1" || conn.Arrival.Canceled == "1" {
			continue
		}

		departure := util.UnixToTime(conn.Departure.Time)
		delay, _ := strconv.Atoi(conn.Departure.Delay)

		fmt.Printf("%02d:%02d +%02d", departure.Hour(), departure.Minute(), delay/60)
		if cCtx.Bool("first") {
			break
		}
		fmt.Println()
	}

	return nil
}
