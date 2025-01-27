GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)

PROTO_FILES=$(shell find ./api ! -path "./api/third_party*" -name "*.proto")
PROTO_FILES_CHW=$(shell find ./api/blog/v1 -name "*.proto") 


.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/srikrsna/protoc-gen-gotag


.PHONY: proto

# generate internal proto
proto:
	protoc --proto_path=./ \
	--proto_path=./api/third_party \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	--go-http_out=paths=source_relative:. \
	--go-errors_out=paths=source_relative:. \
	--openapiv2_out ./ \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt json_names_for_fields=false \
    --openapiv2_opt enums_as_ints=true \
    $(PROTO_FILES)  \
    && protoc --proto_path=./ \
    --proto_path=./api/third_party \
    --gotag_out=paths=source_relative:. \
    $(PROTO_FILES)


.PHONY: proto-chw
# generate proto
proto-chw:
	protoc --proto_path=./ \
	--proto_path=./api/third_party \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	--go-http_out=paths=source_relative:. \
	--go-errors_out=paths=source_relative:. \
	--openapiv2_out ./ \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt json_names_for_fields=false \
    --openapiv2_opt enums_as_ints=true \
    $(PROTO_FILES_CHW) \
    && protoc --proto_path=./ \
    --proto_path=./api/third_party \
    --gotag_out=paths=source_relative:. \
    $(PROTO_FILES_CHW)

		
.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: test
# test
test:
	go test -v ./... -cover

.PHONY: all
# generate all
all:
	make generate;
	make grpc;
	make http;
	make proto;
	make errors;
	make swagger;
	make build;
	make test;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
