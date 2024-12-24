package main

import (
	"strings"

	"golang.org/x/net/html"
)

func main() {

}

func countWordsAndImages(n *html.Node) (words, images int) {

	if n == nil {
		return 0, 0
	}
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	childWords, childImages := countWordsAndImages(n.FirstChild)
	siblingWords, siblingImages := countWordsAndImages(n.NextSibling)

	words += childWords + siblingWords
	images += siblingImages + childImages

	return words, images
}
