package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestUnMarshall(t *testing.T) {
	tests := map[string]struct {
		json            string
		success         bool
		expectedRequest *Request
	}{
		"only one word": {
			json:    `{"phrase": "abandon"}`,
			success: true,
			expectedRequest: &Request{
				Phrase: "abandon",
				Length: 12,
			},
		},
		"word with a length": {
			json:    `{"length": "15", "phrase": "yellow"}`,
			success: true,
			expectedRequest: &Request{
				Phrase: "yellow",
				Length: 15,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var req Request
			err := json.Unmarshal([]byte(test.json), &req)
			if test.success && err != nil {
				t.Errorf("expected success, got error: %v", err)
			}
			if !test.success && err == nil {
				t.Errorf("expected error, got success")
			}
			if !test.success {
				return
			}

			req.AssumeDefaults()
			assert.Equal(t, test.expectedRequest, &req)
		})

	}
}
