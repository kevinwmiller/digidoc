package command

import (
	"fmt"
)

type ServerStartCommand struct {
	*ServerCommand
}

func (c *ServerStartCommand) Run (args []string) int {
	fmt.Println("Server start command Run")
	return 0
}