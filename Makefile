# Don't fear a Makefile

deploy:
	doctl sls connect lambda
	doctl sls deploy src --remote-build

