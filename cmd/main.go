package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

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
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
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

// Return the <title> for a page, falling back to the site name.
func pageTitle(title string) string {
	if title != "" {
		return buildTitle(title)
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
	posts := getMostRecent()
	postsArray := getMostRecentArray(posts)
	most_recent(postsArray)
}

func writeBlog(mdPath string, title string, subtitle string, outPath string) {
	var err error
	var relativePath = relativePath(outPath)
	var content = readFile(mdPath)
	htmlContent := string(MarkdownToHTML(content))

	var image string
	if extracted := extractFirstImage(htmlContent); extracted != "" {
		absURL := resolveAbsoluteImageURL(extracted, outPath)
		image = absURL
		log.Printf("[INFO][main][writeBlog] %s, %s", relativePath, absURL)
	}
	var ogUrl = distPathToUrl(outPath)
	var ogType = "article"

	log.Printf("[INFO][main][writeBlog] relativePath: %s", relativePath)
	var description = OpenGraphDescription(postDescription(htmlContent, subtitle))
	var header = NewHeader(pageTitle(title), title, subtitle, relativePath, ogUrl, ogType, OpenGraphImage(image), description)
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footer = footer(relativePath)
	var footerStr bytes.Buffer
	footerTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/footer.html.tpl"))
	footerTpl.Execute(&footerStr, footer)
	var body = body(template.HTML(htmlContent))
	var contentStr bytes.Buffer
	contentTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/content.html.tpl"))
	contentTpl.Execute(&contentStr, body)
	updatedContent := string(UpdateHtmlImgTags(contentStr.Bytes()))

	outputStr.WriteString(headerStr.String())
	outputStr.WriteString(updatedContent)
	outputStr.WriteString(footerStr.String())

	out, err := os.Create(outPath)
	if err != nil {
		log.Printf("Error: failed to create file %s: %s\n", outPath, err)
		return
	}
	defer out.Close()
	_, err = out.WriteString(outputStr.String())
}

func main() {
	switch os.Args[1] {
	case "write_index":
		writeIndex()
		os.Exit(0)
	case "write_posts":
		writePosts(".")
		os.Exit(0)
	case "write_year_archives":
		for _, jsonFile := range yearFiles(".") {
			yearSummary(jsonFile)
		}
		os.Exit(0)
	case "write_seo_files":
		writeSitemap(".", "dist/sitemap.xml")
		writeRobots("dist/robots.txt")
		os.Exit(0)
	}
	writeBlog(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	os.Exit(0)
}
