package command

import (
	"log"

	"github.com/mitchellh/cli"
)

type digidocCLI struct {
	commands map[string]cli.CommandFactory
}

func Run(args []string) int {

	d := digidocCLI{}
	d.registerCommands()

	cli := &cli.CLI{
		Name:     "digidoc",
		Version:  "0.0.1",
		Args:     args,
		Commands: d.commands,
	}

	exitCode, err := cli.Run()
	if err != nil {
		log.Println(err)
		return 1
	}
	return exitCode
}
