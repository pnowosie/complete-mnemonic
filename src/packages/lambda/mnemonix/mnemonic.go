package main

import (
	"net/http"
	"strings"

	"github.com/pnowosie/complete-mnemonic/bip39"
)

func Main(in Request) (*Response, error) {
	in.AssumeDefaults()

	mn, err := Repeat(in.Phrase, in.Length)
	if err != nil {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body:       ResponseBody{Error: err.Error()},
		}, nil
	}

	en, _ := bip39.EntropyFromMnemonic(mn)
	mn, _ = bip39.NewMnemonic(en)
	ends := possibleLastWords(en)

	words := strings.Fields(mn)
	return &Response{
		StatusCode: http.StatusOK,
		Body: ResponseBody{
			Mnemonic: mn, Ends: strings.Join(ends, " "), Length: len(words)},
	}, nil
}

func possibleLastWords(entropy []byte) []string {
	words := make([]string, 0, MaxCorrectWords)
	bytes := PossibleLastBytes(len(entropy))
	mnWordsLength := len(entropy) / 4 * 3

	for _, last := range bytes {
		entropy[len(entropy)-1] = last
		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			return []string{}
		}

		mnWords := strings.Fields(mnemonic)
		words = append(words, mnWords[mnWordsLength-1])
	}
	return words
}
