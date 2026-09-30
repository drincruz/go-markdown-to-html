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

// Return the YYYY.json files in a directory, newest year first.
func yearFiles(path string) []string {
	yearSearch := filepath.Join(path, "[0-9][0-9][0-9][0-9].json")
	files, err := filepath.Glob(yearSearch)
	check(err, "Failed to get year files")

	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	return files
}

// Return the year of a YYYY.json file path.
//
// ./2026.json -> 2026
func yearFromJsonFile(jsonFile string) string {
	return strings.TrimSuffix(filepath.Base(jsonFile), ".json")
}
