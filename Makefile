# build the binary
.PHONY: build
build:
	go build -o go-markdown-to-html cmd/*

# check and generate third-party licenses
.PHONY: notices
notices:
	bash bin/generate_notices.sh

# run tests
.PHONY: test
test:
	go test -v ./...
