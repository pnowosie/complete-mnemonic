package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	hd "github.com/miguelmota/go-ethereum-hdwallet"
	bip39 "github.com/tyler-smith/go-bip39"
)

func Main(in Request) (*Response, error) {
	if in.Mnemonic == "" && in.Phrase == "" {
		mnemonic, err := randomMnemonic(in)
		if err != nil {
			return &Response{
				StatusCode: http.StatusInternalServerError,
				Body: ResponseBody{
					Error: err.Error(),
				},
			}, nil
		}
		in.Mnemonic = mnemonic
	}

	if in.Mnemonic != "" {
		parsedMnemonic, err := parseMnemonic(in.Mnemonic)
		if err != nil {
			return &Response{
				StatusCode: http.StatusBadRequest,
				Body: ResponseBody{
					Error: err.Error(),
				},
			}, nil
		}
		in.Mnemonic = parsedMnemonic
	}

	if in.Mnemonic == "" && in.Phrase != "" {
		mnemonic, err := constructFromPhrase(in.Phrase, in.Length)
		if err != nil {
			return &Response{
				StatusCode: http.StatusInternalServerError,
				Body: ResponseBody{
					Error: err.Error(),
				},
			}, nil
		}
		in.Mnemonic = mnemonic
	}

	if !bip39.IsMnemonicValid(in.Mnemonic) {
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body: ResponseBody{
				Error: "invalid mnemonic",
			},
		}, nil
	}

	addrs, err := generateAddresses(in.Mnemonic, in.Derivation, in.Count)
	if err != nil {
		return &Response{
			StatusCode: http.StatusInternalServerError,
			Body: ResponseBody{
				Error: err.Error(),
			},
		}, nil
	}
	addresses := make([]AccountBody, len(addrs))
	for i, addr := range addrs {
		addresses[i] = AccountBody{
			Address: addr,
		}
	}

	return &Response{
		StatusCode: http.StatusOK,
		Body: ResponseBody{
			Wallet: WalletBody{
				Mnemonic:   in.Mnemonic,
				Derivation: DefaultDerivation,
				Length:     DefaultPhraseLength,
			},
			Accounts: addresses,
		},
	}, nil
}

func randomMnemonic(in Request) (string, error) {
	entropyBits := in.Length*11 - in.Length/3
	entropy := make([]byte, entropyBits/8)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

func constructFromPhrase(phrase string, length int) (string, error) {
	words, err := toWordList(phrase)
	if err != nil {
		return "", err
	}

	last := words[len(words)-1]
	mnemonic, err := Repeat(strings.Join(words[:len(words)-1], " "), length)
	if err != nil {
		return "", err
	}
	mnemonicWords, _ := toWordList(mnemonic)
	mnemonicWords[len(mnemonicWords)-1] = last
	return strings.Join(mnemonicWords, " "), nil
}

func generateAddresses(mnemonic string, derivation string, count int) ([]string, error) {
	wallet, err := hd.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	addresses := make([]string, count)
	for i := 0; i < count; i++ {
		prePath := fmt.Sprintf("%s%d", derivation, i)
		path, err := accounts.ParseDerivationPath(prePath)
		if err != nil {
			return nil, err
		}
		account, err := wallet.Derive(path, false)
		if err != nil {
			return nil, err
		}
		addresses[i] = account.Address.Hex()
	}

	return addresses, nil
}

func parseMnemonic(mnemonic string) (string, error) {
	words, err := toWordList(mnemonic)
	if err != nil {
		return "", err
	}
	return strings.Join(words, " "), nil
}
