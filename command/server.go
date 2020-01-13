package command

import (
	"strings"

	"github.com/kevinwmiller/digidoc/server"

	"github.com/mitchellh/cli"
)

var _ cli.Command = (*ServerCommand)(nil)

type ServerCommand struct {
	s *server.Server
}

func (c *ServerCommand) RegisterCommands() map[string]cli.CommandFactory {
	serverCommand := &ServerCommand{
		s: server.New(),
	}

	serverCommands := map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return serverCommand, nil
		},
		"server start": func() (cli.Command, error) {
			return &ServerStartCommand{
				serverCommand,
			}, nil
		},
	}
	return serverCommands
}

func (c *ServerCommand) Help() string {
	helpText := `
  Usage: digidoc server <command>

  This command manages the Digidoc server.

  Start a server with the command:
    $ digidoc server start
`
	return strings.TrimSpace(helpText)
}

func (c *ServerCommand) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *ServerCommand) Synopsis() string {
	return `Manage a Digidoc server instance`
}
