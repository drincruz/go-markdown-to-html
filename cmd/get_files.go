package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

const LOG_PREFIX = "[getFiles]"

func getFiles(fileExtension string, path string) []string {
	path = strings.TrimSuffix(path, "/")
	fileSearch := fmt.Sprintf("%s/*.%s", path, fileExtension)
	log.Printf("%s Searching for files: %s\n", LOG_PREFIX, fileSearch)
	jsonFiles, err := filepath.Glob(fileSearch)
	check(err, "Failed to get JSON files")

	sort.Sort(sort.Reverse(sort.StringSlice(jsonFiles)))
	log.Printf("%s Files: %+v\n", LOG_PREFIX, jsonFiles)
	return jsonFiles
}
