package main

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/pnowosie/complete-mnemonic/bip39"
	"github.com/stretchr/testify/assert"
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
		"another-three-12-success": {
			req: &Request{
				Phrase: "quick brown fox",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Mnemonic: "quick brown fox quick brown fox quick brown fox quick brown fox",
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
			correctBytes := PossibleLastBytes(entropyByteLength, 0x0, 16)
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
		phrase   string
		length   int
		endWords int
	}{
		"test-12": {
			phrase:   "test",
			length:   12,
			endWords: 128,
		},
		"test-15": {
			phrase:   "test",
			length:   15,
			endWords: 64,
		},
		"test-18": {
			phrase:   "test",
			length:   18,
			endWords: 32,
		},
		"test-21": {
			phrase:   "test",
			length:   21,
			endWords: 16,
		},
		"test-24": {
			phrase:   "test",
			length:   24,
			endWords: 8,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := Request{
				Phrase:   test.phrase,
				Length:   test.length,
				EndWords: test.endWords}
			res, err := Main(req)
			assert.NoError(t, err)

			mnemonicWithoutEnd := strings.Join(strings.Fields(res.Body.Mnemonic)[:test.length-1], " ")
			//fmt.Println(req.Phrase, ":", res.Body.Ends)
			for _, end := range strings.Fields(res.Body.Ends) {
				mnemonic := mnemonicWithoutEnd + " " + end
				assert.True(t, bip39.IsMnemonicValid(mnemonic), "mnemonic is not valid", mnemonic)
			}
		})
	}
}

func TestEntropyDoesNotContainChecksum(t *testing.T) {
	tests := map[string]struct {
		phraseLength          int
		obtainedEntropyLength int
	}{
		"test-12": {
			phraseLength:          12,
			obtainedEntropyLength: 16,
		},
		"test-15": {
			phraseLength:          15,
			obtainedEntropyLength: 20,
		},
		"test-18": {
			phraseLength:          18,
			obtainedEntropyLength: 24,
		},
		"test-21": {
			phraseLength:          21,
			obtainedEntropyLength: 28,
		},
		"test-24": {
			phraseLength:          24,
			obtainedEntropyLength: 32,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bitSize := test.phraseLength*11 - test.phraseLength/3
			entropy, err := bip39.NewEntropy(bitSize)
			assert.NoError(t, err)
			assert.Equal(t, test.obtainedEntropyLength, len(entropy))

			mn, err := bip39.NewMnemonic(entropy)
			assert.NoError(t, err)
			bytes, _ := bip39.MnemonicToByteArray(mn, false)

			// only then we can have access to checksum byte
			assert.Equal(t, test.obtainedEntropyLength+1, len(bytes))
		})
	}
}

