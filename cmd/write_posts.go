package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Return the dist path of a post based off of its URL path.
//
// /2026/09/26/the-plumbing-of-holistic-engineering.html -> dist/2026/09/26/the-plumbing-of-holistic-engineering.html
func postDistPath(url string) string {
	return "dist" + url
}

// Copy every file in a post's markdown directory, other than the markdown
// itself, into its dist directory (i.e. images).
func copyPostAssets(mdDir string, distDir string) {
	entries, err := os.ReadDir(mdDir)
	check(err, "Failed to read post directory")
	for _, entry := range entries {
		if entry.IsDir() || strings.HasSuffix(entry.Name(), ".markdown") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(mdDir, entry.Name()))
		check(err, "Failed to read post asset")
		check(os.WriteFile(filepath.Join(distDir, entry.Name()), data, 0o644), "Failed to write post asset")
	}
}

// Write the HTML page and assets of a single post.
func writePost(post Blog) {
	mdPath := getMarkdownFilenameFromHtml(post.Url)
	outPath := postDistPath(post.Url)
	log.Printf("[INFO][writePost] %s -> %s", mdPath, outPath)
	check(os.MkdirAll(filepath.Dir(outPath), 0o755), "Failed to create post directory")
	writeBlog(mdPath, post.Title, post.Subtitle, outPath)
	copyPostAssets(filepath.Dir(mdPath), filepath.Dir(outPath))
}

// Write every post listed in the year JSON files in jsonDir.
func writePosts(jsonDir string) {
	for _, jsonFile := range yearFiles(jsonDir) {
		for _, post := range readYearPosts(jsonFile) {
			writePost(post)
		}
	}
}
