package main

import "testing"

func TestFilenameParts(t *testing.T) {
	want := []string{"2021", "json"}
	got := filenameParts("2021.json", ".")
	if len(got) != len(want) {
		t.Errorf("filenameParts: %s != %s", got, want)
	}
	for index, val := range want {
		if val != got[index] {
			t.Errorf("filenameParts: %s != %s", got, want)
		}
	}
}
