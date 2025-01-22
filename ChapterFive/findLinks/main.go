package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Print(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode &&
		(n.Data == "a" || n.Data == "link" || n.Data == "area" || n.Data == "base" || n.Data == "img" || n.Data == "iframe" || n.Data == "script" ||
			n.Data == "audio" || n.Data == "video" || n.Data == "source" || n.Data == "track" || n.Data == "embed" || n.Data == "input" || n.Data == "object") {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}

	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	// return links

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links

}
