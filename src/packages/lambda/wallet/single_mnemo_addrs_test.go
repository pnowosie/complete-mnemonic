package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	AddressCount = 100
	PreRun       = false

	SamplesPath        = "../../../../samples"
	SingleFileNameTmpl = "single%d"
	SingleFileTmpl     = SamplesPath + "/" + SingleFileNameTmpl + ".txt"
	OutputDir          = SamplesPath + "/addresses/single%d/"
)

func TestGenerateAllSingleMnemonicAddresses(t *testing.T) {
	t.Skip("test generates addresses for single-word mnemonics. Single run is enough.")
	t.Parallel()
	lengths := []int{12, 15, 18, 21, 24}

	// read a singleX file from samples
	for _, n := range lengths {
		t.Run("single"+fmt.Sprint(n), closeTestParams(n))
	}
}

func closeTestParams(testCase int) func(t *testing.T) {
	infile := fmt.Sprintf(SingleFileTmpl, testCase)
	outdir := fmt.Sprintf(OutputDir, testCase)

	return func(t *testing.T) {
		t.Parallel()

		fbytes, err := os.ReadFile(infile)
		if err != nil {
			t.Fatal("Cannot open source file", "infile", infile, err)
		}

		for i, line := range strings.Split(string(fbytes), "\n") {
			if PreRun && i >= 5 {
				break
			}

			mnemonic, err := constructFromPhrase(line, testCase)
			if err != nil {
				if strings.HasSuffix(err.Error(), "no words found in ''") {
					break
				}
				t.Fatal("Cannot construct mnemonic", "line", line, err)
			}
			assert.True(t, bip39.IsMnemonicValid(mnemonic), "mnemonic is not valid", "mnemonic", mnemonic)

			addrs, err := generateAddresses(mnemonic, DefaultDerivation, AddressCount)
			fmt.Printf("%-20s %s\n", line, addrs[0])

			outfile := outdir + strings.ReplaceAll(line, " ", "-") + ".txt"
			os.WriteFile(outfile, []byte(strings.Join(addrs, "\n")+"\n"), 0644)
		}
	}
}
