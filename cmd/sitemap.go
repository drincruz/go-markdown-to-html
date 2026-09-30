package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const sitemapNamespace = "http://www.sitemaps.org/schemas/sitemap/0.9"

var postDateRegex = regexp.MustCompile(`^/(\d{4})/(\d{2})/(\d{2})/`)

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
}

type urlSet struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

// Return the ISO date of a post based off of its URL path.
//
// /2026/09/26/the-plumbing-of-holistic-engineering.html -> 2026-09-26
func postDate(url string) string {
	matches := postDateRegex.FindStringSubmatch(url)
	if matches == nil {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", matches[1], matches[2], matches[3])
}

// Return the most recent post date, or an empty string if there are none.
func newestPostDate(posts []Blog) string {
	var newest string
	for _, post := range posts {
		if date := postDate(post.Url); date > newest {
			newest = date
		}
	}
	return newest
}

func readYearPosts(jsonFile string) []Blog {
	data, err := os.ReadFile(jsonFile)
	check(err, "Failed to read year file")
	year := Year{}
	check(json.Unmarshal(data, &year), "Failed to unmarshal year file")
	return year.Posts
}

func yearPagePath(jsonFile string) string {
	yearNum := strings.TrimSuffix(filepath.Base(jsonFile), ".json")
	return fmt.Sprintf("/%s.html", yearNum)
}

func sitemapEntry(path string, lastMod string) sitemapURL {
	return sitemapURL{
		Loc:     BaseURL() + path,
		LastMod: lastMod,
	}
}

// Return the sitemap entries for the static pages, year pages and posts.
func sitemapEntries(jsonDir string) []sitemapURL {
	var yearEntries []sitemapURL
	var allPosts []Blog
	for _, jsonFile := range getFiles("json", jsonDir) {
		posts := readYearPosts(jsonFile)
		allPosts = append(allPosts, posts...)
		yearEntries = append(yearEntries, sitemapEntry(yearPagePath(jsonFile), newestPostDate(posts)))
		for _, post := range posts {
			yearEntries = append(yearEntries, sitemapEntry(post.Url, postDate(post.Url)))
		}
	}

	entries := []sitemapURL{
		sitemapEntry("/", newestPostDate(allPosts)),
		sitemapEntry("/about.html", ""),
		sitemapEntry("/archive.html", ""),
	}
	return append(entries, yearEntries...)
}

func buildSitemap(entries []sitemapURL) ([]byte, error) {
	set := urlSet{Xmlns: sitemapNamespace, URLs: entries}
	out, err := xml.MarshalIndent(set, "", "  ")
	if err != nil {
		return nil, err
	}
	return append([]byte(xml.Header), out...), nil
}

func writeSitemap(jsonDir string, outPath string) {
	sitemap, err := buildSitemap(sitemapEntries(jsonDir))
	check(err, "Failed to build sitemap")
	check(os.WriteFile(outPath, sitemap, 0o644), "Failed to write sitemap")
	log.Printf("[INFO][writeSitemap] Wrote %s", outPath)
}
