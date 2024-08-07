package lib

import "strings"

func ProperTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to de "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
		}
	}
	return strings.Join(words, " ")
}
