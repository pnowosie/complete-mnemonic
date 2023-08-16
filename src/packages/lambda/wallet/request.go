package main

import (
	"encoding/json"
	"strconv"
)

const (
	DefaultPhraseLength = 12
	DefaultDerivation   = "m/44'/60'/0'/0/"
	DefaultAccountCount = 10
)

// Request is the function's request struct
type Request struct {
	Length     int    `json:"length,omitempty"`
	Count      int    `json:"count,omitempty"`
	Mnemonic   string `json:"mnemonic,omitempty"`
	Phrase     string `json:"phrase,omitempty"`
	Derivation string `json:"derivation,omitempty"`
}

// Response is the function's response struct
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       ResponseBody      `json:"body,omitempty"`
}

type ResponseBody struct {
	Wallet   WalletBody    `json:"wallet"`
	Accounts []AccountBody `json:"accounts"`
	Error    string        `json:"error,omitempty"`
}

type WalletBody struct {
	Mnemonic   string `json:"mnemonic"`
	Length     int    `json:"length"`
	Derivation string `json:"derivation"`
}

type AccountBody struct {
	Address string `json:"address"`
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

	// Length
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

	// Count
	var count int64
	if rawMsg, ok := objMap["count"]; ok {
		if err := json.Unmarshal(*rawMsg, &count); err != nil {
			// if failed, unmarshal to string
			var countStr string
			if err := json.Unmarshal(*rawMsg, &countStr); err != nil {
				return err
			}
			count, err = strconv.ParseInt(countStr, 10, 64)
			if err != nil {
				return err
			}
		}
	}
	if req.Count = int(count); req.Count == 0 {
		req.Count = DefaultAccountCount
	}

	// Derivation
	if rawMsg, ok := objMap["derivation"]; ok {
		var derivation string
		if err := json.Unmarshal(*rawMsg, &derivation); err != nil {
			return err
		}
		req.Derivation = derivation
	}
	if req.Derivation == "" {
		req.Derivation = DefaultDerivation
	}

	// Phrase
	if rawMsg, ok := objMap["phrase"]; ok {
		var phrase string
		if err := json.Unmarshal(*rawMsg, &phrase); err != nil {
			return err
		}
		req.Phrase = phrase
	}

	// Mnemonic
	if rawMsg, ok := objMap["mnemonic"]; ok {
		var mnemonic string
		if err := json.Unmarshal(*rawMsg, &mnemonic); err != nil {
			return err
		}
		req.Mnemonic = mnemonic
	}

	return nil
}
