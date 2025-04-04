package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GetCardPrefixWithCardBrand(brandFile, issuerFile, brand, issuer string) (string, error){
	var prefixIssuer string

	bFile, err := os.Open(brandFile)
	if err != nil {
		return "", fmt.Errorf("failed to open brands file, %s", err)
	}
	defer bFile.Close()

	brandExists := false
	scanner := bufio.NewScanner(bFile)
	allBrandVariants := []string{}
	for scanner.Scan() {
		
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("wrong format of data in %s", brandFile)
		} else if strings.TrimSpace(strings.ToLower(parts[0])) == strings.TrimSpace(strings.ToLower(brand)) {
			if strings.TrimSpace(parts[1]) == "" {
				return "", fmt.Errorf("brand %s doesn't have creditcard start value", brand)
			}
			brandExists = true
			allBrandVariants = append(allBrandVariants, parts[1])
		}
	}
	if !brandExists {
		return "", fmt.Errorf("brand %s not found in %s", brand, brandFile)
	}

	iFile, err := os.Open(issuerFile)
	if err != nil {
		return "", fmt.Errorf("filed to open issuers file, %s", err)
	}
	defer iFile.Close()

	scanner = bufio.NewScanner(iFile)

	issuerExists := false
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("wrong format of data in %s", issuerFile)
		} else if strings.ToLower(strings.TrimSpace(parts[0])) == strings.ToLower(strings.TrimSpace(issuer)){
			if strings.TrimSpace(parts[1]) == "" {
				return "", fmt.Errorf("issuer %s doesn't have prefix of credit card number", issuer)
			}
			prefixIssuer = parts[1]
			issuerExists = true
			break
		}
	}
	if !issuerExists {
		return "", fmt.Errorf("issuer %s not found in %s", issuer, issuerFile)
	}
	for _, brandPrefix := range allBrandVariants {
		if len(brandPrefix) <= len(prefixIssuer) {
			if brandPrefix == prefixIssuer[:len(brandPrefix)] {
				return prefixIssuer, nil
			}
		} else {
			return "", fmt.Errorf("prefix of brand %s is greater than prefix of issuer %s", brand, issuer)
		}
	}
	
	
	return "", fmt.Errorf("prefix of brand %s does not match to prefix of issuer %s", brand, issuer)
}

func IssueCardNumber(cardPrefix string, brand string) error{

	switch strings.TrimSpace(strings.ToLower(brand)) {
	case "visa":
		randomLen := rand.Intn(2)
		if randomLen == 1 {
			randomLen = 13
		} else {
			randomLen = 16
		}
		for i := len(cardPrefix); i < randomLen - 1; i++ {
			cardPrefix += string(rand.Intn(10) + '0')
		}
		char := '0'
		for i := 0; i < 10; i++ {
			cardNum := cardPrefix + string(char)
			if valid, _ := LuhnAlgorithm(cardNum); valid {
				fmt.Println(cardNum)
				return nil
			}
			char++
		}
	case "amex": 
		cardLen := 15
		for i := len(cardPrefix); i < cardLen - 1; i++ {
			cardPrefix += string(rand.Intn(10) + '0')
		}
		char := '0'
		for i := 0; i < 10; i++ {
			cardNum := cardPrefix + string(char)
			if valid, _ := LuhnAlgorithm(cardNum); valid {
				fmt.Println(cardNum)
				return nil
			}
			char++
		}
	case "mastercard":
		cardLen := 16
		for i := len(cardPrefix); i < cardLen - 1; i++ {
			cardPrefix += string(rand.Intn(10) + '0')
		}
		char := '0'
		for i := 0; i < 10; i++ {
			cardNum := cardPrefix + string(char)
			if valid, _ := LuhnAlgorithm(cardNum); valid {
				fmt.Println(cardNum)
				return nil
			}
			char++
		}
	}
	return fmt.Errorf("underfined card brand")
}