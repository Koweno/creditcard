package utils

import (
	"fmt"
	"math/rand"
)

func GenerateCardNumbers(cardPrompt string, pick bool) error {
	astNum, err := validateCardPrompt(cardPrompt)
	if err != nil {
		return err
	}
	cardNumWithoutAst := cardPrompt[:len(cardPrompt)-astNum]
	
	digits := make([]byte, astNum)
	if !pick {
		for i := range digits {
			digits[i] = '0'
		}
	}else {
		for i := range digits {
			digits[i] = byte(rand.Intn(10)) + '0'
		}
	}

	var cardNum string

	for {
		cardNum = cardNumWithoutAst + string(digits)
		valid, err := LuhnAlgorithm(cardNum)
		if err != nil {
			return err
		}
		if valid && !pick{
			fmt.Println(cardNum)
		} else if valid && pick{
			fmt.Println(cardNum)
			break
		}

		pos := astNum - 1
		for pos >= 0{
			digits[pos]++
			if digits[pos] <= '9' {
				break
			}
			digits[pos] = '0'
			pos--
		}
		if pos < 0 {
			break
		}
	}
	
	return nil
}

func validateCardPrompt(s string) (int, error) {
	if len(s) < 13 {
		return -1, fmt.Errorf("too short card number prompt")
	}
	cntAst := 0
	cntDigits := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9'{
			cntDigits++
		} else if s[i] == '*' {
			cntAst++
			for j := i+1; j < len(s); j++ {
				if s[j] != '*' {
					return -1, fmt.Errorf("asterisks must be written together")
				} else {
					cntAst++
				}
			}
			break
		} else {
			return -1, fmt.Errorf("prompt %s includes non-digit characters", s)
		}
	}
	if cntAst <= 0 || cntAst > 4 {
		return -1, fmt.Errorf("amount of asterisks must be between 1 and 4")
	} 
	return cntAst, nil
}