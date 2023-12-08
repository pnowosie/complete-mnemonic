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
	in.AssumeDefaults()
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
		parsedMnemonic, length, err := parseMnemonic(in.Mnemonic)
		if err != nil {
			fmt.Println("error in parsing mnemonic", quote(in.Mnemonic), "error", err)
			return &Response{
				StatusCode: http.StatusBadRequest,
				Body: ResponseBody{
					Error: err.Error(),
				},
			}, nil
		}
		in.Mnemonic, in.Length = parsedMnemonic, length
	}

	if in.Mnemonic == "" && in.Phrase != "" {
		mnemonic, err := constructFromPhrase(in.Phrase, in.Length)
		if err != nil {
			fmt.Println("error constructing mnemonic from phrase", "phrase", quote(in.Phrase), "length", in.Length, "error", err)
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
		fmt.Println("given mnemonic is not valid", "phrase", quote(in.Phrase), "length", in.Length, "mnemonic", quote(in.Mnemonic))
		return &Response{
			StatusCode: http.StatusBadRequest,
			Body: ResponseBody{
				Error: "Invalid mnemonic. Don't you mean 'phrase' instead of 'mnemonic'?",
			},
		}, nil
	}

	genAccounts, err := generateAddresses(in.Mnemonic, in.Password, in.Derivation, in.Count, in.RevealPrivate)
	if err != nil {
		return &Response{
			StatusCode: http.StatusInternalServerError,
			Body: ResponseBody{
				Error: err.Error(),
			},
		}, nil
	}

	return &Response{
		StatusCode: http.StatusOK,
		Body: ResponseBody{
			Wallet: WalletBody{
				Mnemonic:   in.Mnemonic,
				Derivation: in.Derivation,
				Length:     in.Length,
			},
			Accounts: genAccounts,
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

func generateAddresses(mnemonic, password, derivation string, count int, includePrivate bool) ([]AccountBody, error) {
	wallet, err := hd.NewFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}

	accs := make([]AccountBody, count)
	//fmt.Println("Generating accs:", "count", count)
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
		accs[i].Address = account.Address.Hex()
		if includePrivate {
			pub, _ := wallet.PublicKeyHex(account)
			priv, _ := wallet.PrivateKeyHex(account)
			accs[i].PublicKey = pub
			accs[i].PrivateKey = priv
		}
	}

	return accs, nil
}

func parseMnemonic(mnemonic string) (string, int, error) {
	words, err := toWordList(mnemonic)
	if err != nil {
		return "", 0, err
	}
	return strings.Join(words, " "), len(words), nil
}
