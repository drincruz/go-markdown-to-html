package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Blog struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Url      string `json:"url"`
}

// Define a struct that matches the JSON structure
type Year struct {
	Posts []Blog `json:"posts"`
}

func getMostRecent() (posts []Blog) {
	jsonFiles := yearFiles(".")
	if len(jsonFiles) == 0 {
		log.Fatalf("No JSON files found")
		return []Blog{}
	}
	posts = []Blog{}
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
	return posts
}
