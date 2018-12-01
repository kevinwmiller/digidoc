package main

import (
	"os"

	"github.com/kevinwmiller/digidoc/command"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}