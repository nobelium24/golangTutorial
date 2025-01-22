package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func main() {}

func Outline(n *html.Node, action string) {
	depth := 0
	if n.Type != html.ElementNode {
		return
	}
	var startElement func(arg *html.Node)
	var endElement func(arg *html.Node)

	if action == "start" {
		startElement = func(arg *html.Node) {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	} else if action == "end" {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}

	var forEachNode func(n *html.Node, pre, post func(n *html.Node))
	forEachNode = func(n *html.Node, pre, post func(n *html.Node)) {
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

	forEachNode(n, startElement, endElement)

}

// func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
// 	if pre != nil {
// 	pre(n)
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 	forEachNode(c, pre, post)
// 	}
// 	if post != nil {
// 	post(n)
// 	}
// 	}

// func startElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 	fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
// 	depth++
// 	}
// 	}
// 	func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 	depth--
// 	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 	}
// 	}
