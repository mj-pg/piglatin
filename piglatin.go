package main

import "strings"

// Store stores texts and their pig latin translations.
type Store interface {
	Save(string, string) error
	Get() ([][2]string, error)
}

// Service represents the core services supported by this pig latin translator.
type Service struct {
	store Store
}

// Translate translates the text to pig latin and saves the translation.
func (s *Service) Translate(text string) (string, error) {
	translated := translate(text)
	// TODO: is this infinite?
	if err := s.store.Save(text, translated); err != nil {
		return translated, err
	}
	return translated, nil
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
