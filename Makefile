# Don't fear a Makefile
.DEFAULT_GOAL := help

.PHONY: help show-words test run deploy random-wallet wallet

WORD := abandon
PHRASE := test_junk
LEN := 12

##@ Usage
help: ## display this helpful message
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


show-words: ## show the available english words of the BIP-39 wordlist
	@less +25 src/packages/lambda/mnemonix/bip39/english.go

run: ## sample invocation with doctl CLI, params: WORD=abandon LEN=12 [15,18,21,24] (words delimited by _)
	@doctl sls fn invoke lambda/mnemonix -p phrase:${WORD},length:${LEN}

random-wallet: ## runs wallet fn to generate at random
	@doctl sls fn invoke lambda/wallet -p count:5

wallet: ## runs wallet fn with count=5, params: LEN=12 and PHRASE=test_junk
	@doctl sls fn invoke lambda/wallet -p count:5,length:${LEN},phrase:${PHRASE}

##@ Develop

test: ## runs a test of the lambda function
	@cd src/packages/lambda/mnemonix && gotestsum -f testname
	@cd src/packages/lambda/wallet && gotestsum -f testname

deploy: ## deploy the lambda function
	doctl sls connect lambda
	doctl sls deploy src --remote-build

