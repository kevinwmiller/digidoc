package command

import (
	"strings"

	"github.com/kevinwmiller/digidoc/server"
)

type ServerStartCommand struct {
	*ServerCommand
}

func (c *ServerStartCommand) Help() string {
	helpText := `
  Usage: digidoc server start

  This command manages the Digidoc server.

  Start a server with the command:
    $ digidoc server start
`
	return strings.TrimSpace(helpText)
}

func (c *ServerStartCommand) Run(args []string) int {
	c.ServerCommand.s.Start(&server.ServerOptions{
		Port: 8080,
	})
	return 0
}

func (c *ServerStartCommand) Synopsis() string {
	return `Start a digidoc server instance`
}
