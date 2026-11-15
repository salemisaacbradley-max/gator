package main 

import ("fmt"
	"gator/internal/config"
	"os")

type state struct {
	pointer *config.Config
}

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading the Config: %v", err)
		return
	}
	ali := state{pointer: &cfg}
	coms := commands{com:make(map[string]func(*state, command) error)}
	coms.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		fmt.Println("Username Required")
		os.Exit(1)
	} else {
		orders := command{name: os.Args[1], args: os.Args[2:]}
		err = coms.run(&ali, orders)
		if err != nil {
			fmt.Printf("Error following orders: %v", err)
			os.Exit(1)
		}
	}
}