package main 

import ("fmt"
	"gator/internal/config")

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading the Config: %v", err)
		return
	}
	_, err = config.SetUser(cfg)
	if err != nil {
		fmt.Printf("Error writing to the Config: %v", err)
		return
	}
	data, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading the second time: %v", err)
		return
	}
	fmt.Println(data)
}