package main

import (
	"fmt"
	"strings"

	"github.com/pnowosie/complete-mnemonic/bip39"
)

func Repeat(phrase string, length int) (string, error) {
	if err := hasCorrectWordsLength(length); err != nil {
		return "", err
	}
	words, err := toWordList(phrase)
	if err != nil {
		return "", err
	}

	// adjust the length to the most fitting and correct value
	if len(words) > 12 && len(words) <= 24 {
		length = len(words) + ((3 - len(words)%3) % 3)
	}

	dst := make([]string, length)
	for i := 0; i < length; i += len(words) {
		copy(dst[i:], words)
	}
	return strings.Join(dst, " "), nil
}

func hasCorrectWordsLength(length int) error {
	if !(length%3 == 0 && length >= 12 && length <= 24) {
		return fmt.Errorf("invalid length of '%d', accepted values: 12, 15, 18, 21, 24", length)
	}
	return nil
}

func toWordList(phrase string) ([]string, error) {
	phrase = strings.ReplaceAll(phrase, "_", " ")
	words := strings.Fields(phrase)
	if len(words) == 0 {
		return []string{}, fmt.Errorf("no words found in '%s'", phrase)
	}

	for i, word := range words {
		if _, ok := bip39.GetWordIndex(word); !ok {
			return []string{}, fmt.Errorf("word '%s' at position %d is not in WordList", word, i)
		}
	}
	return words, nil
}
