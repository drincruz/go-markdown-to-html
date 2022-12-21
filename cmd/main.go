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

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func header(title string, contentTitle string, contentSubtitle string, relativePath string) Header {
	return Header{
		Title:           title,
		ContentTitle:    contentTitle,
		ContentSubtitle: contentSubtitle,
		RelativePath:    relativePath,
	}
}

func body(body template.HTML) Body {
	return Body{
		Content: body,
	}
}

func footer(relativePath string) Footer {
	return Footer{
		RelativePath: relativePath,
	}
}

func getTitle() string {
	if os.Args[2] != "" {
		return buildTitle(os.Args[2])
	}
	return "drincruz.com Blog"
}

func buildTitle(title string) string {
	return fmt.Sprintf("%s | drincruz.com", title)
}

func relativePath(path string) string {
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

func writeIndex() {
	// Hardcoding for now
	most_recent([]string{
		"markdown/2022/12/21/asking-for-help.markdown",
		"markdown/2022/10/23/weekly-writing-prompts.markdown",
		"markdown/2021/09/17/optimizing-for-writing.markdown",
	})
}

func writeBlog() {
	var err error
	var relativePath = relativePath(os.Args[4])
	var header = header(getTitle(), os.Args[2], os.Args[3], relativePath)
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footer = footer(relativePath)
	var footerStr bytes.Buffer
	footerTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/footer.html.tpl"))
	footerTpl.Execute(&footerStr, footer)
	var content = readFile(os.Args[1])
	var body = body(template.HTML(string(MarkdownToHTML(content))))
	var contentStr bytes.Buffer
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/content.html.tpl"))
	contentTpl.Execute(&contentStr, body)

	outputStr.WriteString(headerStr.String())
	outputStr.WriteString(contentStr.String())
	outputStr.WriteString(footerStr.String())

	out, err := os.Create(os.Args[4])
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer out.Close()
	_, err = out.WriteString(outputStr.String())
}

func main() {
	if os.Args[1] == "write_index" {
		writeIndex()
		os.Exit(0)
	} else if os.Args[1] == "write_year_archives" {
		yearSummary("2013.json")
		yearSummary("2014.json")
		yearSummary("2015.json")
		yearSummary("2016.json")
		yearSummary("2017.json")
		yearSummary("2018.json")
		yearSummary("2019.json")
		yearSummary("2021.json")
		yearSummary("2022.json")
		os.Exit(0)
	}
	writeBlog()
	os.Exit(0)
}
