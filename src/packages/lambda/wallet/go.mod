module github.com/pnowosie/complete-mnemonic

go 1.20

require (
	github.com/ethereum/go-ethereum v1.10.17
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/stretchr/testify v1.8.4
	github.com/tyler-smith/go-bip39 v1.1.0
)

require (
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/miguelmota/go-ethereum-hdwallet => github.com/pnowosie/go-ethereum-hdwallet v0.0.0-20231114151632-c79963b23f11
