package main

import ("fmt"
		"errors"
		"gator/internal/config"
		"github.com/google/uuid"
		"time"
		"os"
		"gator/internal/database"
		"context")

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("Error Username Field was Blank\n")
	}
	s.pointer.CurrentUsername = cmd.Args[0]
	_, err := config.SetUser(*s.pointer)
	if err != nil {
		return fmt.Errorf("Error setting user: %v\n", err)
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister (s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Error: handlerRegister takes one argument\n")
		os.Exit(1)
	}
	newID := uuid.New()
	currentTime := time.Now()
	userParams := database.CreateUserParams{ID: newID, 
		CreatedAt: currentTime, UpdatedAt: currentTime, Name: cmd.Args[0]}
	newUser, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("Error Creating User: %v\n", err)
		os.Exit(1)
	}
	s.pointer.CurrentUsername = cmd.Args[0]
	fmt.Println(newUser)
	return nil
}