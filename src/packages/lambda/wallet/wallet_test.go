package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyDerivation(t *testing.T) {
	tests := map[string]struct {
		req              *Request
		expectedResponse *Response
	}{
		"derive from phrase": {
			req: &Request{
				Phrase: "fox_frown",
				Count:  3,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "fox fox fox fox fox fox fox fox fox fox fox frown",
					},
					Accounts: []AccountBody{
						{
							Address: "0x1023e8DbDebAd480C43f6e19b3381c465c74E933",
						},
						{
							Address: "0x977608A20f221f31D0FA10b22664511343CfB3A1",
						},
						{
							Address: "0x7a307954D1337af50c00Aa0e2Dbe92Dd9CcfBA80",
						},
					},
				},
			},
		},
		"derive from mnemonic": {
			req: &Request{
				Mnemonic: "wish_wish_wish_wish_wish_wish_wish_wish_wish_wish_wish_wool",
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "wish wish wish wish wish wish wish wish wish wish wish wool",
					},
					Accounts: []AccountBody{
						{
							Address: "0x59002b96fdf144FCE4F2d357A9978770779E589F",
						},
						{
							Address: "0xAD83edfB4953a2Bd42699D9c72cdf86130E23317",
						},
						{
							Address: "0xCEDa6874fE9007DcFc66EA85E60AA1569D577792",
						},
						{
							Address: "0x516e74fd137F854d5dd75f2702FF00f97fc94CCc",
						},
						{
							Address: "0x4Fe9b7bc50f0a2bc986D438eac366c11eE3CeF55",
						},
						{
							Address: "0x2953621c746EBA0eFDD755Dde1e57fAd364302C0",
						},
						{
							Address: "0x7B7DD69d0b096cF73a7E28D61F51cfdFbDE2914B",
						},
						{
							Address: "0x2BaD8191514DE9F76f6102775cC49B9FCa15181f",
						},
						{
							Address: "0xa3c16BFe270ea9054c7464995617f491e7602Db2",
						},
						{
							Address: "0xb24069bCeE29200FAbadaBf4e8f96E3EDb05b257",
						},
					},
				},
			},
		},
		"two words phrase with a checksum word 1_6": {
			req: &Request{
				Phrase: "one six since",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "one six one six one six one six one six one since",
					},
					Accounts: []AccountBody{
						{
							Address: "0xC634fB51Ee91E771066737fbd483e5EF8b6275AE",
						},
					},
				},
			},
		},
		"two words phrase with a checksum word fox_6": {
			req: &Request{
				Phrase: "fox six silly",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "fox six fox six fox six fox six fox six fox silly",
					},
					Accounts: []AccountBody{
						{
							Address: "0x7d347F41F826d8d95A41e41298Dbdc60fa3435C4",
						},
					},
				},
			},
		},
		"test junk phrase": {
			req: &Request{
				Phrase: "test junk",
				Count:  3,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "test test test test test test test test test test test junk",
					},
					Accounts: []AccountBody{
						{
							Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
						},
						{
							Address: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
						},
						{
							Address: "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
						},
					},
				},
			},
		},
		"test junk phrase with password": {
			req: &Request{
				Phrase:   "test junk",
				Password: "password",
				Count:    3,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "test test test test test test test test test test test junk",
					},
					Accounts: []AccountBody{
						{
							Address: "0xfaFfA9053ac6c6315Aa7806d1336F10F9b280Ee9",
						},
						{
							Address: "0xb1a3B55051E04d44Ce457A6A479c999557521921",
						},
						{
							Address: "0x64ffA20464c6dF3b23f1540327578eBf10C23785",
						},
					},
				},
			},
		},
		"mnemonix generated QBF-1": {
			req: &Request{
				Phrase: "quick brown fox attack",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "quick brown fox quick brown fox quick brown fox quick brown attack",
					},
					Accounts: []AccountBody{
						{
							Address: "0xC81E455a82d2029E2ecDdFaA6365B15CD69589a5",
						},
					},
				},
			},
		},
		"mnemonix generated QBF-2": {
			req: &Request{
				Phrase: "quick_brown_fox_cable",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "quick brown fox quick brown fox quick brown fox quick brown cable",
					},
					Accounts: []AccountBody{
						{
							Address: "0x523063b46e87d419c4b30402170C3ED91dCefD6A",
						},
					},
				},
			},
		},
		"mnemonix generated QBF-3": {
			req: &Request{
				Phrase: "quick brown fox oyster",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "quick brown fox quick brown fox quick brown fox quick brown oyster",
					},
					Accounts: []AccountBody{
						{
							Address: "0x588D620acE82cC864976bD3Cfb44Fdb33DCe0ED4",
						},
					},
				},
			},
		},
		"mnemonix generated BFQ-1": {
			req: &Request{
				Phrase: "brown fox quick wrong",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "brown fox quick brown fox quick brown fox quick brown fox wrong",
					},
					Accounts: []AccountBody{
						{
							Address: "0xa601FAb390f54318642F2e5f9fe4584F7502A769",
						},
					},
				},
			},
		},
		"mnemonix generated BFQ-2": {
			req: &Request{
				Phrase: "brown fox quick garage",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "brown fox quick brown fox quick brown fox quick brown fox garage",
					},
					Accounts: []AccountBody{
						{
							Address: "0x5cD325FeeefaBc5f91C856c71d46a923F9235cE4",
						},
					},
				},
			},
		},
		"mnemonix generated BFQ-3": {
			req: &Request{
				Phrase: "brown_fox_quick_bullet",
				Count:  1,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "brown fox quick brown fox quick brown fox quick brown fox bullet",
					},
					Accounts: []AccountBody{
						{
							Address: "0x11aaa3bfdc8c6669002fb74ABbc33adf4b7cfb92",
						},
					},
				},
			},
		},
		"test junk phrase with private keys revealed": {
			req: &Request{
				Phrase:        "test junk",
				Count:         2,
				RevealPrivate: true,
			},
			expectedResponse: &Response{
				StatusCode: 200,
				Body: ResponseBody{
					Wallet: WalletBody{
						Derivation: DefaultDerivation,
						Length:     DefaultPhraseLength,
						Mnemonic:   "test test test test test test test test test test test junk",
					},
					Accounts: []AccountBody{
						{
							Address:    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
							PublicKey:  "8318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5",
							PrivateKey: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
						},
						{
							Address:    "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
							PublicKey:  "ba5734d8f7091719471e7f7ed6b9df170dc70cc661ca05e688601ad984f068b0d67351e5f06073092499336ab0839ef8a521afd334e53807205fa2f08eec74f4",
							PrivateKey: "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		// set defaults
		req := test.req
		req.Length = DefaultPhraseLength
		req.Derivation = DefaultDerivation
		if req.Count == 0 {
			req.Count = DefaultAccountCount
		}
		t.Run(name, func(t *testing.T) {
			resp, err := Main(*test.req)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, test.expectedResponse, resp)
		})
	}
}

func TestInvalidInputsErrors(t *testing.T) {
	tests := map[string]struct {
		req           *Request
		expectedCode  int
		expectedError string
	}{
		"mnemonic instead of phrase": {
			req: &Request{
				Mnemonic: "fox_six_skill",
				Length:   18,
			},
			expectedCode:  400,
			expectedError: "Invalid mnemonic. Don't you mean 'phrase' instead of 'mnemonic'?",
		},
	}

	for name, test := range tests {
		// set defaults
		req := test.req
		req.Derivation = DefaultDerivation
		if req.Length == 0 {
			req.Length = DefaultPhraseLength
		}
		if req.Count == 0 {
			req.Count = DefaultAccountCount
		}
		t.Run(name, func(t *testing.T) {
			resp, err := Main(*test.req)
			if err != nil {
				t.Fatal("unexpected error, which should be reported by response")
			}
			assert.Equal(t, test.expectedCode, resp.StatusCode)
			assert.Equal(t, test.expectedError, resp.Body.Error)
		})
	}
}

func TestNewMnemonicGeneratedAtRandom(t *testing.T) {
	resp, err := Main(Request{Count: 3, Derivation: DefaultDerivation, Length: DefaultPhraseLength, RevealPrivate: true})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Body.Wallet.Mnemonic)
	assert.Equal(t, 3, len(resp.Body.Accounts))
	for _, addr := range resp.Body.Accounts {
		assert.Equal(t, 42, len(addr.Address))
		fmt.Println("-", addr.Address, addr.PrivateKey)
	}
}
