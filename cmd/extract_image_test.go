package main

import "testing"

func TestExtractFirstImage(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "No image",
			html: "<p>Hello world</p>",
			want: "",
		},
		{
			name: "Simple image",
			html: `<p>Hello</p><img src="pic.jpg"><p>World</p>`,
			want: "pic.jpg",
		},
		{
			name: "Image with single quotes and extra attributes",
			html: `<img class="featured" src='path/to/img.png' alt="Image">`,
			want: "path/to/img.png",
		},
		{
			name: "Multiple images",
			html: `<img src="first.jpg"><img src="second.jpg">`,
			want: "first.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractFirstImage(tt.html)
			if got != tt.want {
				t.Errorf("extractFirstImage() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestResolveAbsoluteImageURL(t *testing.T) {
	tests := []struct {
		name      string
		imagePath string
		postPath  string
		want      string
	}{
		{
			name:      "Empty image path",
			imagePath: "",
			postPath:  "dist/2024/09/17/post.html",
			want:      "",
		},
		{
			name:      "Already absolute https",
			imagePath: "https://example.com/img.jpg",
			postPath:  "dist/2024/09/17/post.html",
			want:      "https://example.com/img.jpg",
		},
		{
			name:      "Relative to root",
			imagePath: "/images/logo.png",
			postPath:  "dist/2024/09/17/post.html",
			want:      "https://www.drincruz.com/images/logo.png",
		},
		{
			name:      "Relative to post directory",
			imagePath: "pexels-cottonbro-5302808.jpg",
			postPath:  "dist/2024/09/17/the-hardest-thing-i-have-done.html",
			want:      "https://www.drincruz.com/2024/09/17/pexels-cottonbro-5302808.jpg",
		},
		{
			name:      "Relative to top level post",
			imagePath: "about.jpg",
			postPath:  "dist/about.html",
			want:      "https://www.drincruz.com/about.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveAbsoluteImageURL(tt.imagePath, tt.postPath)
			if got != tt.want {
				t.Errorf("resolveAbsoluteImageURL() = %q, want %q", got, tt.want)
			}
		})
	}
}
