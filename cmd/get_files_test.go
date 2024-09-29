package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func getTempFiles() (tempFiles []string) {
	return []string{
		"2021.json",
		"2022.json",
		"2023.json",
		"2024.json",
	}
}

func createTempFiles(tmpDir string) {
	for _, file := range getTempFiles() {
		filename := fmt.Sprintf("%s/%s", tmpDir, file)
		_, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %s", filename)
		}
	}
}

func removeTempFiles(tmpDir string) {
	for _, file := range getTempFiles() {
		filename := fmt.Sprintf("%s/%s", tmpDir, file)
		f, err := os.Open(filename)
		check(err, "Failed to open file")
		f.Close()
		err = os.Remove(filename)
		if err != nil {
			log.Fatalf("Failed to remove file: %s", file)
		}
	}
}

func TestGetFiles(t *testing.T) {
	tmpDir := os.TempDir()
	tmpDir = strings.TrimSuffix(tmpDir, "/")
	log.Printf("[TestGetFiles] TempDir: %s\n", tmpDir)
	createTempFiles(tmpDir)

	want := []string{
		fmt.Sprintf("%s/2024.json", tmpDir),
		fmt.Sprintf("%s/2023.json", tmpDir),
		fmt.Sprintf("%s/2022.json", tmpDir),
		fmt.Sprintf("%s/2021.json", tmpDir),
	}
	got := getFiles("json", tmpDir)
	if len(got) != len(want) {
		t.Errorf("getFiles: %d != %d", len(got), len(want))
		return
	}
	for index, val := range want {
		if got[index] != val {
			t.Errorf("getFiles: %s != %s", got[index], want[index])
		}
	}

	removeTempFiles(tmpDir)
}
