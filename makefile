.phony: check
check:
	@echo 'hello world'

.phony: version
version:
	go version

.phony: start
start:
	go run .

.phony: grpc
grpc:
	protoc rpc/$(PARAM)/$(PARAM).proto --go_out=. --go-grpc_out=.
