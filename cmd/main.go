package main

import (
	"creditcard/config"
	"creditcard/utils"
	"log"
)

func main() {
	err := config.ParseCommand()
	if err != nil {
		log.Fatal(err)
	}
	
	switch config.Cfg.Command {
	case "validate":
		if config.Cfg.Stdin {
			err := utils.ValidateFromStdin()
			if err != nil {
				log.Fatal("Error: ", err)
			}
		} else {
			err := utils.ValidateCardNumbers(config.Cfg.CardNumbersToValidate)
			if err != nil {
				log.Fatal("Error: ", err)
			}
		}
		
	case "generate":
		if config.Cfg.Pick {

		}else {
			
		}
	case "information":
	case "issue":
	default: log.Fatal("Error: command not provided")
	}
}
