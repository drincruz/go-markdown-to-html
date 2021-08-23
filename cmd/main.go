package main

import (
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
	RelativePath    string
}

// Body is the content of an HTML page.
// For us here, it will be the content parsed in from markdown.
type Body struct {
	Content template.HTML
}

// Footer is Typical footer Copyright, links, etc.
type Footer struct {
	RelativePath string
}

// MarkdownToHTML transforms markdown to HTML.
//
// Takes markdown as a byte array argument, returns a byte array of HTML.
func MarkdownToHTML(md []byte) []byte {
	return markdown.ToHTML(md, nil, nil)
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
		Title:           getTitle(),
		ContentTitle:    os.Args[2],
		ContentSubtitle: os.Args[3],
		RelativePath:    relativePath(),
	}
}

func body(body template.HTML) Body {
	return Body{
		Content: body,
	}
}

func footer() Footer {
	return Footer{
		RelativePath: relativePath(),
	}
}

func getTitle() string {
	if os.Args[2] != "" {
		return fmt.Sprintf("%s | drincruz.com", os.Args[2])
	}
	return "drincruz.com Blog"

}
func relativePath() string {
	var path string = os.Args[4]
	var outputArr []string
	for i := 1; i < strings.Count(path, "/"); i++ {
		outputArr = append(outputArr, "../")
	}

	var output string = strings.Join(outputArr, "")

	if output == "" {
		return "./"
	}
	return output
}

func main() {
	var err error
	var header = header()
	var output []byte
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footer = footer()
	var footerStr bytes.Buffer
	footerTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/footer.html.tpl"))
	footerTpl.Execute(&footerStr, footer)
	var content = readFile()
	var body = body(template.HTML(string(MarkdownToHTML(content))))
	var contentStr bytes.Buffer
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/content.html.tpl"))
	contentTpl.Execute(&contentStr, body)

	outputStr.WriteString(headerStr.String())
	outputStr.WriteString(contentStr.String())
	outputStr.WriteString(footerStr.String())

	fmt.Printf("%s", outputStr.String())
	var html []byte
	output = MarkdownToHTML(html)
	var out *os.File
	out = os.Stdout
	if _, err = out.Write(output); err != nil {
		os.Exit(-1)
	}
}
