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

func TestPageTitle(t *testing.T) {
	tests := []struct {
		name  string
		title string
		want  string
	}{
		{"with title", "Tradeoffs in Engineering", "Tradeoffs in Engineering | drincruz.com"},
		{"empty title", "", "drincruz.com Blog"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pageTitle(tt.title); got != tt.want {
				t.Errorf("pageTitle(%q) = %q, want %q", tt.title, got, tt.want)
			}
		})
	}
}
