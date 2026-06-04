package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var imgRegex = regexp.MustCompile(`(?i)<img\s+[^>]*src=["']([^"']+)["']`)

// extractFirstImage finds the first <img> tag and returns its src value.
func extractFirstImage(htmlContent string) string {
	matches := imgRegex.FindStringSubmatch(htmlContent)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// resolveAbsoluteImageURL converts a potentially relative image path to an absolute URL.
func resolveAbsoluteImageURL(imagePath string, postPath string) string {
	if imagePath == "" {
		return ""
	}
	if strings.HasPrefix(imagePath, "http://") || strings.HasPrefix(imagePath, "https://") {
		return imagePath
	}

	// If it's a domain-relative path (e.g. /images/pic.png)
	if strings.HasPrefix(imagePath, "/") {
		return fmt.Sprintf("%s/%s", BaseURL(), strings.TrimPrefix(imagePath, "/"))
	}

	// Otherwise, it is relative to the post's directory (removing "dist/")
	relPostPath := strings.TrimPrefix(postPath, "dist/")
	dir := filepath.Dir(relPostPath)
	if dir == "." || dir == "/" {
		return fmt.Sprintf("%s/%s", BaseURL(), imagePath)
	}
	return fmt.Sprintf("%s/%s/%s", BaseURL(), dir, imagePath)
}
