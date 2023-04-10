package link

import "io"

// Link represents a HTML link tag (<a href="...">)
type Link struct {
	Href string
	Text string
}

// Parse will accept a HTML document and will return a slice
// of links extracted from it.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
