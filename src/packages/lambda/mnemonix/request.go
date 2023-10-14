package main

import (
	"encoding/json"
	"strconv"
)

const (
	DefaultPhraseLength = 12
)

// Request is the function's request struct
type Request struct {
	Phrase string `json:"phrase"`
	Length int    `json:"length"`
}

// Response is the function's response struct
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       ResponseBody      `json:"body,omitempty"`
}

type ResponseBody struct {
	Mnemonic string `json:"mnemonic"`
	Length   int    `json:"length"`
	Ends     string `json:"ends,omitempty"`
	Error    string `json:"error,omitempty"`
}

// UnmarshalJSON custom method of UnMarshaller interface for handling non-string types
// of the Request.
func (req *Request) UnmarshalJSON(data []byte) error {
	var objMap map[string]*json.RawMessage

	// unmarshal json to raw messages
	err := json.Unmarshal(data, &objMap)
	if err != nil {
		return err
	}

	var length int64
	if rawMsg, ok := objMap["length"]; ok {
		if err := json.Unmarshal(*rawMsg, &length); err != nil {
			// if failed, unmarshal to string
			var lengthStr string
			if err := json.Unmarshal(*rawMsg, &lengthStr); err != nil {
				return err
			}
			length, err = strconv.ParseInt(lengthStr, 10, 64)
			if err != nil {
				return err
			}
		}
	}
	if req.Length = int(length); req.Length == 0 {
		req.Length = DefaultPhraseLength
	}

	if rawMsg, ok := objMap["phrase"]; ok {
		var phrase string
		if err := json.Unmarshal(*rawMsg, &phrase); err != nil {
			return err
		}
		req.Phrase = phrase
	}

	return nil
}
