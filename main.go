package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	Translate("adog")
	Translate("dog")
	Translate("bog")
	Translate("eog")
	Translate("iog")
	Translate("og")
	Translate("ulog")
	Translate("log")
}

func Translate(word string) string {
	// word is empty
	if word == "" {
		return ""
	}

	// word starts with a vowel, just add suffix
	if isVowel(word[0]) {
		return withSuffix(word)
	}

	// word starts with consonant
	// move starting consonants to the end then add suffix
	//
	start, remaining := splitStart(word)
	return withSuffix(remaining + start)
}

func isVowel(letter byte) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func withSuffix(word string) string {
	return word + "ay"
}

func splitStart(word string) (string, string) {
	consonants := ""
	for i, letter := range word {
		if isVowel(byte(letter)) {
			return consonants, word[i:]
		}
		consonants = consonants + string(letter)
	}
	return consonants, ""
}
