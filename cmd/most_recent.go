package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Card struct {
	Content template.HTML
	Date    string
	Link    string
	Title   template.HTML
}

func linkFromMarkdown(markdown string) string {
	newStr := strings.Replace(markdown, "markdown/", "", -1)
	newStr = strings.Replace(newStr, ".markdown", ".html", -1)

	return newStr
}

func distFileFromMarkdown(markdown string) string {
	newStr := strings.Replace(markdown, "markdown/", "dist/", -1)
	newStr = strings.Replace(newStr, ".markdown", ".html", -1)

	return newStr
}

func dateFromPath(path string) string {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	parts := strings.Split(path, "/")
	if len(parts) > 3 {
		date := fmt.Sprintf("%s-%s-%s", parts[1], parts[2], parts[3])
		t, _ := time.Parse(layoutISO, date)
		return fmt.Sprintf("%s, %d %s %d", t.Weekday(), t.Year(), t.Month(), t.Day())
	}
	// No date, so return one part.
	return fmt.Sprintf("%s", parts[1])
}

func contentMulti(markdownFilename string) string {
	var contentStr bytes.Buffer
	var content, _ = os.Open(distFileFromMarkdown(markdownFilename))
	doc, _ := html.Parse(content)
	p, _ := cardTag(doc, "p")
	title, _ := cardTag(doc, "h1")
	var card = Card{
		Content: renderNode(p),
		Date:    dateFromPath(markdownFilename),
		Link:    linkFromMarkdown(markdownFilename),
		Title:   renderNode(title),
	}
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/most_recent_card.html.tpl"))
	contentTpl.Execute(&contentStr, card)

	return contentStr.String()
}

func cardTag(doc *html.Node, tag string) (*html.Node, error) {
	var tags []*html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			tags = append(tags, node)
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if len(tags) > 0 {
		return firstNoImgTag(tags), nil
	}
	var errorMsg string = fmt.Sprintf("Missing <%s> in the node tree", tag)
	return nil, errors.New(errorMsg)
}

// This can use some cleanup
func firstNoImgTag(tags []*html.Node) *html.Node {
	for _, tag := range tags {
		for child := tag.FirstChild; child != nil; child = child.NextSibling {
			if child.Data != "img" {
				return tag
			}
		}
	}
	// default to empty p tag
	return &html.Node{Type: html.ElementNode, Data: "p", DataAtom: atom.P}
}

func renderNode(n *html.Node) template.HTML {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return template.HTML(buf.String())
}

func cardContent(cards string) string {
	var contentStr bytes.Buffer
	var body = body(template.HTML(cards))
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/most_recent_content.html.tpl"))
	contentTpl.Execute(&contentStr, body)

	return contentStr.String()
}

func most_recent(posts []string) {
	var err error
	var relativePath string = "./"
	var header = header(buildTitle("Welcome!"), "Greetings!", "My name is Adrian", relativePath)
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footer = footer(relativePath)
	var footerStr bytes.Buffer
	footerTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/footer.html.tpl"))
	footerTpl.Execute(&footerStr, footer)

	var content = readFile("markdown/index.markdown")
	var body = body(template.HTML(string(MarkdownToHTML(content))))
	var mainContentStr bytes.Buffer
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/index_content.html.tpl"))
	contentTpl.Execute(&mainContentStr, body)

	var cardsStr string
	for _, post := range posts {
		cardsStr += contentMulti(post)
	}

	var contentStr string
	contentStr = cardContent(cardsStr)

	outputStr.WriteString(headerStr.String())
	outputStr.WriteString(mainContentStr.String())
	outputStr.WriteString(contentStr)
	outputStr.WriteString(footerStr.String())

	out, err := os.Create("dist/index.html")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer out.Close()
	_, err = out.WriteString(outputStr.String())
}
