package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestPostDistPath(t *testing.T) {
	want := "dist/2026/09/26/the-plumbing-of-holistic-engineering.html"
	got := postDistPath("/2026/09/26/the-plumbing-of-holistic-engineering.html")
	if got != want {
		t.Errorf("postDistPath() = %q, want %q", got, want)
	}
}

func TestCopyPostAssets(t *testing.T) {
	mdDir := t.TempDir()
	distDir := t.TempDir()
	writeFixture(t, mdDir, "post.markdown", "# Post")
	writeFixture(t, mdDir, "photo.jpg", "jpg bytes")
	writeFixture(t, mdDir, "diagram.png", "png bytes")
	if err := os.Mkdir(filepath.Join(mdDir, "subdir"), 0o755); err != nil {
		t.Fatalf("Failed to create subdir: %s", err)
	}

	copyPostAssets(mdDir, distDir)

	entries, err := os.ReadDir(distDir)
	if err != nil {
		t.Fatalf("Failed to read dist dir: %s", err)
	}
	var got []string
	for _, entry := range entries {
		got = append(got, entry.Name())
	}
	want := []string{"diagram.png", "photo.jpg"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("copyPostAssets() copied %v, want %v", got, want)
	}

	data, err := os.ReadFile(filepath.Join(distDir, "photo.jpg"))
	if err != nil || string(data) != "jpg bytes" {
		t.Errorf("photo.jpg content = %q (err %v), want %q", data, err, "jpg bytes")
	}
}

func TestYearPostLink(t *testing.T) {
	post := Blog{Title: "Files & Woes", Url: "/2016/03/21/woes.html"}
	want := "<a href=\"/2016/03/21/woes.html\">Files &amp; Woes</a><br />\n"
	if got := yearPostLink(post); got != want {
		t.Errorf("yearPostLink() = %q, want %q", got, want)
	}
}

// Every dated markdown post has a JSON entry and every JSON entry has a
// markdown post, so a post can't be silently left out of the build.
func TestEveryPostHasJsonEntry(t *testing.T) {
	const repoRoot = ".."

	var markdownPosts []string
	err := filepath.WalkDir(filepath.Join(repoRoot, "markdown"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(repoRoot, path)
		if !d.IsDir() && strings.HasSuffix(rel, ".markdown") && strings.Count(rel, "/") == 4 {
			markdownPosts = append(markdownPosts, rel)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Failed to walk markdown dir: %s", err)
	}

	var jsonPosts []string
	for _, jsonFile := range yearFiles(repoRoot) {
		for _, post := range readYearPosts(jsonFile) {
			jsonPosts = append(jsonPosts, getMarkdownFilenameFromHtml(post.Url))
		}
	}

	if len(markdownPosts) == 0 {
		t.Fatal("No markdown posts found")
	}
	if missing := missingFrom(markdownPosts, jsonPosts); len(missing) > 0 {
		t.Errorf("markdown posts with no CCYY.json entry: %v", missing)
	}
	if missing := missingFrom(jsonPosts, markdownPosts); len(missing) > 0 {
		t.Errorf("CCYY.json entries with no markdown post: %v", missing)
	}
}

// Return the items in want that are not in have.
func missingFrom(want []string, have []string) []string {
	seen := map[string]bool{}
	for _, item := range have {
		seen[item] = true
	}
	var missing []string
	for _, item := range want {
		if !seen[item] {
			missing = append(missing, item)
		}
	}
	return missing
}
