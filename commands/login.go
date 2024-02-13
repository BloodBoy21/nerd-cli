package commands

import "fmt"


func (c *CommandService) Login() {
	fmt.Printf("Running %s command\n", c.Command.Flag)
}