package main

const (
	DefaultPhraseLength = 12
)

// Request is the function's request struct
type Request struct {
	Phrase string `json:"phrase"`
	Length int    `json:"length,string,omitempty"`
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

func (req *Request) AssumeDefaults() {
	if req.Length == 0 {
		req.Length = DefaultPhraseLength
	}
}
