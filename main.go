package main 

import ("fmt"
	"gator/internal/config"
	"os"
	"database/sql"
	"gator/internal/database"
 	_ "github.com/lib/pq")

type state struct {
	pointer *config.Config
	db *database.Queries
}

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading the Config: %v\n", err)
		return
	}
	datab, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Printf("Error opening postgres: %v\n", err)
		os.Exit(1)
	}
	newDb := database.New(datab)
	ali := state{pointer: &cfg, db: newDb}
	coms := commands{com:make(map[string]func(*state, command) error)}
	coms.register("login", handlerLogin)
	coms.register("register", handlerRegister)
	coms.register("reset", handlerReset)
	coms.register("users", handlerUsers)
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	} else {
		orders := command{name: os.Args[1], Args: os.Args[2:]}
		err = coms.run(&ali, orders)
		if err != nil {
			fmt.Printf("Error following orders: %v\n", err)
			os.Exit(1)
		}
	}
}