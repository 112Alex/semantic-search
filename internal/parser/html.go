package parser

import (
	"io"

	"golang.org/x/net/html"
)

// ParseTextNodes извлекает все текстовые узлы из HTML и возвращает их вместе с указателем на исходный *html.Node.
// reader: поток HTML.
// Возвращает срез указателей на текстовые узлы.
func ParseTextNodes(reader io.Reader) ([]*html.Node, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	var texts []*html.Node
	var walker func(*html.Node)
	walker = func(n *html.Node) {
		if n.Type == html.TextNode {
			texts = append(texts, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
	}

	walker(doc)
	return texts, nil
}
