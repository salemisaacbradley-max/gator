package main

import ("fmt"
		"errors"
		"gator/internal/config"
		"github.com/google/uuid"
		"time"
		"gator/internal/database"
		"context")

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("Error Username Field was Blank\n")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error getting user data: %v\n", err)
	}
	s.pointer.CurrentUsername = cmd.Args[0]
	_, err = config.SetUser(*s.pointer)
	if err != nil {
		return fmt.Errorf("Error setting user: %v\n", err)
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister (s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Error: handlerRegister takes one argument\n")
	}
	newID := uuid.New()
	currentTime := time.Now()
	userParams := database.CreateUserParams{ID: newID, 
		CreatedAt: currentTime, UpdatedAt: currentTime, Name: cmd.Args[0]}
	//fmt.Println("LOL. LMAO even")
	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("Error Creating User: %v\n", err)
	}
	s.pointer.CurrentUsername = cmd.Args[0]
	_, err = config.SetUser(*s.pointer)
	if err != nil {
    	return fmt.Errorf("Error setting user: %v\n", err)
	}
	fmt.Println(userParams)
	return nil
}

//Only for use in debugging
func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Error resetting: %v", err)
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error retrieving users: %v", err)
	}
	for _, user := range users {
		if user.Name == s.pointer.CurrentUsername {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	data, err := MethodGet("https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Error Getting for Aggregator: %w", err)
	}
	feed, err := cmd.FetchFeed(data)
	if err != nil {
		return fmt.Errorf("Error Fetching Feed: %w", err)
	}
	fmt.Printf()
}