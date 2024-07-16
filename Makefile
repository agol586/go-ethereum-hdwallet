.PHONY: all
all: build

.PHONY: build
build:
	@go build -v -o bin/hdwallet ./cmd/geth-hdwallet 

.PHONY: test
test:
	@go test -v .
