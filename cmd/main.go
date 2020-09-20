package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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

func readFile() []byte {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	var err error
	var output []byte
	var data []byte = readFile()
	output = MarkdownToHTML(data)
	var out *os.File
	out = os.Stdout
	if _, err = out.Write(output); err != nil {
		os.Exit(-1)
	}
}
