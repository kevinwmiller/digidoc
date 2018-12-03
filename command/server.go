package command

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/kevinwmiller/digidoc/digidoc"
	"github.com/mitchellh/cli"

	pb "github.com/kevinwmiller/digidoc/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ cli.Command = (*ServerCommand)(nil)

type ServerCommand struct {
	Listener net.Listener
	Server   *grpc.Server
	Digidoc  *digidoc.Digidoc
}

func (c *ServerCommand) RegisterCommands() map[string]cli.CommandFactory {
	serverCommand := &ServerCommand{
		Listener: nil,
	}

	serverCommands := map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return serverCommand, nil
		},
	}
	return serverCommands
}

func (c *ServerCommand) Help() string {
	helpText := `
Usage: digidoc server

  This command manages the Digidoc server.

  Start a server with the command:
    $ digidoc server
`
	return strings.TrimSpace(helpText)
}

func (c *ServerCommand) Run(args []string) int {
	fmt.Printf("%+v\n", c)
	var err error
	c.Listener, err = net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: Create credentials here. Provide to grpc server
	// creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	// c.ServerCommand.Server := grpc.NewServer(grpc.Creds(creds))
	c.Server = grpc.NewServer()
	pb.RegisterDigiDocServer(c.Server, c.Digidoc)
	reflection.Register(c.Server)
	if err := c.Server.Serve(c.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return 0
}

func (c *ServerCommand) Synopsis() string {
	return `Start a digidoc server instance`
}
