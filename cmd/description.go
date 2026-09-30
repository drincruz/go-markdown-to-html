package main

import (
	"log"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

const maxDescriptionLength = 160

// Return the text content of a node and all of its children.
func nodeText(n *html.Node) string {
	if n == nil {
		return ""
	}
	if n.Type == html.TextNode {
		return n.Data
	}
	var text strings.Builder
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		text.WriteString(nodeText(child))
	}
	return text.String()
}

// Cut text to at most maxLength characters on a word boundary.
func truncateWords(text string, maxLength int) string {
	if utf8.RuneCountInString(text) <= maxLength {
		return text
	}
	const ellipsis = "..."
	runes := []rune(text)
	cut := string(runes[:maxLength-len(ellipsis)])
	if i := strings.LastIndex(cut, " "); i > 0 {
		cut = cut[:i]
	}
	return strings.TrimRight(cut, " ,.;:") + ellipsis
}

// Return the first non-image paragraph of a post as plain text.
func firstParagraphText(htmlContent string) string {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Printf("[WARN][firstParagraphText] Failed to parse HTML: %s", err)
		return ""
	}
	p, err := cardTag(doc, "p")
	if err != nil {
		return ""
	}
	return strings.Join(strings.Fields(nodeText(p)), " ")
}

// Return a meta description for a post: its first paragraph, falling back
// to the subtitle, and then to the site default.
func postDescription(htmlContent string, subtitle string) string {
	if text := firstParagraphText(htmlContent); text != "" {
		return truncateWords(text, maxDescriptionLength)
	}
	if subtitle != "" {
		return subtitle
	}
	return OgDefaultDescription()
}
