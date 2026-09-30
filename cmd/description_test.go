package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestTruncateWords(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		maxLength int
		want      string
	}{
		{"short text unchanged", "Hello there.", 20, "Hello there."},
		{"exact length unchanged", "Hello", 5, "Hello"},
		{"cut on word boundary", "The quick brown fox jumps", 15, "The quick..."},
		{"trailing punctuation trimmed", "One, two, three, four", 13, "One, two..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateWords(tt.text, tt.maxLength); got != tt.want {
				t.Errorf("truncateWords(%q, %d) = %q, want %q", tt.text, tt.maxLength, got, tt.want)
			}
		})
	}
}

func TestPostDescription(t *testing.T) {
	long := strings.Repeat("word ", 50)
	tests := []struct {
		name        string
		htmlContent string
		subtitle    string
		want        string
	}{
		{
			"first paragraph",
			"<h1>Title</h1><p>First <em>paragraph</em>\n here.</p><p>Second.</p>",
			"Subtitle",
			"First paragraph here.",
		},
		{
			"no space before punctuation after inline tags",
			`<p>Hence the <a href="https://getbootstrap.com">Bootstrap</a>. Welcome to <em>me</em>!</p>`,
			"",
			"Hence the Bootstrap. Welcome to me!",
		},
		{
			"skips image-only paragraph",
			`<p><img src="a.jpg" /></p><p>After the image.</p>`,
			"",
			"After the image.",
		},
		{
			"falls back to subtitle",
			"<h1>Only a heading</h1>",
			"Subtitle",
			"Subtitle",
		},
		{
			"falls back to default",
			"<h1>Only a heading</h1>",
			"",
			OgDefaultDescription(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postDescription(tt.htmlContent, tt.subtitle); got != tt.want {
				t.Errorf("postDescription() = %q, want %q", got, tt.want)
			}
		})
	}

	t.Run("truncates long paragraph", func(t *testing.T) {
		got := postDescription("<p>"+long+"</p>", "")
		if utf8.RuneCountInString(got) > maxDescriptionLength || !strings.HasSuffix(got, "word...") {
			t.Errorf("postDescription() = %q, want <= %d chars ending in word...", got, maxDescriptionLength)
		}
	})
}
