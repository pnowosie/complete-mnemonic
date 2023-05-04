package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhraseRepetitionCompletion(t *testing.T) {
	tests := map[string]struct {
		req              *Request
		expectedResponse *Response
	}{
		"abandon-12 sucess": {
			req: &Request{
				Phrase: "abandon",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
					Length:   12,
				},
			},
		},
		"yellow-15-success": {
			req: &Request{
				Phrase: "yellow",
				Length: 15,
			},
			expectedResponse: &Response{
				StatusCode: 200,

				Body: ResponseBody{
					Mnemonic: "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow year",
					Length:   15,
				},
			},
		},
		"angry-bird-24-success": {
			req: &Request{
				Phrase: "angry bird",
				Length: 24,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "angry bird angry bird angry bird angry bird angry bird angry bird angry bird angry bird angry bird angry bird angry bird angry advance",
					Length:   24,
				},
			},
		},
		"non-whitespace-word-separator-success": {
			req: &Request{
				Phrase: "angry_bird",
				Length: 12,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "angry bird angry bird angry bird angry bird angry bird angry birth",
					Length:   12,
				},
			},
		},
		"three-short-12-success": {
			req: &Request{
				Phrase: "air age act",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "air age act air age act air age act air age addict",
					Length:   12,
				},
			},
		},
		"last word affects checksum": {
			req: &Request{
				Phrase: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon zoo",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon wrap",
					Length:   12,
				},
			},
		},
		"overwrite-default-length-when-more-words-15": {
			req: &Request{
				Phrase: "air age act air age act air age act air age act fox",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "air age act air age act air age act air age act fox air airport",
					Length:   15,
				},
			},
		},
		"overwrite-default-length-when-more-words-18": {
			req: &Request{
				Phrase: "air age act air age act air age act air age act blue fox blue fox green zebra",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "air age act air age act air age act air age act blue fox blue fox green window",
					Length:   18,
				},
			},
		},
		"word out of word list": {
			req: &Request{
				Phrase: "not-here",
			},
			expectedResponse: &Response{
				StatusCode: 400,
				Body: ResponseBody{
					Error: "word 'not-here' at position 0 is not in WordList",
				},
			},
		},
		"incorrect length": {
			req: &Request{
				Phrase: "zero",
				Length: 13,
			},
			expectedResponse: &Response{
				StatusCode: 400,
				Body: ResponseBody{
					Error: "invalid length of '13', accepted values: 12, 15, 18, 21, 24",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resp, err := Main(*test.req)
			if err != nil {
				t.Errorf("expected success, got error: %v", err)
			}
			assert.Equal(t, test.expectedResponse, resp)
		})
	}
}
