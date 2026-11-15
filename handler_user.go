package main

import ("fmt"
		"errors"
		"gator/internal/config")

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("Error Username Field was Blank")
	}
	s.pointer.CurrentUsername = cmd.args[0]
	_, err := config.SetUser(*s.pointer)
	if err != nil {
		return fmt.Errorf("Error setting user: %v", err)
	}
	fmt.Println("User has been set")
	return nil
}