// Observation: last byte of entropy is not used for checksum
// to get the checksum you need MnemonicToByteArray method
// checksum is a few first bits from the first byte of the hash of the entropy
func TestHowChecksumBitsArePlacedInBytesArr(t *testing.T) {
	tests := map[string]struct {
		phraseLength int
		wordIndex    int
		firstByte    byte
		butLastByte  byte
		lastByte     byte
	}{
		"test-A-12": {
			phraseLength: 12,
			wordIndex:    0,
			firstByte:    0x0,
			butLastByte:  0x0,
			lastByte:     0x3,
		},
		"test-Z-12": {
			phraseLength: 12,
			wordIndex:    2047,
			firstByte:    0x0f,
			butLastByte:  0xff,
			lastByte:     0xf5,
		},
		"test-A-15": {
			phraseLength: 15,
			wordIndex:    0,
			firstByte:    0x0,
			butLastByte:  0x0,
			lastByte:     0x1b,
		},
		"test-Z-15": {
			phraseLength: 15,
			wordIndex:    2047,
			firstByte:    0x1f,
			butLastByte:  0xff,
			lastByte:     0xf3,
		},
		"test-A-18": {
			phraseLength: 18,
			wordIndex:    0,
			firstByte:    0x0,
			butLastByte:  0x0,
			lastByte:     0x27,
		},
		"test-Z-18": {
			phraseLength: 18,
			wordIndex:    2047,
			firstByte:    0x3f,
			butLastByte:  0xff,
			lastByte:     0xd1,
		},
		"test-A-21": {
			phraseLength: 21,
			wordIndex:    0,
			firstByte:    0x0,
			butLastByte:  0x0,
			lastByte:     0x1d,
		},
		"test-Z-21": {
			phraseLength: 21,
			wordIndex:    2047,
			firstByte:    0x7f,
			butLastByte:  0xff,
			lastByte:     0x99,
		},
		"test-A-24": {
			phraseLength: 24,
			wordIndex:    0,
			firstByte:    0x0,
			butLastByte:  0x0,
			lastByte:     0x66,
		},
		"test-Z-24": {
			phraseLength: 24,
			wordIndex:    2047,
			firstByte:    0xff,
			butLastByte:  0xff,
			lastByte:     0xaf,
		},
	}
	for name, test := range tests {
		words := bip39.GetWordList()
		t.Run(name, func(t *testing.T) {
			res, _ := Main(Request{Phrase: words[test.wordIndex], Length: test.phraseLength})
			phrase := res.Body.Mnemonic
			bytes, _ := bip39.MnemonicToByteArray(phrase, false)
			//fmt.Println(phrase)
			//fmt.Printf("%d: %x\n", test.phraseLength, bytes)
			//entropy, _ := bip39.EntropyFromMnemonic(phrase)
			//fmt.Printf("%d: %x\n", test.phraseLength, entropy)
			assert.Equal(t, test.firstByte, bytes[0])
			assert.Equal(t, test.butLastByte, bytes[len(bytes)-2])
			assert.Equal(t, test.lastByte, bytes[len(bytes)-1])
		})
	}
}

func TestPossibleLastBytesPreservesMask(t *testing.T) {
	tests := map[string]struct {
		entropyLength int
		lastByte      byte
		freeBits      int
	}{
		"test-12": {
			entropyLength: 16,
			lastByte:      0b10000000,
			freeBits:      7,
		},
		"test-15": {
			entropyLength: 20,
			lastByte:      0b01000000,
			freeBits:      6,
		},
		"test-18": {
			entropyLength: 24,
			lastByte:      0b10100000,
			freeBits:      5,
		},
		"test-21": {
			entropyLength: 28,
			lastByte:      0b10010000,
			freeBits:      4,
		},
		"test-24": {
			entropyLength: 32,
			lastByte:      0b10001000,
			freeBits:      3,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			possibleChecksums := 1 << test.freeBits
			bytes := PossibleLastBytes(test.entropyLength, test.lastByte, possibleChecksums)
			lastByte := fmt.Sprintf("%08b", test.lastByte)
			preserve := lastByte[:8-test.freeBits]
			fmt.Printf("%d: free %d bits (%d) %08b\n", test.entropyLength, test.freeBits, 1<<test.freeBits, test.lastByte)
			for _, b := range bytes {
				bib := fmt.Sprintf("%08b", b)
				//fmt.Println(" -", bib)
				assert.True(t, strings.HasPrefix(bib, preserve), "lastByte not preserved", lastByte, bib)
			}
		})
	}
}

func TestCanSpecifyExactlyEndWords(t *testing.T) {
	tests := map[string]struct {
		length           int
		endWords         int
		expectedReturned int
	}{
		"can supress end words": {
			length:           12,
			endWords:         0,
			expectedReturned: 0,
		},
		"can specify end words": {
			length:           15,
			endWords:         3,
			expectedReturned: 3,
		},
		"can specify more possible": {
			length:           18,
			endWords:         100,
			expectedReturned: 32,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := Request{
				Phrase:   "zoo",
				Length:   test.length,
				EndWords: test.endWords,
			}
			res, err := Main(req)
			assert.NoError(t, err)
			assert.Empty(t, res.Body.Error)
			assert.Equal(t, test.expectedReturned, len(strings.Fields(res.Body.Ends)))
		})
	}
}
