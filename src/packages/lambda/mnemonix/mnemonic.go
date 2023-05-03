package main

import (
	"net/http"
	"strings"

	"github.com/pnowosie/complete-mnemonic/bip39"
)

func Main(in Request) (*Response, error) {
	if in.Length == 0 {
		// params of type int cannot be passed via CLI
		in.Length = 12
	}

	mn, err := Repeat(in.Phrase, in.Length)
	if err != nil {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body:       []string{err.Error()},
		}, nil
	}

	en, _ := bip39.EntropyFromMnemonic(mn)
	mn, _ = bip39.NewMnemonic(en)

	words := strings.Fields(mn)
	return &Response{
		StatusCode: http.StatusOK,
		Body:       words,
	}, nil
}
