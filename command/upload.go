package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/kevinwmiller/digidoc/client"
)

var _ cli.Command = (*UploadCommand)(nil)

type UploadCommand struct {
	ServerAddr string,
	Filename string,
}

func (c *UploadCommand) RegisterCommands() map[string]cli.CommandFactory {
	UploadCommand := &UploadCommand{
		filename: "",
	}

	UploadCommands := map[string]cli.CommandFactory{
		"upload": func() (cli.Command, error) {
			return UploadCommand, nil
		},
	}
	return UploadCommands
}

func (c *UploadCommand) Help() string {
	helpText := `
Usage: digidoc upload -file <filename>

  This command uploads a file to the Digidoc server
`
	return strings.TrimSpace(helpText)
}

func (c *UploadCommand) Run(args []string) int {
	return client.Upload(c.filename)
}

func (c *UploadCommand) Synopsis() string {
	return `Upload a file to the Digidoc server`
}
