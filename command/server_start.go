package command

import (
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/kevinwmiller/digidoc/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c.ServerCommand)
	var err error
	c.ServerCommand.Listener, err = net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: Create credentials here. Provide to grpc server
	// creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	// c.ServerCommand.Server := grpc.NewServer(grpc.Creds(creds))
	c.ServerCommand.Server = grpc.NewServer()
	pb.RegisterDigiDocServer(c.ServerCommand.Server, c.ServerCommand.Digidoc)
	reflection.Register(c.Server)
	if err := c.ServerCommand.Server.Serve(c.ServerCommand.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return 0
}

func (c *ServerStartCommand) Synopsis() string {
	return `Start a digidoc server instance`
}
