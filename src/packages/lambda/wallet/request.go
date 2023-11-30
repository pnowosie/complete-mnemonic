package main

const (
	DefaultPhraseLength = 12
	DefaultDerivation   = "m/44'/60'/0'/0/"
	DefaultAccountCount = 10
)

// Request is the function's request struct
type Request struct {
	Length        int    `json:"length,string,omitempty"`
	Count         int    `json:"count,string,omitempty"`
	Mnemonic      string `json:"mnemonic,omitempty"`
	Phrase        string `json:"phrase,omitempty"`
	Derivation    string `json:"derivation,omitempty"`
	Password      string `json:"password,omitempty"`
	RevealPrivate bool   `json:"reveal,omitempty"`
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
	Address    string `json:"address"`
	PublicKey  string `json:"publicKey,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
}

func (req *Request) AssumeDefaults() {
	if req.Length == 0 {
		req.Length = DefaultPhraseLength
	}
	if req.Count == 0 {
		req.Count = DefaultAccountCount
	}
	if req.Derivation == "" {
		req.Derivation = DefaultDerivation
	}
}
