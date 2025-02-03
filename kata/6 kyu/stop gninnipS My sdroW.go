package kata

import "strings"


func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverse(word)
		}
	}
	return strings.Join(words, " ")
}

func reverse(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
