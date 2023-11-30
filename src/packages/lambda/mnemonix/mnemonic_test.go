package main

import (
	"github.com/pnowosie/complete-mnemonic/bip39"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
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
					Ends:     "about attract burger cool disagree exhaust furnace huge moment own question sand solid tent urge wrap",
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
					Ends:     "account autumn buffalo clown dirt excite found hurdle minimum ordinary protect rubber snow switch urge year",
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
					Ends:     "brisk castle faint guilt life pluck task update",
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
					Ends:     "absent audit burden company distance exist garbage husband modify panel quiz safe sort tattoo urban wrist",
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
					Ends:     "ability asthma burden convince dinosaur evoke game humor mixed paper quiz save soon thank unlock wrong",
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
					Ends:     "about attract burger cool disagree exhaust furnace huge moment own question sand solid tent urge wrap",
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
					Ends:     "acid auto bundle concert discover era garage hover mobile owner quality rural snap tattoo urge wrist",
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
					Ends:     "among asset brand club deliver equal flight habit mixture panther praise roof sorry team tuna winner",
					Length:   18,
				},
			},
		},
		"another-three-12-success": {
			req: &Request{
				Phrase: "quick brown fox",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "quick brown fox quick brown fox quick brown fox quick brown fox",
					Ends:     "accuse attack cable confirm discover excess funny hurry moment oyster question safe solution tent urge wreck",
					Length:   12,
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

func TestEntropyInformations(t *testing.T) {
	const (
		wordEntropyBitLength = 11
	)

	tests := map[string]struct {
		wordsLength                int
		entropyByteLength          int
		entropyBitLength           int
		checksumBitLength          int
		lastWordOfEntropyBitLength int
		numberOfCorrectLastWords   int
		someOfCorrectLastBytes     []byte
	}{
		"mnemonic of 12": {
			wordsLength:                12,
			entropyByteLength:          16,
			entropyBitLength:           128,
			checksumBitLength:          4,
			lastWordOfEntropyBitLength: 7,
			numberOfCorrectLastWords:   128,
			someOfCorrectLastBytes:     []byte{0x0, 0x7, 0xf, 0x17, 0x1f, 0x27, 0x2f, 0x37, 0x47, 0x4f, 0x57, 0x5f, 0x67, 0x6f, 0x77, 0x7f},
		},
		"mnemonic of 15": {
			wordsLength:                15,
			entropyByteLength:          20,
			entropyBitLength:           160,
			checksumBitLength:          5,
			lastWordOfEntropyBitLength: 6,
			numberOfCorrectLastWords:   64,
			someOfCorrectLastBytes:     []byte{0x0, 0x3, 0x7, 0xb, 0xf, 0x13, 0x17, 0x1b, 0x23, 0x27, 0x2b, 0x2f, 0x33, 0x37, 0x3b, 0x3f},
		},
		"mnemonic of 18": {
			wordsLength:                18,
			entropyByteLength:          24,
			entropyBitLength:           192,
			checksumBitLength:          6,
			lastWordOfEntropyBitLength: 5,
			numberOfCorrectLastWords:   32,
			someOfCorrectLastBytes:     []byte{0x0, 0x1, 0x3, 0x5, 0x7, 0x9, 0xb, 0xd, 0x11, 0x13, 0x15, 0x17, 0x19, 0x1b, 0x1d, 0x1f},
		},
		"mnemonic of 21": {
			wordsLength:                21,
			entropyByteLength:          28,
			entropyBitLength:           224,
			checksumBitLength:          7,
			lastWordOfEntropyBitLength: 4,
			numberOfCorrectLastWords:   16,
			someOfCorrectLastBytes:     []byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf},
		},
		"mnemonic of 24": {
			wordsLength:                24,
			entropyByteLength:          32,
			entropyBitLength:           256,
			checksumBitLength:          8,
			lastWordOfEntropyBitLength: 3,
			numberOfCorrectLastWords:   8,
			someOfCorrectLastBytes:     []byte{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var (
				entropyByteLength          = test.wordsLength / 3 * 4
				entropyBitLength           = entropyByteLength * 8
				checksumBitLength          = entropyBitLength / 32
				lastWordOfEntropyBitLength = wordEntropyBitLength - checksumBitLength
				numberOfCorrectLastWords   = int(math.Pow(2, float64(lastWordOfEntropyBitLength)))
			)

			assert.Equal(t, test.entropyByteLength, entropyByteLength)
			assert.Equal(t, test.entropyBitLength, entropyBitLength)
			assert.Equal(t, test.checksumBitLength, checksumBitLength)
			assert.Equal(t, test.lastWordOfEntropyBitLength, lastWordOfEntropyBitLength)
			assert.Equal(t, test.numberOfCorrectLastWords, numberOfCorrectLastWords)

			// lets generate some of the correct last bytes
			correctBytes := PossibleLastBytes(entropyByteLength)
			assert.Equal(t, test.someOfCorrectLastBytes, correctBytes)

			end := uint8(test.numberOfCorrectLastWords - 1)
			actualWords := len(correctBytes)
			assert.Equal(t, byte(0), test.someOfCorrectLastBytes[0])
			assert.Equal(t, end, test.someOfCorrectLastBytes[actualWords-1])
			assert.Equal(t, end, correctBytes[actualWords-1])
			span01 := test.someOfCorrectLastBytes[2] - test.someOfCorrectLastBytes[1]
			span7f := test.someOfCorrectLastBytes[actualWords-1] - test.someOfCorrectLastBytes[actualWords-2]
			assert.Equal(t, span01, span7f)
		})
	}
}

func TestLastWordsForLongerPhrases(t *testing.T) {
	tests := map[string]struct {
		phrase string
		length int
	}{
		"test-12": {
			phrase: "test",
			length: 12,
		},
		"test-15": {
			phrase: "test",
			length: 15,
		},
		"test-18": {
			phrase: "test",
			length: 18,
		},
		"test-21": {
			phrase: "test",
			length: 21,
		},
		"test-24": {
			phrase: "test",
			length: 24,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := Request{Phrase: test.phrase, Length: test.length}
			res, err := Main(req)
			assert.NoError(t, err)

			mnemonicWithoutEnd := strings.Join(strings.Fields(res.Body.Mnemonic)[:test.length-1], " ")
			for _, end := range strings.Fields(res.Body.Ends) {
				mnemonic := mnemonicWithoutEnd + " " + end
				assert.True(t, bip39.IsMnemonicValid(mnemonic), "mnemonic is not valid", mnemonic)
			}
		})
	}
}
