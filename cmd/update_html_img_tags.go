package main

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func UpdateHtmlImgTags(content []byte) []byte {
	reader := bytes.NewReader(content)
	context := &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body}

	nodes, err := html.ParseFragment(reader, context)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	for _, n := range nodes {
		UpdateHtmlImgNodes(n)
		if err := html.Render(&buf, n); err != nil {
			panic(err)
		}
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
