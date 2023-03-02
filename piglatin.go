package main

import (
	"strings"
	"unicode"
)

const (
	VSUFF = "way"
	CSUFF = "ay"
)

// Store stores texts and their pig latin translations.
type Store interface {
	Save(string, string) error
	Get() ([][2]string, error)
	//GetByText(string) (string, error)
}

// Service represents the core services supported by this pig latin translator.
type Service struct {
	store Store
}

// Translate translates the text to pig latin and saves the translation.
func (s *Service) Translate(text string) (string, error) {

	/* // already exists
	if translated, err := s.store.GetByText(text); err == nil {
		log.Println("Got translation from storage")
		return translated, nil
	}*/

	translated := translate(text)
	// should saving try forever?
	// no, it is not that important compared to actually giving the translation
	return translated, s.store.Save(text, translated)
}

// List returns all the text and their pig latin translations.
func (s *Service) List() ([][2]string, error) {
	return s.store.Get()
}

// translate turns each word in the text to its pig latin version.
func translate(text string) string {
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
	if !isValid(word) {
		return word
	}

	// the rule is
	// if word starts with vowel then just add suffix 'way'
	// else move starting constant/cluster first before adding suffix 'ay'
	//

	// word starts with vowel
	first := rune(word[0])
	if isVowel(first) {
		return word + VSUFF
	}

	// word starts with consonant
	start, remaining := splitStart(word)
	return remaining + start + CSUFF
}

// splitStart splits the starting consonant/s from the rest of the word.
func splitStart(word string) (string, string) {
	// split at the 1st vowel
	i := strings.IndexFunc(word, isVowel)
	// no vowels
	if i < 0 {
		return word, ""
	}
	return word[:i], word[i:]
}

// isVowel checks if a character is a vowel.
func isVowel(letter rune) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// isValid checks if a word is probably valid. It allows hyphen and single quote.
func isValid(word string) bool {
	for _, ch := range word {
		if !unicode.IsLetter(ch) &&
			byte(ch) != '\'' &&
			byte(ch) != '-' {
			return false
		}
	}
	return true
}
