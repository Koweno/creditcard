package utils

func trimSpaces(str string) string {
	newStr := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		newStr += string(str[i])
	}
	return newStr
}
