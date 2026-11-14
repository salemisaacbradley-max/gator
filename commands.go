package main

import ("fmt"
		)

type Command struct {
	Args []string

}

type Commands struct {
	cmd map[string]func(*state, command) error
}

