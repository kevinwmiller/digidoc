package command

import (
	"github.com/mitchellh/cli"
)

func (d *digidocCLI) registerCommands() {
	d.commands = map[string]cli.CommandFactory{}

	// Register commands here
	d.registerCommand((&ServerCommand{}).RegisterCommands())
}

func (d *digidocCLI) registerCommand(commands map[string]cli.CommandFactory) {
	for k, v := range commands {
		d.commands[k] = v
	}
}
