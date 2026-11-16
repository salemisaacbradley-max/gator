package main 

import ("fmt"
	"gator/internal/config"
	"os"
	"database/sql")
import _ "github.com/lib/pq"

type state struct {
	pointer *config.Config
	db *database.Queries
}

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading the Config: %v", err)
		return
	}
	datab, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("Error opening postgres: %v", err)
		os.Exit(1)
	}
	newDb := database.New(datab)
	ali := state{pointer: &cfg, db: newDb}
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