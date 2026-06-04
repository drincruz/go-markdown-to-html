package main

import "testing"

func TestDistPathToUrl(t *testing.T) {
	path := "dist/2024/09/17/the-hardest-thing-i-have-done.html"
	want := "https://www.drincruz.com/2024/09/17/the-hardest-thing-i-have-done.html"
	got := distPathToUrl(path)
	if got != want {
		t.Errorf("BaseURL: %s != %s", got, want)
	}
}
