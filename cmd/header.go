package main

// Header is Typical HTML data
// Header data is the "top" part of a site (metadata & such )
type Header struct {
	Title           string
	ContentTitle    string
	ContentSubtitle string
	RelativePath    string
	OpenGraphMeta   OpenGraphMetadata
}

type OpenGraphMetadata struct {
	Image string
	Type  string
	Url   string
}

type OpenGraphOptions func(*OpenGraphMetadata)

func OpenGraphImage(image string) OpenGraphOptions {
	return func(og *OpenGraphMetadata) {
		og.Image = image
	}
}

func NewHeader(t string, ct string, cs string, rp string, url string, ogType string, opts ...OpenGraphOptions) Header {
	og := &OpenGraphMetadata{
		Image: OgDefaultImage(),
		Type:  ogType,
		Url:   url,
	}

	for _, opt := range opts {
		opt(og)
	}

	header := Header{
		Title:           t,
		ContentTitle:    ct,
		ContentSubtitle: cs,
		RelativePath:    rp,
		OpenGraphMeta:   *og,
	}

	return header

}
