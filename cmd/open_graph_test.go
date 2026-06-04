package main

import "testing"

func TestOgDefaultImage(t *testing.T) {
	want := "https://www.drincruz.com/favicon-32x32.png"
	got := OgDefaultImage()
	if got != want {
		t.Errorf("BaseURL: %s != %s", got, want)
	}
}
