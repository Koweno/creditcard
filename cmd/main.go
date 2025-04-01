package main

import (
	"creditcard/config"
	"fmt"
	"log"
)

func main() {
	err := config.ParseCommand()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(config.Cfg) 
}
