package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/pnowosie/complete-mnemonic/bip39"
	"github.com/stretchr/testify/assert"
)

const (
	PreRun = false

	SamplesPath        = "../../../../samples"
	SingleFileNameTmpl = "single%d"
	SingleFileTmpl     = SamplesPath + "/" + SingleFileNameTmpl + ".txt"
)

func TestGenerateAllSingleMnemonicAddresses(t *testing.T) {
	t.Skip("test generates single-word mnemonics. Single run is enough.")
	t.Parallel()
	lengths := []int{12, 15, 18, 21, 24}

	// read a singleX file from samples
	for _, n := range lengths {
		t.Run("single"+fmt.Sprint(n), closeTestParams(n))
	}
}

func closeTestParams(testCase int) func(t *testing.T) {
	outfile := fmt.Sprintf(SingleFileTmpl, testCase)

	return func(t *testing.T) {
		t.Parallel()

		words := []string{}
		for i, word := range bip39.GetWordList() {
			if PreRun && i >= 5 {
				break
			}

			mn, _ := Repeat(word, testCase)
			en, _ := bip39.EntropyFromMnemonic(mn)
			mn, _ = bip39.NewMnemonic(en)

			mWs := strings.Fields(mn)
			assert.Equalf(t, word, mWs[0], "first word is not the same", "word", word, "first", mWs[0])
			words = append(words, mWs[0]+" "+mWs[len(mWs)-1])
		}
		os.WriteFile(outfile, []byte(strings.Join(words, "\n")+"\n"), 0644)
	}
}
