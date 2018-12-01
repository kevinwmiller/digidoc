package command

import (
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func registerCommands() {
	Commands = map[string]cli.CommandFactory{}
	register((ServerCommand{}).RegisterCommands())
}

func register(commands map[string]cli.CommandFactory) {
	for k, v := range commands {
	    Commands[k] = v
	}
}