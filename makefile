.phony: check
check:
	@echo 'hello world'

.phony: version
version:
	go version

.phony: start
start:
	go run .