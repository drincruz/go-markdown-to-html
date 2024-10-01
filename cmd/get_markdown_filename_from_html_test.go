package main

import (
	"log"
	"testing"
)

func TestGetMarkdownFilenameFromHtml(t *testing.T) {
	log.Printf("[TestGetMarkdownFilenameFromHtml]\n")

	want := "markdown/2023/10/21/engineering-guilds-at-work.markdown"
	got := getMarkdownFilenameFromHtml("/2023/10/21/engineering-guilds-at-work.html")

	if got != want {
		t.Errorf("getMarkdownFilenameFromHtml: %s != %s", got, want)
	}
}
