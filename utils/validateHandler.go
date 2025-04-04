package utils

import (
	"bufio"
	"fmt"
	"os"
)


func ValidateCardNumbers(cardNums []string) error {
	for _, number := range cardNums {
		if valid, err := LuhnAlgorithm(number); valid{
			fmt.Println("CORRECT")
		}else {
			if err != nil {
				return err
			} else {
				fmt.Println("INCORRECT")
				os.Exit(1)
			}
		}
	}
	return nil
}

func ValidateFromStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	for scanner.Scan(){ 
		line = scanner.Text()
		line = trimSpaces(line) 
		if line == "" {
			return fmt.Errorf("empty input")
		}
		ValidateCardNumbers([]string{line})
	}
	return nil
}



func LuhnAlgorithm(num string) (bool, error) {
		sum := 0
		doubleNum := 0
		if len(num) > 0 && num[0] == '0'{
			return false, fmt.Errorf("provided card number %s starts with 0", num)
		}
		for i := len(num)-1; i >= 0; i-=2 {
			if num[i] >= '0' && num[i] <= '9' {
				sum += int(num[i] - '0')
			} else {
				return false, fmt.Errorf("provided card number %s contains non-digit characters", num)
			}
		}

		for i := len(num)-2; i >= 0; i-=2 {
			if num[i] >= '0' && num[i] <= '9' {
				doubleNum = int(num[i] - '0')*2
				if doubleNum >= 10 {
					sum += doubleNum/10
					sum += doubleNum%10
				} else {
					sum += doubleNum
				}
			} else {
				return false, fmt.Errorf("provided card number %s contains non-digit characters", num)
			}
		}
		return sum % 10 == 0 && sum != 0, nil
}

func ValidateForKnownCard(cardNumber string) bool {
	if (len(cardNumber) == 13 || len(cardNumber) == 16) && cardNumber[0] == '4' {
		return true // Visa
	} else if len(cardNumber) == 16 && cardNumber[0] == '5' && cardNumber[1] >= '1' && cardNumber[1] <= '5' {
		return true // MasterCard
	} else if len(cardNumber) == 15 && cardNumber[0] == '3' && (cardNumber[1] == '4' || cardNumber[1] == '7') {
		return true // American Express
	} 
	return false
} 