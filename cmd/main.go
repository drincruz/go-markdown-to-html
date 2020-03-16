package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
)

// MarkdownToHTML transforms markdown to HTML.
//
// Takes markdown as a byte array argument, returns a byte array of HTML.
func MarkdownToHTML(md []byte) []byte {
	return markdown.ToHTML(md, nil, nil)
}

func readStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(MarkdownToHTML(scanner.Bytes()))
	}
}

func main() {
	var err error
	var output []byte
	output = MarkdownToHTML([]byte("# This is a title `right` *ok* _really_?"))
	var out *os.File
	out = os.Stdout
	if _, err = out.Write(output); err != nil {
		os.Exit(-1)
	}
}
