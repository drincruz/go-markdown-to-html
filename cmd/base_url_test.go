package main

import "testing"

func TestBaseUrl(t *testing.T) {
	want := "https://www.drincruz.com"
	got := BaseURL()
	if got != want {
		t.Errorf("BaseURL: %s != %s", got, want)
	}
}
