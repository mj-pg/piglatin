package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

// Translate turns each word in the text to its pig latin version.
func Translate(text string) string {
	ww := strings.Fields(text)
	res := make([]string, len(ww))
	for i, word := range ww {
		pl := pigLatinize(word)
		res[i] = pl
	}
	return strings.Join(res, " ")
}

// pigLatinize returns the pig latin version of a word.
func pigLatinize(word string) string {
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

// isVowel checks if a character is a vowel.
func isVowel(letter byte) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// withSuffix appends the suffix 'ay' to the word.
func withSuffix(word string) string {
	return word + "ay"
}

// splitStart splits the starting consonant or cluster from the rest of the word.
func splitStart(word string) (string, string) {
	var consonants string
	for i, letter := range word {
		if isVowel(byte(letter)) {
			return consonants, word[i:]
		}
		consonants = consonants + string(letter)
	}
	return consonants, ""
}
