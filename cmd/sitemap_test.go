package main

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestPostDate(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{"dated post", "/2026/09/26/the-plumbing-of-holistic-engineering.html", "2026-09-26"},
		{"static page", "/about.html", ""},
		{"year page", "/2026.html", ""},
		{"malformed date", "/2026/9/26/post.html", ""},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postDate(tt.url); got != tt.want {
				t.Errorf("postDate(%q) = %q, want %q", tt.url, got, tt.want)
			}
		})
	}
}

func writeFixture(t *testing.T, dir string, name string, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatalf("Failed to write fixture %s: %s", name, err)
	}
}

func TestSitemapEntries(t *testing.T) {
	dir := t.TempDir()
	writeFixture(t, dir, "2024.json", `{"posts": [
		{"title": "Hardest", "url": "/2024/09/17/hardest.html"}
	]}`)
	writeFixture(t, dir, "2026.json", `{"posts": [
		{"title": "Plumbing", "url": "/2026/09/26/plumbing.html"},
		{"title": "Tradeoffs", "url": "/2026/07/29/tradeoffs.html"}
	]}`)

	want := []sitemapURL{
		{Loc: "https://www.drincruz.com/", LastMod: "2026-09-26"},
		{Loc: "https://www.drincruz.com/about.html"},
		{Loc: "https://www.drincruz.com/archive.html"},
		{Loc: "https://www.drincruz.com/2026.html", LastMod: "2026-09-26"},
		{Loc: "https://www.drincruz.com/2026/09/26/plumbing.html", LastMod: "2026-09-26"},
		{Loc: "https://www.drincruz.com/2026/07/29/tradeoffs.html", LastMod: "2026-07-29"},
		{Loc: "https://www.drincruz.com/2024.html", LastMod: "2024-09-17"},
		{Loc: "https://www.drincruz.com/2024/09/17/hardest.html", LastMod: "2024-09-17"},
	}
	got := sitemapEntries(dir)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sitemapEntries() =\n%+v\nwant\n%+v", got, want)
	}

	for _, entry := range got {
		for _, path := range robotsDisallowed {
			if strings.HasSuffix(entry.Loc, path) {
				t.Errorf("sitemap should not include disallowed page %s", entry.Loc)
			}
		}
	}
}

func TestBuildSitemap(t *testing.T) {
	entries := []sitemapURL{
		{Loc: "https://www.drincruz.com/", LastMod: "2026-09-26"},
		{Loc: "https://www.drincruz.com/about.html"},
	}
	out, err := buildSitemap(entries)
	if err != nil {
		t.Fatalf("buildSitemap() error: %s", err)
	}
	sitemap := string(out)

	if !strings.HasPrefix(sitemap, xml.Header) {
		t.Errorf("sitemap is missing the XML header:\n%s", sitemap)
	}
	if !strings.Contains(sitemap, `xmlns="`+sitemapNamespace+`"`) {
		t.Errorf("sitemap is missing the namespace:\n%s", sitemap)
	}
	if strings.Count(sitemap, "<lastmod>") != 1 {
		t.Errorf("sitemap should omit empty lastmod:\n%s", sitemap)
	}

	parsed := urlSet{}
	if err := xml.Unmarshal(out, &parsed); err != nil {
		t.Fatalf("sitemap does not parse: %s", err)
	}
	if !reflect.DeepEqual(parsed.URLs, entries) {
		t.Errorf("parsed URLs = %+v, want %+v", parsed.URLs, entries)
	}
}
