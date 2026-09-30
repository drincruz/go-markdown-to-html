package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
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

func TestGetFiles(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test_get_files")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(tmpDir)

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
}

func TestYearFiles(t *testing.T) {
	tmpDir := t.TempDir()
	for _, name := range []string{"2024.json", "2026.json", "package.json", "notes.json", "20260.json", "2026.txt"} {
		if err := os.WriteFile(filepath.Join(tmpDir, name), []byte("{}"), 0o644); err != nil {
			t.Fatalf("Failed to create file %s: %s", name, err)
		}
	}

	want := []string{
		filepath.Join(tmpDir, "2026.json"),
		filepath.Join(tmpDir, "2024.json"),
	}
	got := yearFiles(tmpDir)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("yearFiles() = %v, want %v", got, want)
	}
}

func TestYearFromJsonFile(t *testing.T) {
	for _, path := range []string{"2026.json", "./2026.json", "/abs/path/2026.json"} {
		if got := yearFromJsonFile(path); got != "2026" {
			t.Errorf("yearFromJsonFile(%q) = %q, want %q", path, got, "2026")
		}
	}
}
