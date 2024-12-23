package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {

}

func TextContent() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	printTexts(doc)

}

func printTexts(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	printTexts(n.FirstChild)
	printTexts(n.NextSibling)
}
