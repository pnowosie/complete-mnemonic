# Don't fear a Makefile

deploy:
	doctl sls connect lambda
	doctl sls deploy src --remote-build

WORD := abandon
LEN := 12
run:
	# single word invocation is as easy as
	doctl sls fn invoke lambda/mnemonix -p phrase:${WORD},length:${LEN}
