package main

import (
	"net/http"
	"strings"

	"github.com/pnowosie/complete-mnemonic/bip39"
)

// Request is the function's request struct
type Request struct {
	Word   string `json:"word"`
	Length int    `json:"length"`
}

// Response is the function's response struct
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       []string          `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	if in.Length == 0 {
		// params of type int cannot be passed via CLI
		in.Length = 12
	}

	correctLength := in.Length%3 == 0 && in.Length >= 12 && in.Length <= 24
	if !correctLength {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body:       []string{"Invalid length, accepted values: 12, 15, 18, 21, 24"},
		}, nil
	}

	if _, ok := bip39.GetWordIndex(in.Word); !ok {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body:       []string{"Invalid word"},
		}, nil
	}

	mn := strings.TrimSpace(strings.Repeat(in.Word+" ", in.Length))
	en, err := bip39.EntropyFromMnemonic(mn)
	if err != nil {
		mn, _ = bip39.NewMnemonic(en)
	}
	words := strings.Fields(mn)
	lastWordIdx := len(words) - 1

	return &Response{
		StatusCode: http.StatusOK,
		Body:       []string{words[0], words[lastWordIdx]},
	}, nil
}
