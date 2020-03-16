package main

import (
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	want := string("<h1>This is a title</h1>\n")
	got := string(MarkdownToHTML([]byte("# This is a title")))
	if got != want {
		t.Errorf("markdownToHTML: %s != %s", got, want)
	}
}
