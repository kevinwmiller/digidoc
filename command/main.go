package command

import (
	"log"

	"github.com/mitchellh/cli"
)

func Run(args []string) int {

	registerCommands()

	cli := &cli.CLI{
		Name:     "digidoc",
		Version:  "0.0.1",
		Args:     args,
		Commands: Commands,
	}

	exitCode, err := cli.Run()
	if err != nil {
		log.Println(err)
		return 1
	}
	return exitCode
}