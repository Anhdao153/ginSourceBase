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

.phony: grpc-swagger
grpc-swagger:
	protoc -I . --openapiv2_out ./rpc/gen rpc/$(PARAM)/$(PARAM).proto


# just re-generate the Go client * generate 
# sync the database with your schema for development * db push
# create a prisma schema from your existing database * db pull
# for production use, create a migration locally * migrate dev
# sync your production database with your migrations * migrate deploy
.phony: prisma
prisma:
	go run github.com/steebchen/prisma-client-go $(PARAM)

.phony: seed
seed: 
	go run seeds/prismaSeed.go