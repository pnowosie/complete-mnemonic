# Don't fear a Makefile
.DEFAULT_GOAL := help

.PHONY: test run deploy show-words single random-wallet wallet
test: ## runs a test of the lambda function
	@cd src/packages/lambda/mnemonix && gotestsum -f testname
	@cd src/packages/lambda/wallet && gotestsum -f testname

WORD := abandon
PHRASE := test_junk
LEN := 12
run: ## sample invocation with doctl CLI, params: WORD=abandon LEN=12 [15,18,21,24] (words delimited by _)
	@doctl sls fn invoke lambda/mnemonix -p phrase:${WORD},length:${LEN}

deploy: ## deploy the lambda function
	doctl sls connect lambda
	doctl sls deploy src --remote-build

show-words: ## show the available english words of the BIP-39 wordlist
	@less +25 src/packages/lambda/mnemonix/bip39/english.go

single: ## creates a single word mnemonic from the BIP-39 samples, params: WORD=abandon LEN=12
	@env LEN=${LEN} WORD=${WORD} bash samples/make-single.sh

random-wallet:
	@doctl sls fn invoke lambda/wallet -p count:5

wallet:
	@doctl sls fn invoke lambda/wallet -p count:5,length:${LEN},phrase:${PHRASE}

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
