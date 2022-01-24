package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type year struct {
	Posts []post
}

type post struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func yearSummary(jsonFile string) {
	var err error
	var relativePath string = "./"
	var jsonFilename []string = filenameParts(jsonFile, ".")
	var yearNum string = jsonFilename[0]
	var header = header(buildTitle(yearNum), "Post Archive", yearNum, relativePath)
	var outputStr strings.Builder
	var headerStr bytes.Buffer
	tpl := template.Must(template.ParseFiles("bootstrap/clean-blog/header.html.tpl"))
	tpl.Execute(&headerStr, header)
	var footer = footer(relativePath)
	var footerStr bytes.Buffer
	footerTpl := template.Must(template.ParseFiles("bootstrap/clean-blog/footer.html.tpl"))
	footerTpl.Execute(&footerStr, footer)

	yearData, _ := ioutil.ReadFile(jsonFile)
	year := year{}
	json.Unmarshal([]byte(yearData), &year)
	var postLinks string
	for _, post := range year.Posts {
		postLinks += fmt.Sprintf("<a href=\"%s\">%s</a><br />\n", post.Url, post.Title)
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
