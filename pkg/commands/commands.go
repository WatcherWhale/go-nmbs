package commands

import (
	"github.com/urfave/cli/v2"
)

func LoadComands() []*cli.Command {
	return []*cli.Command{
		disruptionCommand(),
		departureCommand(),
	}
}
