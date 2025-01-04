package main

import (
	"fmt"

	"github.com/alisjj/gatorade/internal/config"
	"github.com/alisjj/gatorade/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cm, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("error: command %s doesn't exist", cmd.name)
	}
	return cm(s, cmd)

}
