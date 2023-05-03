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
				Body:       []string{"abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "abandon", "about"},
			},
		},
		"yellow-15-success": {
			req: &Request{
				Phrase: "yellow",
				Length: 15,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body:       []string{"yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "yellow", "year"},
			},
		},
		"angry-bird-24-success": {
			req: &Request{
				Phrase: "angry bird",
				Length: 24,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body:       []string{"angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "bird", "angry", "advance"},
			},
		},
		"three-short-12-success": {
			req: &Request{
				Phrase: "air age act",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body:       []string{"air", "age", "act", "air", "age", "act", "air", "age", "act", "air", "age", "addict"},
			},
		},
		"word out of word list": {
			req: &Request{
				Phrase: "not-here",
			},
			expectedResponse: &Response{
				StatusCode: 400,
				Body:       []string{"word 'not-here' at position 0 is not in WordList"},
			},
		},
		"incorrect length": {
			req: &Request{
				Phrase: "zero",
				Length: 13,
			},
			expectedResponse: &Response{
				StatusCode: 400,
				Body:       []string{"invalid length of '13', accepted values: 12, 15, 18, 21, 24"},
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
