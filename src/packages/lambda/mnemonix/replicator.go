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

	// TODO: more is comming soon
	return strings.TrimSpace(strings.Repeat(words[0]+" ", length)), nil
}

func hasCorrectWordsLength(length int) error {
	if !(length%3 == 0 && length >= 12 && length <= 24) {
		return fmt.Errorf("invalid length of '%d', accepted values: 12, 15, 18, 21, 24", length)
	}
	return nil
}

func toWordList(phrase string) ([]string, error) {
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

