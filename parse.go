package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a HTML link tag (<a href="...">)
type Link struct {
	Href string
	Text string
}

// Parse will accept a HTML document and will return a slice
// of links extracted from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var retVal []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retVal = append(retVal, linkNodes(c)...)
	}
	return retVal
}
