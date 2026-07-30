package main

import (
	"strings"
	"testing"
)

func TestUpdateHtmlImgTags(t *testing.T) {
	tests := []struct {
		name           string
		html           string
		wantContains   string
		wantNotContain string
	}{
		{
			name:         "No image",
			html:         "<p>Hello world</p>",
			wantContains: "<p>Hello world</p>",
		},
		{
			name:         "Image with no class attribute",
			html:         `<img src="pic.jpg">`,
			wantContains: `<img src="pic.jpg" class="img-fluid"/>`,
		},
		{
			name:         "Image with existing class not containing img-fluid",
			html:         `<img src="pic.jpg" class="featured">`,
			wantContains: `class="featured img-fluid"`,
		},
		{
			name:           "Image with class already containing img-fluid",
			html:           `<img src="pic.jpg" class="featured img-fluid">`,
			wantContains:   `class="featured img-fluid"`,
			wantNotContain: `img-fluid img-fluid`,
		},
		{
			name:         "Multiple images",
			html:         `<img src="first.jpg"><img src="second.jpg" class="foo">`,
			wantContains: `<img src="first.jpg" class="img-fluid"/><img src="second.jpg" class="foo img-fluid"/>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(UpdateHtmlImgTags([]byte(tt.html)))
			if tt.wantContains != "" && !strings.Contains(got, tt.wantContains) {
				t.Errorf("UpdateHtmlImgTags() = %q, want substring %q", got, tt.wantContains)
			}
			if tt.wantNotContain != "" && strings.Contains(got, tt.wantNotContain) {
				t.Errorf("UpdateHtmlImgTags() = %q, unexpectedly contains %q", got, tt.wantNotContain)
			}
		})
	}
}

// UpdateHtmlImgTags is called on inner fragments (e.g. content.html.tpl's
// <div>-wrapped output, which itself sits inside <body> opened by
// header.html.tpl and closed by footer.html.tpl). It must not synthesize a
// document wrapper around the fragment, or the assembled page would end up
// with an <html>/<head>/<body> injected in the middle of the document.
func TestUpdateHtmlImgTagsDoesNotWrapFragment(t *testing.T) {
	input := `<div class="container"><div class="row"><div class="col-lg-8 col-md-10 mx-auto"><p>Hello</p><img src="x.jpg"></div></div></div>`
	got := string(UpdateHtmlImgTags([]byte(input)))

	for _, tag := range []string{"<html", "<head", "<body"} {
		if strings.Contains(got, tag) {
			t.Errorf("UpdateHtmlImgTags() = %q, unexpectedly contains %q wrapper", got, tag)
		}
	}
	if !strings.Contains(got, `class="img-fluid"`) {
		t.Errorf("UpdateHtmlImgTags() = %q, want img-fluid class applied", got)
	}
}
