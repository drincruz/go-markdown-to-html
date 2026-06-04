package main

import (
	"fmt"
	"log"
	"strings"
)

//	Return a URL based off of a `dist` path.
//
// dist/2026/05/04/sw.html -> https://www.drincruz.com/2026/05/04/sw.html
func distPathToUrl(path string) string {
	var relPostPath string
	switch {
	case strings.HasPrefix(path, "./"):
		relPostPath = strings.TrimPrefix(path, "./")
	default:
		relPostPath = strings.TrimPrefix(path, "dist/")
	}
	url := fmt.Sprintf("%s/%s", BaseURL(), relPostPath)
	log.Printf("[distPathToUrl] path: %s, url: %s", path, url)

	return url
}
