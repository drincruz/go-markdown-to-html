# go-markdown-to-html
Markdown to HTML Converter.

It takes markdown and converts it to HTML. :)

## How to Build
To build the `go-markdown-to-html` binary, simply run `go build -o go-markdown-to-html cmd/*`.

## Running Tests
To run the tests, you can run `go test -v ./...`.

## Generating the Static HTML
After building the binary (`go-markdown-to-html`), you can just run `./bin/build_pages.sh` to generate the HTML files. These files will be generated in the `dist` directory.

## Managing Third-Party Licenses

This project tracks and checks compliance of third-party licenses using `go-licenses`. To generate/update the [THIRD_PARTY_NOTICES.md](file:///home/drin/src/github.com/drincruz/go-markdown-to-html/THIRD_PARTY_NOTICES.md) file and save local license files, run:

```bash
make notices
```

This will:
1. Verify all dependencies comply with allowed open-source licenses.
2. Save raw license files to the `third_party/` directory (git-ignored).
3. Regenerate [THIRD_PARTY_NOTICES.md](file:///home/drin/src/github.com/drincruz/go-markdown-to-html/THIRD_PARTY_NOTICES.md) using the custom template in `bin/notices.tmpl`.
