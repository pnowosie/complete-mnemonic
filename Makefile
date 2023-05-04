# Don't fear a Makefile
.DEFAULT_GOAL := help

.PHONY: test run deploy
test: ## runs a test of the lambda function
	@cd src/packages/lambda/mnemonix && gotestsum -f testname

WORD := abandon
LEN := 12
run: ## sample invocation with doctl CLI, params: WORD=abandon LEN=12
	# single word invocation is as easy as
	doctl sls fn invoke lambda/mnemonix -p phrase:${WORD},length:${LEN}

deploy: ## deploy the lambda function
	doctl sls connect lambda
	doctl sls deploy src --remote-build

show-words: ## show the available english words of the BIP-39 wordlist
	@less +25 src/packages/lambda/mnemonix/bip39/english.go

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
