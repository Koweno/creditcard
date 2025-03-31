package main

import (
	"creditcard/config"
	"log"
)

func main() {
	err := config.ParseCommand()
	if err != nil {
		log.Fatal(err)
	}
}
