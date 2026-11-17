package main

import ("fmt"
		)

type command struct {
	name string
	Args []string

}

type commands struct {
	com map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	err := c.com[cmd.name](s, cmd)
	if err != nil {
		return fmt.Errorf("Error trying to run %v command: %w\n", cmd.name, err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.com[name] = f
}