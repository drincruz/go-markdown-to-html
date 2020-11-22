package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

// Header is Typical HTML data
// Header data is the "top" part of a site (metadata & such )
type Header struct {
	Title           string
	ContentTitle    string
	ContentSubtitle string
}

// Body is the content of an HTML page.
// For us here, it will be the content parsed in from markdown.
type Body struct {
	Content template.HTML
}

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

func readHeader() []byte {
	data, err := ioutil.ReadFile("bootstrap/clean-blog/header.html.tpl")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func readFooter() []byte {
	data, err := ioutil.ReadFile("bootstrap/clean-blog/footer.html.tpl")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func readFile() []byte {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func header() Header {
	return Header{
		Title:           "drincruz.com Blog",
		ContentTitle:    "This is the main content",
		ContentSubtitle: "and this would be a subtitle",
	}
}

func body(body template.HTML) Body {
	return Body{
		Content: body,
	}
}

func main() {
	var err error
	var header = header()
	var output []byte
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footerData = string(readFooter())
	var content = readFile()
	var body = body(template.HTML(string(MarkdownToHTML(content))))
	var contentStr bytes.Buffer
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/content.html.tpl"))
	contentTpl.Execute(&contentStr, body)

	outputStr.WriteString(headerStr.String())
	outputStr.WriteString(contentStr.String())
	outputStr.WriteString(footerData)

	fmt.Printf("%s", outputStr.String())
	var html []byte
	output = MarkdownToHTML(html)
	var out *os.File
	out = os.Stdout
	if _, err = out.Write(output); err != nil {
		os.Exit(-1)
	}
}
