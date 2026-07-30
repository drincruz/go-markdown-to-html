package main

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

func UpdateHtmlImgTags(content []byte) []byte {
	reader := bytes.NewReader(content)

	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	UpdateHtmlImgNodes(doc)

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func UpdateHtmlImgNodes(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		classIndex := -1
		for i, attr := range n.Attr {
			if attr.Key == "class" {
				classIndex = i
				break
			}
		}

		if classIndex == -1 {
			n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "img-fluid"})
		} else {
			classes := strings.Fields(n.Attr[classIndex].Val)
			hasImgFluid := false
			for _, c := range classes {
				if c == "img-fluid" {
					hasImgFluid = true
					break
				}
			}
			if !hasImgFluid {
				classes = append(classes, "img-fluid")
				n.Attr[classIndex].Val = strings.Join(classes, " ")
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		UpdateHtmlImgNodes(c)
	}
}
