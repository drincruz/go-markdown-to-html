package main

import "testing"

func TestOgDefaultImage(t *testing.T) {
	want := "https://www.drincruz.com/favicon-32x32.png"
	got := OgDefaultImage()
	if got != want {
		t.Errorf("BaseURL: %s != %s", got, want)
	}
}

func TestNewHeaderDescription(t *testing.T) {
	header := NewHeader("t", "ct", "cs", "./", "https://www.drincruz.com/", "website")
	if header.OpenGraphMeta.Description != OgDefaultDescription() {
		t.Errorf("default description = %q, want %q", header.OpenGraphMeta.Description, OgDefaultDescription())
	}

	header = NewHeader("t", "ct", "cs", "./", "https://www.drincruz.com/", "article", OpenGraphDescription("Custom"))
	if header.OpenGraphMeta.Description != "Custom" {
		t.Errorf("description = %q, want %q", header.OpenGraphMeta.Description, "Custom")
	}
}
