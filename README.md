# go-markdown-to-html
Markdown to HTML Converter.

It takes markdown and converts it to HTML. :)

## How to Build
To build the `go-markdown-to-html` binary, simply run `go build -o go-markdown-to-html cmd/*`.

## Running Tests
To run the tests, you can run `go test -v ./...`.

## Generating the Static HTML
After building the binary (`go-markdown-to-html`), you can just run `./bin/build_pages.sh` to generate the HTML files. These files will be generated in the `dist` directory.
