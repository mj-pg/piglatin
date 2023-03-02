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

		// ignore punctuations at the ends because they may be valid
		// e.g.  (something less important) ,comma question?
		//
		left, mid, right := trimPuncts(word)
		pl := pigLatinize(mid)
		res[i] = left + pl + right
	}
	return strings.Join(res, " ")
}

// pigLatinize returns the pig latin version of a word.
func pigLatinize(word string) string {
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
	// split at the 1st non-consonant
	i := strings.IndexFunc(word, isNonConsonant)
	// no vowels
	if i < 0 {
		return word, ""
	}
	return word[:i], word[i:]
}

// isVowel checks if a character is a vowel.
func isVowel(ch rune) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// isNonConsonant checks if a character is not a consonant.
func isNonConsonant(ch rune) bool {
	isConsonant := unicode.IsLetter(ch) && !isVowel(ch)
	return !isConsonant
}

// isValid checks if a word is probably valid.
// A valid word has letters, single quote and hyphen only.
func isValid(word string) bool {
	if word == "" {
		return false
	}
	for _, ch := range word {
		if !unicode.IsLetter(ch) &&
			byte(ch) != '\'' &&
			byte(ch) != '-' {
			return false
		}
	}
	return true
}

// trimPuncs returns the trimmed word and the subsequents punctuations at the start(leftmost) and end(rightmost) of the word.
func trimPuncts(word string) (string, string, string) {
	// remove punctuations at the start
	withoutLeft := strings.TrimLeftFunc(word, unicode.IsPunct)

	// word is all punctuation
	if withoutLeft == "" {
		return word, "", ""
	}

	// remove punctuations at the end
	mid := strings.TrimRightFunc(withoutLeft, unicode.IsPunct)

	// get starting punctuations(at the left)
	//
	iLeftEnd := strings.Index(word, withoutLeft)
	left := word[:iLeftEnd]

	// zero ending punctuations
	if len(left)+len(mid) == len(word) {
		return left, mid, ""
	}

	// get start of the ending punctuations(at the right)
	//
	iRightStart := iLeftEnd + len(mid)
	right := word[iRightStart:]

	return left, mid, right
}
