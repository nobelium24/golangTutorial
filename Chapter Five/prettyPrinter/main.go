package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// forEachNode(doc, startElement, endElement)

}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s='%s'", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf(">\n")
			depth++
		}

	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", text)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startEle(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if id == attr.Val && attr.Key == "id" {
				return false
			}
			return true
		}
	}
	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var result *html.Node
	pre := func(n *html.Node, id string) bool {
		if !startEle(n, id) {
			result = n
			return false
		}
		return true
	}
	forEachNodeTwo(doc, pre, nil, id)
	return result
}

func forEachNodeTwo(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) {
	if pre != nil && !pre(n, id) {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNodeTwo(c, pre, post, id)
	}
	if post != nil {
		post(n, id)
	}
}
