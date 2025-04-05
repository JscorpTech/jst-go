package utils

import "strings"

func StringToName(val string) string {
	words := strings.Fields(val)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + string(word[1:])
	}
	return strings.Join(words, "")
}
