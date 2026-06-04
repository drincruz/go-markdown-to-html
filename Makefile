# build the binary
.PHONY: build
build:
	go build -o go-markdown-to-html cmd/*

# run the build script to generate HTML
.PHONY: html
html:
	bash bin/build_pages.sh

# check and generate third-party licenses
.PHONY: notices
notices:
	bash bin/generate_notices.sh

# run tests
.PHONY: test
test:
	go test -v ./...
