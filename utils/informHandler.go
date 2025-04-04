package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type CardInformation struct {
	Correct bool
	CardNum string
	brand string
	issuer string
}

func PrintInfoResult(cardInfo *CardInformation) {
	if !cardInfo.Correct {
		fmt.Println(cardInfo.CardNum)
		fmt.Println("Correct: no")
		fmt.Println("Card Brand: -")
		fmt.Println("Card Issuer: -")
	} else {
		fmt.Println(cardInfo.CardNum)
		fmt.Println("Correct: yes")
		if strings.TrimSpace(cardInfo.brand) != "" {
			fmt.Println("Card Brand:", cardInfo.brand)
		} else {
			fmt.Println("Card Brand: -")
		}
		if strings.TrimSpace(cardInfo.issuer) != "" {
			fmt.Println("Card Issuer:", cardInfo.issuer)
		} else {
			fmt.Println("Card Issuer: -")
		}
	}
}

func CheckInBrands(brandsFile, CardNum string, cardInfo *CardInformation) error {
	cardInfo.CardNum = CardNum
	bFile, err := os.Open(brandsFile)
	if err != nil {
		return fmt.Errorf("failed to open brands file, %s", err)
	}
	defer bFile.Close()

	scanner := bufio.NewScanner(bFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) > 2 {
			return fmt.Errorf("wrong data format in brands file")
		} else if len(parts) < 2{
			continue
		} 
		if len(parts[1]) > len(CardNum)  {
			return fmt.Errorf("length of brand prefix (%s) is longer than card number (%s)", parts[1], CardNum)
		}
		if len(parts[1]) != 0 && parts[1] == CardNum[:len(parts[1])] {
			cardInfo.brand = parts[0]
			break
		}
	}
	return nil
}

func CheckInIssuers(issuersFile, CardNum string, cardInfo *CardInformation) error {
	iFile, err := os.Open(issuersFile)
	if err != nil {
		return fmt.Errorf("failed to open issuers file, %s", err)
	}
	defer iFile.Close()

	scanner := bufio.NewScanner(iFile) 

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) > 2 {
			return fmt.Errorf("wrong data format in issuers file")
		} else if len(parts) < 2{
			continue
		} 
		if len(parts[1]) > len(CardNum)  {
			return fmt.Errorf("length of issuer prefix (%s) is longer than card number (%s)", parts[1], CardNum)
		}
		if len(parts[1]) != 0 && parts[1] == CardNum[:len(parts[1])] {
			cardInfo.issuer = parts[0]
			break
		}
	}
	return nil
}