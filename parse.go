package link

import (
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
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var retVal Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			retVal.Href = attr.Val
			break
		}
	}
	retVal.Text = "TODO: extract link text"
	return retVal
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
