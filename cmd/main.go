package main

import (
	"os"

	"github.com/gomarkdown/markdown"
)

func markdownToHTML(md []byte) []byte {
	return markdown.ToHTML(md, nil, nil)
}

func main() {
	var err error
	var output []byte
	output = markdownToHTML([]byte("# This is a title `right` *ok* _really_?"))
	var out *os.File
	out = os.Stdout
	if _, err = out.Write(output); err != nil {
		os.Exit(-1)
	}
}
