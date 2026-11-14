package main

import ("fmt"
		"errors")

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("Error Username Field was Blank")
	}
	s.pointer.CurrentUsername = cmd.Args[0]
	fmt.Println("User has been set")
	return nil
}