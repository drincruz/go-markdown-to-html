package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Blog struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

// Define a struct that matches the JSON structure
type Year struct {
	Posts []Blog `json:"posts"`
}

func getMostRecent() {
	ex, err := os.Executable()
	check(err, "Failed to get executable")
	path := filepath.Dir(ex)
	jsonFiles := getFiles("json", path)
	if len(jsonFiles) == 0 {
		log.Fatalf("No JSON files found")
		return
	}
	posts := []Blog{}
	maxPosts := 3

	for _, file := range jsonFiles {
		jsonFile, err := os.ReadFile(file)
		check(err, "Failed to read file")
		data := Year{}
		if err := json.Unmarshal(jsonFile, &data); err != nil {
			log.Fatalf("Failed to unmarshal JSON: %s", err)
		}
		posts = append(posts, data.Posts...)
		if len(posts) >= maxPosts {
			posts = posts[:maxPosts]
			break
		}
	}
	fmt.Printf("Posts: %+v, Length: %d\n", posts, len(posts))
}
