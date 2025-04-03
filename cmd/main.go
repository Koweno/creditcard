package main

import (
	"creditcard/config"
	"creditcard/utils"
	"log"
	"strings"
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
		prompt := strings.TrimSpace(config.Cfg.CardPromptToGenerate)
		err := utils.GenerateCardNumbers(prompt, config.Cfg.Pick)
		if err != nil {
			log.Fatal("Error: ", err)
		} 
	case "information":
	case "issue":
	default: log.Fatal("Error: command not provided")
	}
}
