package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

// Return an HTML-escaped link to a post for its year page.
func yearPostLink(post Blog) string {
	return fmt.Sprintf("<a href=\"%s\">%s</a><br />\n", template.HTMLEscapeString(post.Url), template.HTMLEscapeString(post.Title))
}

func yearSummary(jsonFile string) {
	var err error
	var relativePath string = "./"
	var yearNum string = yearFromJsonFile(jsonFile)
	var ogUrl = distPathToUrl(fmt.Sprintf("%s%s.html", relativePath, yearNum))
	var ogType = "article"
	var description = OpenGraphDescription(fmt.Sprintf("Blog posts from %s on drincruz.com.", yearNum))
	var header = NewHeader(buildTitle(yearNum), "Post Archive", yearNum, relativePath, ogUrl, ogType, description)
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footer = footer(relativePath)
	var footerStr bytes.Buffer
	footerTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/footer.html.tpl"))
	footerTpl.Execute(&footerStr, footer)

	var postLinks string
	for _, post := range readYearPosts(jsonFile) {
		postLinks += yearPostLink(post)
	}

	var contentStr string
	contentStr = cardContent(postLinks)

	outputStr.WriteString(headerStr.String())
	outputStr.WriteString(contentStr)
	outputStr.WriteString(footerStr.String())

	var outputFilename string = fmt.Sprintf("dist/%s.html", yearNum)
	out, err := os.Create(outputFilename)
	if err != nil {
		error := fmt.Errorf("[ERROR] Could not create the output file: %s. Error: %v", outputFilename, err.Error())
		fmt.Println(error)
		return
	}
	defer out.Close()
	_, err = out.WriteString(outputStr.String())
}
