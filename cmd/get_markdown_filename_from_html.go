package main

import (
	"fmt"
	"strings"
)

func getMarkdownFilenameFromHtml(htmlFilename string) string {
	replaceFileType := strings.Replace(htmlFilename, ".html", ".markdown", 1)
	return fmt.Sprintf("markdown%s", replaceFileType)
}
