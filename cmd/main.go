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
		cardNum := config.Cfg.CardNumberToInform
		cardInfo := utils.CardInformation{}
		cardInfo.CardNum = cardNum

		if valid, _ := utils.LuhnAlgorithm(cardNum); valid {
			cardInfo.Correct = true
		}
		if !utils.ValidateForKnownCard(cardNum) {
			log.Fatal("Error: provided card is not in known card format")
		}
		if !cardInfo.Correct {
			utils.PrintInfoResult(&cardInfo)
			return
		}
		err := utils.CheckInBrands(config.Cfg.BrandsFile, config.Cfg.CardNumberToInform, &cardInfo)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		err = utils.CheckInIssuers(config.Cfg.IssuersFile, config.Cfg.CardNumberToInform, &cardInfo)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		utils.PrintInfoResult(&cardInfo)
	case "issue":
		cardPrefix, err := utils.GetCardPrefixWithCardBrand(config.Cfg.BrandsFile, config.Cfg.IssuersFile , config.Cfg.Brand, config.Cfg.Issuer)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		err = utils.IssueCardNumber(cardPrefix, config.Cfg.Brand)
		if err != nil {
			log.Fatal("Error: ", err)
		}
	default: log.Fatal("Error: command not provided")
	}
}
